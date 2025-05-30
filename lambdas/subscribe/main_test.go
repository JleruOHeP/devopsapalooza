package main

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
)

// MockDynamoDB is a mock implementation of DynamoDBAPI
type MockDynamoDB struct {
	PutItemFunc func(ctx context.Context, input *dynamodb.PutItemInput, opts ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

func (m *MockDynamoDB) PutItem(ctx context.Context, input *dynamodb.PutItemInput, opts ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return m.PutItemFunc(ctx, input, opts...)
}

func TestHandler_ValidEmail(t *testing.T) {
	mock := &MockDynamoDB{
		PutItemFunc: func(ctx context.Context, input *dynamodb.PutItemInput, opts ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
			assert.Equal(t, "user@example.com", input.Item["Email"].(*types.AttributeValueMemberS).Value)
			_, err := time.Parse(time.RFC3339, input.Item["Date"].(*types.AttributeValueMemberS).Value)
			assert.NoError(t, err)
			return &dynamodb.PutItemOutput{}, nil
		},
	}

	ddbClient = mock

	response, err := handler(context.Background(), "user@example.com")

	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Email and date saved", response.Body)
}

func TestHandler_InvalidEmail(t *testing.T) {
	ddbClient = &MockDynamoDB{} // it won't be called

	response, err := handler(context.Background(), "not-email")

	assert.NoError(t, err)
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, "Invalid request body", response.Body)
}

func TestHandler_DynamoError(t *testing.T) {
	mock := &MockDynamoDB{
		PutItemFunc: func(ctx context.Context, input *dynamodb.PutItemInput, opts ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
			return nil, errors.New("DynamoDB error")
		},
	}

	ddbClient = mock

	response, err := handler(context.Background(), "user@example.com")

	assert.NoError(t, err)
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, "Failed to save email", response.Body)
}
