package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/",
	}

	resp, err := Handler(context.Background(), req)
	assert.NoError(t, err)

	var respBody Response
	err = json.Unmarshal([]byte(resp.Body), &respBody)
	assert.NoError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, respBody.Message, "Hello World")
}
