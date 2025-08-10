package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	awsclient "github.com/nelskin/go-derma-practice/internal/aws"
	"github.com/nelskin/go-derma-practice/internal/config"
	"github.com/nelskin/go-derma-practice/internal/specials"
	"github.com/nelskin/go-derma-practice/internal/utils"
)

var (
	handler *specials.Handler
)

func init() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dynamoClient, err := awsclient.NewDynamoDBClient(cfg.DynamoDB)
	if err != nil {
		log.Fatalf("failed to create dynamodb client: %v", err)
	}

	repo := specials.NewDynamoDBRepository(dynamoClient, cfg.DynamoDB.SpecialsTable)
	handler = specials.NewHandler(repo)
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic recovered: %v", r)
		}
	}()

	resp, err := handler.Handle(ctx, req)
	if err != nil {
		log.Printf("handler error: %v", err)
		return utils.CreateErrorResponse(500, "Internal server error"), nil
	}
	return resp, nil
}

func main() {
	lambda.Start(handleRequest)
}
