package main

import (
	"context"
	"fmt"
	"net/mail"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	ddbClient DynamoDBAPI
	tableName = "Customers"
)

type DynamoDBAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

func handler(ctx context.Context, request string) (events.APIGatewayProxyResponse, error) {
	_, err := mail.ParseAddress(request)
	if err != nil {
		fmt.Println(request, "is invalid email:", err)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request body"}, nil
	}

	currentTime := time.Now().UTC().Format(time.RFC3339)

	_, err = ddbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"Email":   &types.AttributeValueMemberS{Value: request},
			"Date":    &types.AttributeValueMemberS{Value: currentTime},
			"Message": &types.AttributeValueMemberS{Value: "Source: Go Go Go"},
		},
	})

	if err != nil {
		fmt.Println("Failed to write to DynamoDB:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to save email"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Email and date saved"}, nil
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	ddbClient = dynamodb.NewFromConfig(cfg)
	lambda.Start(handler)
}
