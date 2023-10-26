package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{Message: "Hello World"})
	})

	respBody, err := json.Marshal(r.Run())
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to marshal response body: %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
