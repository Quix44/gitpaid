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
)

const MaxBodyBytes = 1048576 // 1MB

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

func HandleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	jsonString, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("Request: %s\n", jsonString)

	// If the user is asking for issues
	if request.QueryStringParameters["issues"] == "true" {
		return handleListIssuesRequest(request)
	}

	// If the user is asking for repos
	if request.QueryStringParameters["repos"] == "true" {
		return handleListReposRequest(request)
	}

	// If the user is asking for labels
	if request.QueryStringParameters["labels"] == "true" {
		return handleListLabelsRequest(request)
	}

	// If we should import a single repo
	if request.QueryStringParameters["import"] == "true" {
		return handleImportRequest(request)
	}

	// Catch all throw
	return events.APIGatewayProxyResponse{
		StatusCode: 405,
		Body:       "Method Not Allowed",
	}, nil
}

// handleGetRequest handles GET requests
func handleListReposRequest(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	targetUser := request.QueryStringParameters["user"]

	keyCondition := expression.Key("typename").Equal(expression.Value("Repository"))
	filter := expression.Name("data.owner.login").Equal(expression.Value(targetUser))

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()

	if err != nil {
		fmt.Println("Got error building expression:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}

	// Prepare the query input
	queryInput := dynamodb.QueryInput{
		TableName:                 aws.String("gitpaid"),
		IndexName:                 aws.String("byTypename"),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	ctx := context.TODO()
	items, err := QueryDynamoDB[DynamoItem](ctx, queryInput)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}
	itemsJSON, err := json.Marshal(items)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err

	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(itemsJSON),
	}, nil
}

func handleListLabelsRequest(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	targetUser := request.QueryStringParameters["user"]

	keyCondition := expression.Key("typename").Equal(expression.Value("Label"))
	filter := expression.Name("data.Sender.Login").Equal(expression.Value(targetUser))

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()

	if err != nil {
		fmt.Println("Got error building expression:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}

	// Prepare the query input
	queryInput := dynamodb.QueryInput{
		TableName:                 aws.String("gitpaid"),
		IndexName:                 aws.String("byTypename"),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	ctx := context.TODO()
	items, err := QueryDynamoDB[DynamoItem](ctx, queryInput)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}
	itemsJSON, err := json.Marshal(items)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(itemsJSON),
	}, nil
}

func handleListIssuesRequest(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	// targetUser := request.QueryStringParameters["user"]
	keyCondition := expression.Key("typename").Equal(expression.Value("Issue"))
	filter := expression.Name("data.Action").NotEqual(expression.Value("closed"))

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()

	if err != nil {
		fmt.Println("Got error building expression:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}

	// Prepare the query input
	queryInput := dynamodb.QueryInput{
		TableName:                 aws.String("gitpaid"),
		IndexName:                 aws.String("byTypename"),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	ctx := context.TODO()
	items, err := QueryDynamoDB[DynamoItem](ctx, queryInput)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}
	itemsJSON, err := json.Marshal(items)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(itemsJSON),
	}, nil
}

type RepositoryImport struct {
	Username       string `json:"username"`
	RepositoryName string `json:"repositoryName"`
}

// handlePostRequest handles POST requests
func handleImportRequest(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	// unmarshal the body into the import repo struct
	requestBody := request.Body
	var importRepo RepositoryImport
	err := json.Unmarshal([]byte(requestBody), &importRepo)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Bad Request",
		}, err
	}

	repos, err := ListRepos(importRepo.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}

	ctx := context.TODO()
	for i := 0; i < len(repos); i++ {
		if (repos[i].Name == importRepo.RepositoryName) && (repos[i].Owner.Login == importRepo.Username) {
			itemID := strconv.Itoa(repos[i].ID)
			item := NewDynamoItem[GithubRepo](itemID, repos[i], "Repository")
			err = PutItemInDynamoDB[DynamoItem](ctx, "gitpaid", item)
			if err != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: 500,
					Body:       "Internal Server Error",
				}, err
			}
		}
	}

	// Example: Echo back the request body
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("POST request body: %s", request.Body),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

func ListRepos(githubUsername string) ([]GithubRepo, error) {
	url := "https://api.github.com/users/" + githubUsername + "/repos"
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

func QueryDynamoDB[T any](ctx context.Context, queryInput dynamodb.QueryInput) ([]T, error) {
	// Execute the query.
	resp, err := svc.Query(ctx, &queryInput)
	if err != nil {
		return nil, fmt.Errorf("failed to query DynamoDB: %w", err)
	}

	var items []T
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &items)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB response items: %w", err)
	}

	return items, nil
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

type DynamoItem struct {
	ID        string    `dynamodbav:"id"` // Ensure this tag matches DynamoDB's key name
	CreatedAt time.Time `dynamodbav:"createdAt"`
	UpdatedAt time.Time `dynamodbav:"updatedAt"`
	Typename  string    `dynamodbav:"typename"`
	Data      any       `dynamodbav:"data"`
}

type GithubRepo struct {
	ID                       int       `json:"id"`
	NodeID                   string    `json:"node_id"`
	Name                     string    `json:"name"`
	FullName                 string    `json:"full_name"`
	Private                  bool      `json:"private"`
	Owner                    Owner     `json:"owner"`
	HTMLURL                  string    `json:"html_url"`
	Description              any       `json:"description"`
	Fork                     bool      `json:"fork"`
	URL                      string    `json:"url"`
	ForksURL                 string    `json:"forks_url"`
	KeysURL                  string    `json:"keys_url"`
	CollaboratorsURL         string    `json:"collaborators_url"`
	TeamsURL                 string    `json:"teams_url"`
	HooksURL                 string    `json:"hooks_url"`
	IssueEventsURL           string    `json:"issue_events_url"`
	EventsURL                string    `json:"events_url"`
	AssigneesURL             string    `json:"assignees_url"`
	BranchesURL              string    `json:"branches_url"`
	TagsURL                  string    `json:"tags_url"`
	BlobsURL                 string    `json:"blobs_url"`
	GitTagsURL               string    `json:"git_tags_url"`
	GitRefsURL               string    `json:"git_refs_url"`
	TreesURL                 string    `json:"trees_url"`
	StatusesURL              string    `json:"statuses_url"`
	LanguagesURL             string    `json:"languages_url"`
	StargazersURL            string    `json:"stargazers_url"`
	ContributorsURL          string    `json:"contributors_url"`
	SubscribersURL           string    `json:"subscribers_url"`
	SubscriptionURL          string    `json:"subscription_url"`
	CommitsURL               string    `json:"commits_url"`
	GitCommitsURL            string    `json:"git_commits_url"`
	CommentsURL              string    `json:"comments_url"`
	IssueCommentURL          string    `json:"issue_comment_url"`
	ContentsURL              string    `json:"contents_url"`
	CompareURL               string    `json:"compare_url"`
	MergesURL                string    `json:"merges_url"`
	ArchiveURL               string    `json:"archive_url"`
	DownloadsURL             string    `json:"downloads_url"`
	IssuesURL                string    `json:"issues_url"`
	PullsURL                 string    `json:"pulls_url"`
	MilestonesURL            string    `json:"milestones_url"`
	NotificationsURL         string    `json:"notifications_url"`
	LabelsURL                string    `json:"labels_url"`
	ReleasesURL              string    `json:"releases_url"`
	DeploymentsURL           string    `json:"deployments_url"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	PushedAt                 time.Time `json:"pushed_at"`
	GitURL                   string    `json:"git_url"`
	SSHURL                   string    `json:"ssh_url"`
	CloneURL                 string    `json:"clone_url"`
	SvnURL                   string    `json:"svn_url"`
	Homepage                 any       `json:"homepage"`
	Size                     int       `json:"size"`
	StargazersCount          int       `json:"stargazers_count"`
	WatchersCount            int       `json:"watchers_count"`
	Language                 any       `json:"language"`
	HasIssues                bool      `json:"has_issues"`
	HasProjects              bool      `json:"has_projects"`
	HasDownloads             bool      `json:"has_downloads"`
	HasWiki                  bool      `json:"has_wiki"`
	HasPages                 bool      `json:"has_pages"`
	HasDiscussions           bool      `json:"has_discussions"`
	ForksCount               int       `json:"forks_count"`
	MirrorURL                any       `json:"mirror_url"`
	Archived                 bool      `json:"archived"`
	Disabled                 bool      `json:"disabled"`
	OpenIssuesCount          int       `json:"open_issues_count"`
	License                  any       `json:"license"`
	AllowForking             bool      `json:"allow_forking"`
	IsTemplate               bool      `json:"is_template"`
	WebCommitSignoffRequired bool      `json:"web_commit_signoff_required"`
	Topics                   []any     `json:"topics"`
	Visibility               string    `json:"visibility"`
	Forks                    int       `json:"forks"`
	OpenIssues               int       `json:"open_issues"`
	Watchers                 int       `json:"watchers"`
	DefaultBranch            string    `json:"default_branch"`
}
type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
