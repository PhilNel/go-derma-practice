package aws

import (
	"context"
	"errors"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/nelskin/go-derma-practice/internal/config"
)

type dynamoDBClient struct {
	client *dynamodb.Client
}

func NewDynamoDBClient(dcfg *config.DynamoDB) (DynamoDBClient, error) {
	if dcfg == nil {
		return nil, errors.New("nil dynamodb config")
	}

	cfg, err := awsconfig.LoadDefaultConfig(context.Background(), awsconfig.WithRegion(dcfg.Region))
	if err != nil {
		return nil, fmt.Errorf("load aws config: %w", err)
	}
	return &dynamoDBClient{client: dynamodb.NewFromConfig(cfg)}, nil
}

func (d *dynamoDBClient) Scan(ctx context.Context, table string, filterExpression string, exprValues map[string]any, resultsPtr any) error {
	var ev map[string]types.AttributeValue
	var err error
	if exprValues != nil {
		ev, err = attributevalue.MarshalMap(exprValues)
		if err != nil {
			return fmt.Errorf("marshal expr values: %w", err)
		}
	}

	input := &dynamodb.ScanInput{TableName: &table}
	if filterExpression != "" {
		input.FilterExpression = &filterExpression
		input.ExpressionAttributeValues = ev
	}

	out, err := d.client.Scan(ctx, input)
	if err != nil {
		return err
	}
	return attributevalue.UnmarshalListOfMaps(out.Items, resultsPtr)
}
