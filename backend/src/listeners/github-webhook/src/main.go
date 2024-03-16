package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
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

type PaymentInfo struct {
	// closer name
	SolverUsername string `json:"solverUsername" dynamodbav:"solverUsername"`
	// closer avatar
	SolverAvatar string `json:"solverAvatar" dynamodbav:"solverAvatar"`
	// amount paid
	Amount int `json:"amount" dynamodbav:"amount"`
	// issue id - link to issues on frontend
	IssueID string `json:"issueID" dynamodbav:"issueID"`
	// tx id
	TxID string `json:"txID" dynamodbav:"txID"`
}

type User struct {
	ID             string            `json:"id" dynamodbav:"id"`
	GitHubUsername string            `json:"githubUsername" dynamodbav:"githubUsername"`
	Typename       string            `json:"typename" dynamodbav:"typename"`
	CreatedAt      time.Time         `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt" dynamodbav:"updatedAt"`
	Metadata       map[string]string `json:"metadata" dynamodbav:"metadata"`
}

type WalletLinkData struct {
	ID                string    `json:"id"`
	Chain             string    `json:"chain"`
	CreatedAt         time.Time `json:"createdAt"`
	DeletedAt         any       `json:"deletedAt"`
	HardwareWallet    any       `json:"hardwareWallet"`
	LowerPublicKey    string    `json:"lowerPublicKey"`
	Name              string    `json:"name"`
	Provider          string    `json:"provider"`
	PublicKey         string    `json:"publicKey"`
	SignerWalletID    any       `json:"signerWalletId"`
	TurnkeyHDWalletID string    `json:"turnkeyHDWalletId"`
	UpdatedAt         time.Time `json:"updatedAt"`
	UserID            string    `json:"userId"`
}

type GitHubCommit struct {
	SHA    string `json:"sha"`
	Commit struct {
		Author struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		// ... other fields as necessary
	} `json:"commit"`
	// ... other fields as necessary
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
		handlePingEvent(ctx, eventMap.Repository.ID)
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
	case *github.IssuesEvent:
		if *event.Action == "opened" || *event.Action == "closed" {
			log.Printf("Issues Event: %+v\n", event)
			putIssuesEvent(ctx, event, nil)
			// issue closed, now go and find the PR for the issue using github api
			if *event.Action == "closed" {
				log.Printf("Issue closed")

				/*
					OBTAIN LABEL AND PULL REQUEST OBJECTS
				*/
				shouldPay, paymentLabel, pullRequest, err := shouldPayUser(event)
				if err != nil {
					log.Printf("Error checking if user should be paid: %s", err)
					// return bad api response
					return events.APIGatewayProxyResponse{
						StatusCode: 500, // Internal Server Error
						Headers: map[string]string{
							"Content-Type": "application/json",
						},
						Body:            "{\"message\": \"Error checking if user should be paid\"}",
						IsBase64Encoded: false,
					}, nil
				}
				if shouldPay { // if the label is not null then we should pay the user, all the should pay checks have passed
					// pay the user
					log.Printf("User should be paid, label name is: %s", *paymentLabel.Name)

					// 1. Get the users name, then get the users record from dynamo (User object)
					// 2. Get the repo object to get the payment metadata (contract address etc)
					// 3. Pay the user by calling the contract
					/*
						OBTAIN THE REPO AND USER OBJECTS
					*/
					keyCondition := expression.Key("typename").Equal(expression.Value("Repository"))
					log.Printf("Repo name: %s", *pullRequest.Base.Repo.Name)
					filter := expression.Name("data.name").Equal(expression.Value(*pullRequest.Base.Repo.Name))

					expr, err := expression.NewBuilder().
						WithKeyCondition(keyCondition).
						WithFilter(filter).
						Build()
					if err != nil {
						return events.APIGatewayProxyResponse{
							StatusCode: 500, // Internal Server Error
							Headers: map[string]string{
								"Content-Type": "application/json",
							},
							Body:            "{\"message\": \"Error building query expression\"}",
							IsBase64Encoded: false,
						}, nil
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
					dynamoRepoItem, err := QueryDynamoDB[DynamoRepoItem](ctx, queryInput)
					log.Printf("Dynamo Repo Item: %+v", dynamoRepoItem)
					if err != nil || len(dynamoRepoItem) == 0 {
						return events.APIGatewayProxyResponse{
							StatusCode: 500, // Internal Server Error
							Headers: map[string]string{
								"Content-Type": "application/json",
							},
							Body:            "{\"message\": \"Error getting repository from dynamo\"}",
							IsBase64Encoded: false,
						}, nil
					}
					repo := dynamoRepoItem[0]
					// repoMetadata := make(map[string]string)
					log.Printf("Repo Metadata: %+v", repo.Metadata)
					paymentAddress, addressErr := getPaymentAddress(ctx, *pullRequest.User.Login)
					amount, err := parseAmount(paymentLabel.Name)
					log.Printf("Payment Address: %s", paymentAddress)
					log.Printf("Amount: %d", amount)
					if err != nil || addressErr != nil {
						log.Fatal(err)
					}
					if err != nil {
						return events.APIGatewayProxyResponse{
							StatusCode: 500, // Internal Server Error
							Headers: map[string]string{
								"Content-Type": "application/json",
							},
							Body:            "{\"message\": \"Error getting payment address\"}",
							IsBase64Encoded: false,
						}, nil
					}

					log.Printf("READY TO PAY USER ... INPUTS ARE")
					log.Printf("Repo Metadata: %+v", repo.Metadata)
					log.Printf("Payment Address: %s", paymentAddress)
					log.Printf("Amount: %d", amount)
					log.Printf("User Name: %s", *pullRequest.User.Login)
					log.Printf("Finished")
					txHash, err := payUser(ctx, repo.Metadata, paymentAddress, amount, repo.ID)
					if err != nil {
						log.Printf("Error paying user: %s", err)
						return events.APIGatewayProxyResponse{
							StatusCode: 500, // Internal Server Error
							Headers: map[string]string{
								"Content-Type": "application/json",
							},
							Body:            "{\"message\": \"Error paying user\"}",
							IsBase64Encoded: false,
						}, nil
					}
					log.Printf("Transaction Hash: %s", txHash)
					issueID := strconv.FormatInt(*event.Issue.ID, 10)
					paymentInfo := PaymentInfo{
						SolverUsername: *pullRequest.User.Login,
						SolverAvatar:   *pullRequest.User.AvatarURL,
						Amount:         amount,
						IssueID:        issueID,
						TxID:           txHash,
					}

					issue, err := getIssue(issueID)
					if err != nil {
						log.Printf("Error getting issue: %s", err)
						return events.APIGatewayProxyResponse{
							StatusCode: 500, // Internal Server Error
							Headers: map[string]string{
								"Content-Type": "application/json",
							},
							Body:            "{\"message\": \"Error getting issue\"}",
							IsBase64Encoded: false,
						}, nil
					}
					if issue.Metadata == nil {
						issue.Metadata = make(map[string]string)
					}

					issue.Metadata["solverUsername"] = paymentInfo.SolverUsername
					issue.Metadata["solverAvatar"] = paymentInfo.SolverAvatar
					issue.Metadata["amount"] = strconv.Itoa(paymentInfo.Amount)
					issue.Metadata["txID"] = paymentInfo.TxID
					putIssuesEvent(ctx, issue.Data, issue.Metadata)
				}
			}
		} else {
			if *event.Action == "labeled" || *event.Action == "unlabeled" {
				log.Printf("Running labeled or unlabeled event")
				dynamoItem, err := getIssue(strconv.FormatInt(*event.Issue.ID, 10))
				if err != nil {
					log.Printf("Error getting issue: %s", err)
					return events.APIGatewayProxyResponse{
						StatusCode: 500, // Internal Server Error
						Headers: map[string]string{
							"Content-Type": "application/json",
						},
						Body:            "{\"message\": \"Error getting issue when trying to check labeled or unlabeled\"}",
						IsBase64Encoded: false,
					}, nil
				}
				if *dynamoItem.Data.Action == "opened" {
					// issue is open, go and modify label metadata
					log.Printf("Issue is open, modifying label metadata")
					if *event.Action == "labeled" {
						if dynamoItem.Metadata == nil {
							dynamoItem.Metadata = make(map[string]string)
						}
						dynamoItem.Metadata["label"] = *event.Label.Name
					} else {
						for k := range dynamoItem.Metadata {
							if k == "label" {
								delete(dynamoItem.Metadata, k)
							}
						}
					}
				} else {
					log.Printf("Issue is not open, not modifying label metadata")
				}
				putIssuesEvent(ctx, dynamoItem.Data, dynamoItem.Metadata)
			}
		}

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

func parseAmount(label *string) (int, error) {
	if label == nil {
		return 0, fmt.Errorf("label is nil")
	}

	// Split the string on the space
	parts := strings.Split(*label, " ")

	// Check if the split resulted in at least two parts
	if len(parts) < 2 {
		return 0, fmt.Errorf("expected at least two parts, got %d", len(parts))
	}

	// Parse the first part as an integer
	amount, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("failed to parse amount: %v", err)
	}

	return amount, nil
}

func getPaymentAddress(ctx context.Context, userName string) (string, error) {
	keyCondition := expression.Key("typename").Equal(expression.Value("User"))
	filter := expression.Name("githubUsername").Equal(expression.Value(userName))

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()
	if err != nil {
		return "", fmt.Errorf("failed to build expression: %w", err)
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
	dynamoRepoItem, err := QueryDynamoDB[User](ctx, queryInput)
	if len(dynamoRepoItem) == 0 || err != nil {
		return "error obtaining wallet address for user to be paid", fmt.Errorf("failed to query DynamoDB: %w", err)
	}

	paymentAddress := &dynamoRepoItem[0]
	log.Printf("Payment Address: %+v", paymentAddress)
	if paymentAddress.Metadata == nil || paymentAddress.Metadata["publicKey"] == "" {
		return "", fmt.Errorf("no payment address found")
	}

	return dynamoRepoItem[0].Metadata["publicKey"], nil
}

func shouldPayUser(event *github.IssuesEvent) (bool, *github.Label, *github.PullRequest, error) {
	// ensure the issue has a valid payment label
	paymentLabel, err := determinePaymentLabel(event)
	if err != nil || paymentLabel == nil {
		return false, nil, nil, fmt.Errorf("error determining payment label: %s", err)
	}
	// ensure the issue closing is the result of a merge request by inspecting commit messages
	owner := *event.Repo.Owner.Login
	repo := *event.Repo.Name
	// Construct the URL with the variables
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", owner, repo)
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_ACCESS_TOKEN"))
	commits, err := fetchData[[]GitHubCommit](url, headers)
	// error fetching commits
	if err != nil {
		log.Printf("Error fetching data: %s", err)
		return false, nil, nil, err
	}
	log.Printf("Commits: %+v", commits)
	if len(commits) == 0 {
		return false, nil, nil, fmt.Errorf("no commits found")
	} else {
		log.Printf("Commits: %+v", commits)
		// look at the last commit message
		lastCommit := commits[0]
		log.Printf("Last commit: %+v", lastCommit)
		// check if the last commit message contains the issue number
		pullRequestNumber, issueSolvedNumber, err := extractNumbers(lastCommit.Commit.Message)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Pull Request Number: %d\n", pullRequestNumber)
			fmt.Printf("Issue Solved Number: %d\n", issueSolvedNumber)
		}

		// verify the pull request is actually merged and the issue is closed
		// get the pull request
		url = fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls/%d", owner, repo, pullRequestNumber)
		data, err := fetchData[github.PullRequest](url, headers)
		if err != nil {
			log.Printf("Error fetching data: %s", err)
			return false, nil, nil, err
		}
		log.Printf("PR data: %+v", data)
		if data.MergedAt == nil {
			return false, nil, nil, fmt.Errorf("PR not merged")
		}
		// get the issue
		url = fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", owner, repo, issueSolvedNumber)
		issue, err := fetchData[github.Issue](url, headers)
		if err != nil {
			log.Printf("Error fetching data: %s", err)
			return false, nil, nil, err
		}
		log.Printf("Issue data: %+v", issue)
		if issue.State != nil && *issue.State == "closed" {
			return true, paymentLabel, &data, nil
		} else {
			return false, nil, nil, fmt.Errorf("issue not closed")
		}
	}
}

func extractNumbers(commitMessage string) (pullRequestNumber int, issueSolvedNumber int, err error) {
	// Regex to match "Merge pull request #<number>"
	prRegex := regexp.MustCompile(`Merge pull request #(\d+)`)
	prMatches := prRegex.FindStringSubmatch(commitMessage)

	// Regex to match "fixes #<number>"
	fixesRegex := regexp.MustCompile(`fixes #(\d+)`)
	fixesMatches := fixesRegex.FindStringSubmatch(commitMessage)

	// If matches found, convert captured strings to integers
	if len(prMatches) > 1 {
		pullRequestNumber, err = strconv.Atoi(prMatches[1])
		if err != nil {
			return 0, 0, fmt.Errorf("failed to convert pull request number to integer: %v", err)
		}
	} else {
		err = fmt.Errorf("no pull request number found")
		return
	}

	if len(fixesMatches) > 1 {
		issueSolvedNumber, err = strconv.Atoi(fixesMatches[1])
		if err != nil {
			return 0, 0, fmt.Errorf("failed to convert issue solved number to integer: %v", err)
		}
	} else {
		err = fmt.Errorf("no issue solved number found")
		return
	}

	return pullRequestNumber, issueSolvedNumber, nil
}

func getValidLabelsForRepo(repo *string) ([]DynamoLabelEventItem, error) {
	// query the dynamodb table for all the labels for the repo
	// get the labels that are valid for the repo
	// return the labels
	keyCondition := expression.Key("typename").Equal(expression.Value("Label"))
	filter := expression.Name("data.Repo.Name").Equal(expression.Value(repo)).And(expression.Name("data.Action").Equal(expression.Value("created")))

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build expression: %w", err)
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
	ctx := context.TODO()
	items, err := QueryDynamoDB[DynamoLabelEventItem](ctx, queryInput)
	if err != nil {
		return nil, fmt.Errorf("failed to query DynamoDB: %w", err)
	}
	return items, nil
}

func determinePaymentLabel(event *github.IssuesEvent) (*github.Label, error) {
	verifiedLabels, err := getValidLabelsForRepo(event.Repo.Name)
	if err != nil {
		return nil, fmt.Errorf("error getting valid labels for repo: %w", err)
	}
	// iterate through the labels on the event

	for _, eventLabel := range event.Issue.Labels {
		log.Printf("Checking label: %s", *eventLabel.Name)
		for _, verifiedLabel := range verifiedLabels {
			log.Printf("Checking verified label: %s", *verifiedLabel.Data.Label.Name)
			// found a verified label, return it
			_, correctValue := findMatchingString(eventLabel.Name)
			log.Printf("Correct value: %t", correctValue)
			log.Printf("Verified label: %+v", verifiedLabel)
			log.Printf("Event label: %+v", eventLabel)
			if *verifiedLabel.Data.Label.Name == *eventLabel.Name && correctValue {
				return &eventLabel, nil
			}
		}
	}
	return nil, fmt.Errorf("no verified label found")
}

func findMatchingString(input *string) (string, bool) {
	// Create a regex pattern
	pattern := regexp.MustCompile(`^\d+\s[a-zA-Z]{1,5}$`)

	// Find a match
	match := pattern.FindString(*input)

	// If match is found, return it along with true; otherwise return an empty string with false
	return match, match != ""
}

func fetchData[T any](url string, headers map[string]string) (T, error) {
	var result T // Initialize a variable of type T

	// Create a new HTTP GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}

	// Set headers for the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// Unmarshal the response body into the result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func getIssue(ID string) (DynamoIssuesEventItem, error) {
	keyCondition := expression.Key("typename").Equal(expression.Value("Issue"))
	filter := expression.Name("id").Equal(expression.Value(ID))

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCondition).
		WithFilter(filter).
		Build()
	if err != nil {
		return DynamoIssuesEventItem{}, fmt.Errorf("failed to build expression: %w", err)
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
	ctx := context.TODO()
	items, err := QueryDynamoDB[DynamoIssuesEventItem](ctx, queryInput)
	if err != nil {
		return DynamoIssuesEventItem{}, fmt.Errorf("failed to query DynamoDB: %w", err)
	}
	// issue := items[0].Data.(*github.IssuesEvent)
	return items[0], nil
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

func handlePingEvent(ctx context.Context, repoID int) (string, error) {
	keyCondition := expression.Key("typename").Equal(expression.Value("Repository"))
	filter := expression.Name("data.id").Equal(expression.Value(repoID))

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
