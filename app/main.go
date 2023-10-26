package main

import (
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gin-gonic/gin"
)

func LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Crie uma instância do servidor Gin
	r := gin.Default()

	// Defina uma rota Gin que corresponda à rota da sua API Gateway
	r.GET("/hello", func(c *gin.Context) {
		// Extraia os parâmetros da solicitação da API Gateway
		name := c.DefaultQuery("name", "World")

		// Execute sua lógica de negócios
		message := "Hello, " + name

		// Responda com a mensagem
		c.JSON(200, gin.H{"message": message})
	})

	// Crie um contexto Gin manualmente e configure-o com a solicitação da API Gateway
	ctx := &gin.Context{
		Request: &http.Request{
			Method: "GET", // Defina o método HTTP da solicitação
			URL: &url.URL{
				Path: "/hello", // Defina o caminho da solicitação
			},
			Header: make(http.Header),
		},
		Writer: c.Writer, // Use o gravador de resposta do Gin
	}

	// Execute a solicitação no servidor Gin
	r.ServeHTTP(ctx.Writer, ctx.Request)

	// Converta a resposta de Gin em uma resposta da API Gateway
	response := events.APIGatewayProxyResponse{
		StatusCode:      ctx.Writer.Status(),
		Headers:         ctx.Writer.Header(),
		Body:            ctx.Writer.Body.String(),
		IsBase64Encoded: false,
	}

	return response, nil
}
