package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
