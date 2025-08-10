package aws

import "context"

type DynamoDBClient interface {
	Scan(ctx context.Context, table string, filterExpression string, exprValues map[string]any, resultsPtr any) error
}
