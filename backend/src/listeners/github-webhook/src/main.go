package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/github"
)

type LambdaEvent struct {
	Headers map[string]string
	Event   string
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
	case *github.LabelEvent:
		log.Printf("Label Event: %+v\n", event)
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

func main() {
	lambda.Start(HandleRequest)
}
