package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	"github.com/google/go-github/github"
)

var svc *dynamodb.Client

const tableName = "gitpaid"

type LambdaEvent struct {
	Headers map[string]string
	Event   string
}

type GithubZen struct {
	Repository GithubRepo `json:"repository"`
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

type DynamoIssuesEventItem struct {
	ID        string              `json:"id" dynamodbav:"id"`
	Typename  string              `json:"typename" dynamodbav:"typename"`
	CreatedAt time.Time           `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt" dynamodbav:"updatedAt"`
	Data      *github.IssuesEvent `json:"data" dynamodbav:"data"`
	Metadata  map[string]string   `json:"metadata" dynamodbav:"metadata"`
}

type DynamoRepoItem struct {
	ID        string            `json:"id" dynamodbav:"id"`
	Typename  string            `json:"typename" dynamodbav:"typename"`
	CreatedAt time.Time         `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt" dynamodbav:"updatedAt"`
	Data      GithubRepo        `json:"data" dynamodbav:"data"`
	Metadata  map[string]string `json:"metadata" dynamodbav:"metadata"`
}

type DynamoLabelEventItem struct {
	ID        string             `json:"id" dynamodbav:"id"`
	Typename  string             `json:"typename" dynamodbav:"typename"`
	CreatedAt time.Time          `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" dynamodbav:"updatedAt"`
	Data      *github.LabelEvent `json:"data" dynamodbav:"data"`
	Metadata  map[string]string  `json:"metadata" dynamodbav:"metadata"`
}

func init() {
	accessKey := os.Getenv("HACKATHON_ACCESS_KEY")
	secretKey := os.Getenv("HACKATHON_SECRET_KEY")
	if accessKey == "" || secretKey == "" {
		err := fmt.Errorf("access Key and secret key not found")
		log.Fatal(err)
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
		log.Printf("Error serializing to JSON: %s", err)
		return
	}

	// Create a new DynamoDB client
	svc = dynamodb.NewFromConfig(cfg)
}

func HandleRequest(ctx context.Context, r LambdaEvent) (events.APIGatewayProxyResponse, error) {
	// Log the request object
	log.Printf("Received request:")
	payload := []byte(r.Event)

	log.Printf("Received payload: %s\n", payload)
	// Parse the event type from the headers
	eventType := r.Headers["X-GitHub-Event"]

	if eventType == "ping" {
		var eventMap GithubZen
		if err := json.Unmarshal(payload, &eventMap); err != nil {
			log.Fatalf("Failed to unmarshal Event string: %s", err)
		}
		handlePingEvent(ctx, eventMap.Repository.Name)
		return events.APIGatewayProxyResponse{
			StatusCode: 200, // OK
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body:            "{\"message\": \"All Good!\"}",
			IsBase64Encoded: false,
		}, nil

	}

	log.Print("Received event type from map: " + eventType)
	// Parse the event
	eventObject, err := github.ParseWebHook(eventType, payload)
	if err != nil {
		log.Printf("error parsing webhook: %+v\n", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500, // Internal Server Error
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body:            "{\"message\": \"Error parsing webhook\"}",
			IsBase64Encoded: false,
		}, nil
	}

	switch event := eventObject.(type) {
	case *github.PingEvent:
		log.Printf("Received ping event: %+v\n", event)
		// handlePingEvent(ctx, event.)
	case *github.IssuesEvent:
		log.Printf("Received issue event: %+v\n", event)
		putIssuesEvent(ctx, event, r.Headers)

	case *github.LabelEvent:
		log.Printf("Label Event: %+v\n", event)
		putLabelEvent(ctx, event)

	default:
		log.Printf("Unknown event type: %+v\n", event)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200, // OK
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:            "{\"message\": \"All Good!\"}",
		IsBase64Encoded: false,
	}, nil
}

func putLabelEvent(ctx context.Context, event *github.LabelEvent) {
	item := DynamoLabelEventItem{
		ID:        strconv.FormatInt(*event.Label.ID, 10),
		Typename:  "Label",
		CreatedAt: time.Now().UTC(),
		Data:      event,
	}
	// check if item exists
	PutItemInDynamoDB(ctx, item)
}

func putIssuesEvent(ctx context.Context, event *github.IssuesEvent, metaData map[string]string) {
	item := DynamoIssuesEventItem{
		ID:        strconv.FormatInt(*event.Issue.ID, 10),
		Typename:  "Issue",
		CreatedAt: time.Unix(0, event.Issue.UpdatedAt.UnixNano()).UTC(),
		Data:      event,
		Metadata:  metaData,
	}

	// check if item exists
	PutItemInDynamoDB(ctx, item)
}

func PutItemInDynamoDB[T any](ctx context.Context, item T) error {
	// Marshal the item into a map of DynamoDB attribute values
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("failed to marshal item into DynamoDB attribute values: %w", err)
	}

	// Put the item into the specified DynamoDB table
	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	if err != nil {
		return fmt.Errorf("failed to put item into DynamoDB table: %w", err)
	}

	return nil
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

func handlePingEvent(ctx context.Context, repoName string) (string, error) {
	keyCondition := expression.Key("typename").Equal(expression.Value("Repository"))
	filter := expression.Name("data.name").Equal(expression.Value(repoName))

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()
	if err != nil {
		return "error", fmt.Errorf("failed to build expression: %w", err)
	}
	// Prepare the query input
	queryInput := dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		IndexName:                 aws.String("byTypename"),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}
	// when we receive a ping, we should add metadata flag all repos in dynamo as active
	repo, error := QueryDynamoDB[DynamoRepoItem](ctx, queryInput)
	if error != nil || len(repo) == 0 {
		return "error", fmt.Errorf("failed to query DynamoDB: %w", err)
	}
	_repo := repo[0]
	if _repo.Metadata == nil {
		_repo.Metadata = make(map[string]string)
	}
	_repo.Metadata["active"] = "true"
	PutItemInDynamoDB(ctx, _repo)

	return "success", nil
}

func main() {
	lambda.Start(HandleRequest)
}
