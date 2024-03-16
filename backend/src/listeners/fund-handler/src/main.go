package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

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

// type LambdaInput struct {
// 	Event                 any    `json:"event"`
// 	Network               any    `json:"network"`
// 	ListenerId            any    `json:"listenerId"`
// 	Abi                   any    `json:"abi"`
// 	Api                   any    `json:"api"`
// 	Args                  any    `json:"args"`
// 	Headers               any    `json:"headers"`
// 	Method                string `json:"method"`
// 	QueryStringParameters any    `json:"queryStringParameters"`
// }

func init() {
	Init()
}

type LambdaEvent struct {
	Headers map[string]string
	Event   UnpackedEvent
}

type UnpackedEvent struct {
	Args []string `json:"args"`
}

func logAsJSON(v interface{}) {
	// Marshal the interface{} to JSON with indentation for pretty-printing
	bytes, err := json.MarshalIndent(v, "", "    ") // Use 4 spaces for indentation; you can also use "\t" for tabs
	if err != nil {
		log.Printf("Error serializing to JSON: %s", err)
		return
	}

	// Convert bytes to string and log
	log.Println(string(bytes))
}

type EventStruct struct {
	Args []string `json:"args"`
}

func HandleRequest(ctx context.Context, r LambdaEvent) {

	// var eventMap UnpackedEvent
	// if err := json.Unmarshal(payload, &eventMap); err != nil {
	// 	log.Fatalf("Failed to unmarshal Event string: %s", err)
	// }

	// print out the object
	// fmt.Println("Lambda Input: ", lambdaInput)
	// args := json.Unmarshal(lambdaInput.Event, &EventStruct{})
	// logAsJSON(args)
	tokenAddress := r.Event.Args[0]
	repository := r.Event.Args[1]
	payeeAddress := r.Event.Args[2]
	amount := r.Event.Args[3]

	// fmt.Println("Token Address: ", tokenAddress)
	// fmt.Println("Repository: ", repository)
	// fmt.Println("Payee Address: ", payeeAddress)
	// fmt.Println("Amount: ", amount)

	// // convert amount into a uint64
	// amountUint64, err := strconv.ParseUint(amount, 10, 64)
	// if err != nil {
	// 	fmt.Println("Error converting amount to uint64: ", err)
	// 	return
	// }

	keyCond := expression.KeyEqual(expression.Key("typename"), expression.Value("Repository"))
	filter := expression.Name("id").Equal(expression.Value(repository))

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithFilter(filter).
		Build()

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return
	}

	queryInput := dynamodb.QueryInput{
		TableName:                 aws.String("gitpaid"),
		IndexName:                 aws.String("byTypename"),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	repos, err := QueryDynamoDB[DynamoItem](ctx, queryInput)
	if err != nil {
		fmt.Println("Error querying DynamoDB: ", err)
		return
	}

	if len(repos) == 0 {
		fmt.Println("Repository not found ", repository)
		return
	}

	newMap := TokenMetadata{
		"421614",
		"https://arb-sepolia.g.alchemy.com/v2/Z8Y0CZXvhPgiTt8akdr4Z_dS03C2-H0X",
		tokenAddress,
		payeeAddress,
		amount,
	}
	marshalMap, err := attributevalue.MarshalMap(newMap)
	if err != nil {
		fmt.Println("Error marshalling map: ", err)
		return
	}

	// newMap := map[string]types.AttributeValue{
	// 	"payeeAddress": &types.AttributeValueMemberS{Value: payeeAddress},
	// 	"amount":       &types.AttributeValueMemberS{Value: amount},
	// 	"tokenAddress": &types.AttributeValueMemberS{Value: tokenAddress},
	// 	"chainID":      &types.AttributeValueMemberS{Value: "421614"},
	// 	"rpc":          &types.AttributeValueMemberS{Value: "https://arb-sepolia.g.alchemy.com/v2/Z8Y0CZXvhPgiTt8akdr4Z_dS03C2-H0X"},
	// }

	updates := map[string]types.AttributeValue{
		"updatedAt": &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
		"metadata":  &types.AttributeValueMemberM{Value: marshalMap},
	}

	err = UpdateItem(ctx, "gitpaid", repository, updates)
	if err != nil {
		fmt.Println("Error updating item: ", err)
		return
	}
}

func main() {
	lambda.Start(HandleRequest)
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

type TokenMetadata struct {
	ChainID      string `json:"chainID" dynamodbav:"chainID"`
	RPC          string `json:"rpc" dynamodbav:"rpc"`
	TokenAddress string `json:"tokenAddress" dynamodbav:"tokenAddress"`
	PayeeAddress string `json:"payeeAddress" dynamodbav:"payeeAddress"`
	Amount       string `json:"amount" dynamodbav:"amount"`
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

type DynamoItem struct {
	ID        string    `dynamodbav:"id"`
	CreatedAt time.Time `dynamodbav:"createdAt"`
	UpdatedAt time.Time `dynamodbav:"updatedAt"`
	Typename  string    `dynamodbav:"typename"`
	Data      any       `dynamodbav:"data"`
}
