package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func GetCORSHeaders() map[string]string {
	return map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type, Authorization, x-api-key",
	}
}

func CreateErrorResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	headers := GetCORSHeaders()
	body, _ := json.Marshal(map[string]string{"error": message})
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       string(body),
	}
}

func CreateEmptyResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 204,
		Headers:    GetCORSHeaders(),
	}
}

func CreateSuccessResponse(data any) (events.APIGatewayProxyResponse, error) {
	headers := GetCORSHeaders()
	body, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    headers,
		Body:       string(body),
	}, nil
}
