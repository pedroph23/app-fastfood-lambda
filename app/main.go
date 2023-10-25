package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		resp := gin.H{
			"statusCode": http.StatusOK,
			"headers":    map[string]string{"Content-Type": "application/json"},
			"body":       "Hello, World!",
		}
		ctx.JSON(http.StatusOK, resp)
	})

	// Converte a solicitação do API Gateway em uma solicitação HTTP do Gin
	req, err := gin.RequestFromAPIGatewayProxyRequest(&request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	// Executa a solicitação no roteador Gin
	r.ServeHTTP(req.ResponseWriter, req)

	// Converte a resposta do Gin em uma resposta do API Gateway
	resp, err := gin.APIGatewayProxyResponseFromHTTPRequest(req.Request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
