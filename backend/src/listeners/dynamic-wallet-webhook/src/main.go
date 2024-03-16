package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var svc *dynamodb.Client

func Init() {
	accessKey := os.Getenv("HACKATHON_ACCESS_KEY")
	secretKey := os.Getenv("HACKATHON_SECRET_KEY")
	if accessKey == "" || secretKey == "" {
		fmt.Println("Access Key and Secret Key not found")
		return
	}

	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion("eu-west-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	if err != nil {
		fmt.Printf("error loading AWS configuration: %v\n", err)
		return
	}

	svc = dynamodb.NewFromConfig(cfg)
}

func init() {
	Init()
}

type User struct {
	ID             string         `json:"id" dynamodbav:"id"`
	GitHubUsername string         `json:"githubUsername" dynamodbav:"githubUsername"`
	Typename       string         `json:"typename" dynamodbav:"typename"`
	Data           WalletLinkData `json:"metadata" dynamodbav:"metadata"`
	CreatedAt      time.Time      `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt" dynamodbav:"updatedAt"`
}

type WebhookPayload struct {
	Event                 string                 `json:"event"`
	Network               map[string]interface{} `json:"network"`
	ListenerId            string                 `json:"listenerId"`
	Abi                   map[string]interface{} `json:"abi"`
	Api                   string                 `json:"api"`
	Args                  []interface{}          `json:"args"`
	Headers               map[string]interface{} `json:"headers"`
	Method                string                 `json:"method"`
	QueryStringParameters map[string]interface{} `json:"queryStringParameters"`
}

func HandleRequest(ctx context.Context, request WebhookPayload) (events.APIGatewayProxyResponse, error) {
	jsonString, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	// Just for Logging
	fmt.Printf("Request: %s\n", jsonString)

	var eventMap map[string]interface{}
	if err := json.Unmarshal([]byte(request.Event), &eventMap); err != nil {
		log.Fatalf("Failed to unmarshal Event string: %s", err)
	}

	// Unmarshal the WebhookPayload.Event into a struct SocialLinkedEvent or SocialUnlinkEvent
	// This will allow us to switch on the event name and perform the correct action
	eventName, ok := eventMap["eventName"].(string)
	if !ok {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Bad Request: eventName not found"}, nil
	}

	println("Event Type: ", eventName)
	if eventName == "wallet.linked" {
		var linkedEvent WalletLinkEvent
		if err := json.Unmarshal([]byte(request.Event), &linkedEvent); err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error: Unmarshalling linked event failed"}, nil
		}
		fmt.Println("Linked Event: ", linkedEvent)

		metadataAV, err := attributevalue.MarshalMap(linkedEvent.Data)
		fmt.Println("Metadata: ", metadataAV)
		if err != nil {
			fmt.Printf("Failed to marshal WalletLinkData: %v\n", err)
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal Server Error",
			}, err
		}

		updates := map[string]types.AttributeValue{
			"updatedAt": &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
			"metadata":  &types.AttributeValueMemberM{Value: metadataAV},
		}

		err = UpdateItem(ctx, "gitpaid", linkedEvent.Data.UserID, updates)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal Server Error",
			}, err
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "OK",
		}, err
	}

	if eventName == "user.social.linked" {
		var linkedEvent SocialLinkedEvent
		if err := json.Unmarshal([]byte(request.Event), &linkedEvent); err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error: Unmarshalling linked event failed"}, nil
		}

		githubUsername := linkedEvent.Data.AccountUsername
		println(("Github Username: " + githubUsername))
		ctx := context.TODO()

		keyCond := expression.KeyEqual(expression.Key("typename"), expression.Value("User"))
		filter := expression.Name("githubUsername").Equal(expression.Value(githubUsername))

		queryInput := QueryInput{
			TableName:    "gitpaid",
			IndexName:    "byTypename",
			KeyCondition: keyCond,
			Filter:       filter,
		}

		var users []User

		users, err = QueryDynamoDB[User](ctx, queryInput)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal Server Error",
			}, err
		}

		// If the user is not in the database, add them
		if len(users) == 0 {
			updates := map[string]types.AttributeValue{
				"githubUsername": &types.AttributeValueMemberS{Value: githubUsername},
				"typename":       &types.AttributeValueMemberS{Value: "User"},
				"createdAt":      &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
			}

			err = UpdateItem(ctx, "gitpaid", linkedEvent.Data.UserID, updates)
			if err != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: 500,
					Body:       "Internal Server Error",
				}, err
			}
		}

		// Fetch the repos for the user, for each of the repos add them into dynamo as this user as the owner with typename of Repository
		repos, err := ListRepos(githubUsername)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal Server Error",
			}, err
		}

		// for each repo, add it to dynamo
		for _, repo := range repos {
			itemID := strconv.Itoa(repo.ID)
			item := NewDynamoItem[GithubRepo](itemID, repo, "Repository")
			err = PutItemInDynamoDB[DynamoItem](ctx, "gitpaid", item)
			if err != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: 500,
					Body:       "Internal Server Error",
				}, err
			}
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "OK",
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body:       "Couldnt find event type",
	}, nil
}

func GetItemFromDynamoDB[T any](tableName string, key map[string]types.AttributeValue) (T, error) {
	var item T

	// Make the GetItem API call
	result, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})
	if err != nil {
		return item, fmt.Errorf("failed to get item from DynamoDB: %w", err)
	}

	// Unmarshal the result into the provided type
	err = attributevalue.UnmarshalMap(result.Item, &item)
	if err != nil {
		return item, fmt.Errorf("failed to unmarshal DynamoDB item into type: %w", err)
	}

	return item, nil
}

func PutItemInDynamoDB[T any](ctx context.Context, tableName string, item T) error {
	// Marshal the item into a map of DynamoDB attribute values
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("failed to marshal item into DynamoDB attribute values: %w", err)
	}

	// Put the item into the specified DynamoDB table
	fmt.Println("Adding item to DynamoDB table", item)
	_, err = svc.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	if err != nil {
		return fmt.Errorf("failed to put item into DynamoDB table: %w", err)
	}

	return nil
}

func UpdateItem(ctx context.Context, tableName, id string, updates map[string]types.AttributeValue) error {
	key, err := attributevalue.MarshalMap(map[string]string{"id": id})
	if err != nil {
		return fmt.Errorf("failed to marshal key: %w", err)
	}

	updateExpr := "SET "
	exprAttrValues := make(map[string]types.AttributeValue)
	i := 0
	for attr, val := range updates {
		placeholder := ":val" + fmt.Sprint(i)
		updateExpr += fmt.Sprintf("%s = %s, ", attr, placeholder)
		exprAttrValues[placeholder] = val
		i++
	}
	updateExpr = updateExpr[:len(updateExpr)-2] // Trim the trailing ", "

	_, err = svc.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(tableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpr),
		ExpressionAttributeValues: exprAttrValues,
	})
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	return nil
}

// QueryDynamoDB performs a query operation on DynamoDB with the given input and returns items of type T.
func QueryDynamoDB[T any](ctx context.Context, input QueryInput) ([]T, error) {
	// Initialize the expression builder with the key condition.
	exprBuilder := expression.NewBuilder().WithKeyCondition(input.KeyCondition)

	// Assuming we have a way to know if a filter should be applied,
	// for example, a boolean flag in QueryInput or based on application logic.
	// Since we can't check the filter directly, this part of your application
	// logic needs to decide when to add the filter to the expression.
	if input.Filter.IsSet() {
		exprBuilder = exprBuilder.WithFilter(input.Filter)
	}

	// Build the final expression.
	expr, err := exprBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf("error building expression: %w", err)
	}

	// Prepare the DynamoDB query input.
	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(input.TableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	// Include the index name if provided.
	if input.IndexName != "" {
		queryInput.IndexName = aws.String(input.IndexName)
	}

	// Execute the query.
	resp, err := svc.Query(ctx, queryInput)
	if err != nil {
		return nil, fmt.Errorf("failed to query DynamoDB: %w", err)
	}

	// Unmarshal the results into the slice of T.
	var items []T
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &items)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB response items: %w", err)
	}

	return items, nil
}

// Add in pagination for this later
func ListRepos(githubUsername string) ([]GithubRepo, error) {
	url := "https://api.github.com/users/" + githubUsername + "/repos?per_page=100"
	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
	}

	resp, err := FetchData[[]GithubRepo](url, headers)
	if err != nil {
		panic(err)
	}

	return resp, nil
}

func FetchData[T any](url string, headers map[string]string) (T, error) {
	var result T // Initialize a variable of type T

	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

type GithubRepo struct {
	ID                       int       `json:"id" dynamodbav:"id"`
	NodeID                   string    `json:"node_id" dynamodbav:"node_id"`
	Name                     string    `json:"name" dynamodbav:"name"`
	FullName                 string    `json:"full_name" dynamodbav:"full_name"`
	Private                  bool      `json:"private" dynamodbav:"private"`
	Owner                    Owner     `json:"owner" dynamodbav:"owner"`
	HTMLURL                  string    `json:"html_url" dynamodbav:"html_url"`
	Description              any       `json:"description" dynamodbav:"description"`
	Fork                     bool      `json:"fork" dynamodbav:"fork"`
	URL                      string    `json:"url" dynamodbav:"url"`
	ForksURL                 string    `json:"forks_url" dynamodbav:"forks_url"`
	KeysURL                  string    `json:"keys_url" dynamodbav:"keys_url"`
	CollaboratorsURL         string    `json:"collaborators_url" dynamodbav:"collaborators_url"`
	TeamsURL                 string    `json:"teams_url" dynamodbav:"teams_url"`
	HooksURL                 string    `json:"hooks_url" dynamodbav:"hooks_url"`
	IssueEventsURL           string    `json:"issue_events_url" dynamodbav:"issue_events_url"`
	EventsURL                string    `json:"events_url" dynamodbav:"events_url"`
	AssigneesURL             string    `json:"assignees_url" dynamodbav:"assignees_url"`
	BranchesURL              string    `json:"branches_url" dynamodbav:"branches_url"`
	TagsURL                  string    `json:"tags_url" dynamodbav:"tags_url"`
	BlobsURL                 string    `json:"blobs_url" dynamodbav:"blobs_url"`
	GitTagsURL               string    `json:"git_tags_url" dynamodbav:"git_tags_url"`
	GitRefsURL               string    `json:"git_refs_url" dynamodbav:"git_refs_url"`
	TreesURL                 string    `json:"trees_url" dynamodbav:"trees_url"`
	StatusesURL              string    `json:"statuses_url" dynamodbav:"statuses_url"`
	LanguagesURL             string    `json:"languages_url" dynamodbav:"languages_url"`
	StargazersURL            string    `json:"stargazers_url" dynamodbav:"stargazers_url"`
	ContributorsURL          string    `json:"contributors_url" dynamodbav:"contributors_url"`
	SubscribersURL           string    `json:"subscribers_url" dynamodbav:"subscribers_url"`
	SubscriptionURL          string    `json:"subscription_url" dynamodbav:"subscription_url"`
	CommitsURL               string    `json:"commits_url" dynamodbav:"commits_url"`
	GitCommitsURL            string    `json:"git_commits_url" dynamodbav:"git_commits_url"`
	CommentsURL              string    `json:"comments_url" dynamodbav:"comments_url"`
	IssueCommentURL          string    `json:"issue_comment_url" dynamodbav:"issue_comment_url"`
	ContentsURL              string    `json:"contents_url" dynamodbav:"contents_url"`
	CompareURL               string    `json:"compare_url" dynamodbav:"compare_url"`
	MergesURL                string    `json:"merges_url" dynamodbav:"merges_url"`
	ArchiveURL               string    `json:"archive_url" dynamodbav:"archive_url"`
	DownloadsURL             string    `json:"downloads_url" dynamodbav:"downloads_url"`
	IssuesURL                string    `json:"issues_url" dynamodbav:"issues_url"`
	PullsURL                 string    `json:"pulls_url" dynamodbav:"pulls_url"`
	MilestonesURL            string    `json:"milestones_url" dynamodbav:"milestones_url"`
	NotificationsURL         string    `json:"notifications_url" dynamodbav:"notifications_url"`
	LabelsURL                string    `json:"labels_url" dynamodbav:"labels_url"`
	ReleasesURL              string    `json:"releases_url" dynamodbav:"releases_url"`
	DeploymentsURL           string    `json:"deployments_url" dynamodbav:"deployments_url"`
	CreatedAt                time.Time `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt                time.Time `json:"updated_at" dynamodbav:"updated_at"`
	PushedAt                 time.Time `json:"pushed_at" dynamodbav:"pushed_at"`
	GitURL                   string    `json:"git_url" dynamodbav:"git_url"`
	SSHURL                   string    `json:"ssh_url" dynamodbav:"ssh_url"`
	CloneURL                 string    `json:"clone_url" dynamodbav:"clone_url"`
	SvnURL                   string    `json:"svn_url" dynamodbav:"svn_url"`
	Homepage                 any       `json:"homepage" dynamodbav:"homepage"`
	Size                     int       `json:"size" dynamodbav:"size"`
	StargazersCount          int       `json:"stargazers_count" dynamodbav:"stargazers_count"`
	WatchersCount            int       `json:"watchers_count" dynamodbav:"watchers_count"`
	Language                 any       `json:"language" dynamodbav:"language"`
	HasIssues                bool      `json:"has_issues" dynamodbav:"has_issues"`
	HasProjects              bool      `json:"has_projects" dynamodbav:"has_projects"`
	HasDownloads             bool      `json:"has_downloads" dynamodbav:"has_downloads"`
	HasWiki                  bool      `json:"has_wiki" dynamodbav:"has_wiki"`
	HasPages                 bool      `json:"has_pages" dynamodbav:"has_pages"`
	HasDiscussions           bool      `json:"has_discussions" dynamodbav:"has_discussions"`
	ForksCount               int       `json:"forks_count" dynamodbav:"forks_count"`
	MirrorURL                any       `json:"mirror_url" dynamodbav:"mirror_url"`
	Archived                 bool      `json:"archived" dynamodbav:"archived"`
	Disabled                 bool      `json:"disabled" dynamodbav:"disabled"`
	OpenIssuesCount          int       `json:"open_issues_count" dynamodbav:"open_issues_count"`
	License                  any       `json:"license" dynamodbav:"license"`
	AllowForking             bool      `json:"allow_forking" dynamodbav:"allow_forking"`
	IsTemplate               bool      `json:"is_template" dynamodbav:"is_template"`
	WebCommitSignoffRequired bool      `json:"web_commit_signoff_required" dynamodbav:"web_commit_signoff_required"`
	Topics                   []any     `json:"topics" dynamodbav:"topics"`
	Visibility               string    `json:"visibility" dynamodbav:"visibility"`
	Forks                    int       `json:"forks" dynamodbav:"forks"`
	OpenIssues               int       `json:"open_issues" dynamodbav:"open_issues"`
	Watchers                 int       `json:"watchers" dynamodbav:"watchers"`
	DefaultBranch            string    `json:"default_branch" dynamodbav:"default_branch"`
}

type DynamoItem struct {
	ID        string    `dynamodbav:"id"` // Ensure this tag matches DynamoDB's key name
	CreatedAt time.Time `dynamodbav:"createdAt"`
	UpdatedAt time.Time `dynamodbav:"updatedAt"`
	Typename  string    `dynamodbav:"typename"`
	Data      any       `dynamodbav:"data"`
}

func NewDynamoItem[T any](id string, data T, typename string) DynamoItem {
	return DynamoItem{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Typename:  typename,
		Data:      data,
	}
}

type Owner struct {
	Login             string `json:"login" dynamodbav:"login"`
	ID                int    `json:"id" dynamodbav:"id"`
	NodeID            string `json:"node_id" dynamodbav:"node_id"`
	AvatarURL         string `json:"avatar_url" dynamodbav:"avatar_url"`
	GravatarID        string `json:"gravatar_id" dynamodbav:"gravatar_id"`
	URL               string `json:"url" dynamodbav:"url"`
	HTMLURL           string `json:"html_url" dynamodbav:"html_url"`
	FollowersURL      string `json:"followers_url" dynamodbav:"followers_url"`
	FollowingURL      string `json:"following_url" dynamodbav:"following_url"`
	GistsURL          string `json:"gists_url" dynamodbav:"gists_url"`
	StarredURL        string `json:"starred_url" dynamodbav:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url" dynamodbav:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url" dynamodbav:"organizations_url"`
	ReposURL          string `json:"repos_url" dynamodbav:"repos_url"`
	EventsURL         string `json:"events_url" dynamodbav:"events_url"`
	ReceivedEventsURL string `json:"received_events_url" dynamodbav:"received_events_url"`
	Type              string `json:"type" dynamodbav:"type"`
	SiteAdmin         bool   `json:"site_admin" dynamodbav:"site_admin"`
}

type WalletLinkEvent struct {
	EventID         string         `json:"eventId"`
	WebhookID       string         `json:"webhookId"`
	EnvironmentID   string         `json:"environmentId"`
	Data            WalletLinkData `json:"data"`
	EnvironmentName string         `json:"environmentName"`
	MessageID       string         `json:"messageId"`
	EventName       string         `json:"eventName"`
	UserID          string         `json:"userId"`
	Timestamp       time.Time      `json:"timestamp"`
}

type WalletLinkData struct {
	Chain             string    `json:"chain" dynamodbav:"chain"`
	LowerPublicKey    string    `json:"lowerPublicKey" dynamodbav:"lowerPublicKey"`
	PublicKey         string    `json:"publicKey" dynamodbav:"publicKey"`
	UserID            string    `json:"userId" dynamodbav:"userId"`
	TurnkeyHDWalletID string    `json:"turnkeyHDWalletId" dynamodbav:"turnkeyHDWalletId"`
	CreatedAt         time.Time `json:"createdAt" dynamodbav:"createdAt"`
	DeletedAt         any       `json:"deletedAt" dynamodbav:"deletedAt"`
	Provider          string    `json:"provider" dynamodbav:"provider"`
	Name              string    `json:"name" dynamodbav:"name"`
	ID                string    `json:"id" dynamodbav:"id"`
	HardwareWallet    any       `json:"hardwareWallet" dynamodbav:"hardwareWallet"`
	SignerWalletID    any       `json:"signerWalletId" dynamodbav:"signerWalletId"`
	UpdatedAt         time.Time `json:"updatedAt" dynamodbav:"updatedAt"`
}

type SocialLinkedEvent struct {
	EventID         string     `json:"eventId"`
	WebhookID       string     `json:"webhookId"`
	EnvironmentID   string     `json:"environmentId"`
	Data            SocialData `json:"data"`
	EnvironmentName string     `json:"environmentName"`
	MessageID       string     `json:"messageId"`
	EventName       string     `json:"eventName"`
	UserID          string     `json:"userId"`
	Timestamp       time.Time  `json:"timestamp"`
}

type Plan struct {
	Name          string `json:"name"`
	Collaborators int    `json:"collaborators"`
	PrivateRepos  int    `json:"private_repos"`
	Space         int    `json:"space"`
}

type AccountProfile struct {
	GistsURL                string    `json:"gists_url"`
	ReposURL                string    `json:"repos_url"`
	TwoFactorAuthentication bool      `json:"two_factor_authentication"`
	FollowingURL            string    `json:"following_url"`
	TwitterUsername         any       `json:"twitter_username"`
	Bio                     any       `json:"bio"`
	CreatedAt               time.Time `json:"created_at"`
	Login                   string    `json:"login"`
	Type                    string    `json:"type"`
	Blog                    string    `json:"blog"`
	PrivateGists            int       `json:"private_gists"`
	TotalPrivateRepos       int       `json:"total_private_repos"`
	SubscriptionsURL        string    `json:"subscriptions_url"`
	UpdatedAt               time.Time `json:"updated_at"`
	SiteAdmin               bool      `json:"site_admin"`
	DiskUsage               int       `json:"disk_usage"`
	Collaborators           int       `json:"collaborators"`
	Company                 any       `json:"company"`
	OwnedPrivateRepos       int       `json:"owned_private_repos"`
	ID                      int       `json:"id"`
	PublicRepos             int       `json:"public_repos"`
	GravatarID              string    `json:"gravatar_id"`
	Plan                    Plan      `json:"plan"`
	Email                   any       `json:"email"`
	OrganizationsURL        string    `json:"organizations_url"`
	Hireable                any       `json:"hireable"`
	StarredURL              string    `json:"starred_url"`
	FollowersURL            string    `json:"followers_url"`
	PublicGists             int       `json:"public_gists"`
	URL                     string    `json:"url"`
	ReceivedEventsURL       string    `json:"received_events_url"`
	Followers               int       `json:"followers"`
	AvatarURL               string    `json:"avatar_url"`
	EventsURL               string    `json:"events_url"`
	HTMLURL                 string    `json:"html_url"`
	Following               int       `json:"following"`
	Name                    string    `json:"name"`
	Location                any       `json:"location"`
	NodeID                  string    `json:"node_id"`
}

type SocialData struct {
	AccountDisplayName   string         `json:"accountDisplayName"`
	WalletID             any            `json:"walletId"`
	AccountPhotos        []string       `json:"accountPhotos"`
	AccountProfile       AccountProfile `json:"accountProfile"`
	AccountEmails        []string       `json:"accountEmails"`
	UserID               string         `json:"userId"`
	AccountID            string         `json:"accountId"`
	CreatedAt            time.Time      `json:"createdAt"`
	DeletedAt            any            `json:"deletedAt"`
	ProjectEnvironmentID string         `json:"projectEnvironmentId"`
	UserEmailID          any            `json:"userEmailId"`
	Provider             string         `json:"provider"`
	AccountUsername      string         `json:"accountUsername"`
	ID                   string         `json:"id"`
	UpdatedAt            time.Time      `json:"updatedAt"`
}

type QueryInput struct {
	TableName    string
	IndexName    string // Optional
	KeyCondition expression.KeyConditionBuilder
	Filter       expression.ConditionBuilder  // Optional
	Projection   expression.ProjectionBuilder // Optional, if you want to specify returned attributes
}

func main() {
	lambda.Start(HandleRequest)
}
