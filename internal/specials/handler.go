package specials

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"

	"github.com/nelskin/go-derma-practice/internal/types"
	"github.com/nelskin/go-derma-practice/internal/utils"
)

type Repository interface {
	ListActiveSpecials(ctx context.Context) ([]types.Special, error)
}

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	items, err := h.repo.ListActiveSpecials(ctx)
	if err != nil {
		log.Printf("service error: %v", err)
		return utils.CreateErrorResponse(500, "Internal server error"), nil
	}

	if len(items) == 0 {
		return utils.CreateEmptyResponse(), nil
	}
	return utils.CreateSuccessResponse(types.SpecialsResponse{Specials: items})
}
