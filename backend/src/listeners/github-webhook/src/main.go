package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/go-github/github"
)

var svc *dynamodb.Client

const tableName = "gitpaid"

type LambdaEvent struct {
	Headers map[string]string
	Event   string
}

type DynamoIssuesEventItem struct {
	ID        string              `json:"id" dynamodbav:"id"`
	Typename  string              `json:"typename" dynamodbav:"typename"`
	CreatedAt time.Time           `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt" dynamodbav:"updatedAt"`
	Data      *github.IssuesEvent `json:"data" dynamodbav:"data"`
	Metadata  map[string]string   `json:"metadata" dynamodbav:"metadata"`
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
}

func HandleRequest(ctx context.Context, r LambdaEvent) (events.APIGatewayProxyResponse, error) {
	// Log the request object
	log.Printf("Received request:")
	payload := []byte(r.Event)
	log.Printf("Received payload: %s\n", payload)
	// Parse the event type from the headers
	eventType := r.Headers["X-GitHub-Event"]
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

func main() {
	lambda.Start(HandleRequest)
}
