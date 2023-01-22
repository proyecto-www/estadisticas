package main

import (
	// add this
	"encoding/json"
	"errors"
	"fmt"
	database "hello-world/adapters"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	var query = database.GetAllVehicles()
	response, err := json.Marshal(query)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf(string(response)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
