package specials

import (
	"context"
	"time"

	awsclient "github.com/nelskin/go-derma-practice/internal/aws"
	"github.com/nelskin/go-derma-practice/internal/types"
)

type DynamoDBRepository struct {
	client awsclient.DynamoDBClient
	table  string
}

func NewDynamoDBRepository(client awsclient.DynamoDBClient, table string) *DynamoDBRepository {
	return &DynamoDBRepository{client: client, table: table}
}

func (r *DynamoDBRepository) ListActiveSpecials(ctx context.Context) ([]types.Special, error) {
	var items []types.Special
	filter := "Status = :s AND ActiveFrom <= :now AND ActiveUntil >= :now"
	nowStr := time.Now().UTC().Format(time.RFC3339)

	values := map[string]any{":s": "active", ":now": nowStr}
	if err := r.client.Scan(ctx, r.table, filter, values, &items); err != nil {
		return nil, err
	}
	return items, nil
}
