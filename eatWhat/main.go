package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/csv"
	"io"
	"strings"
	"math/rand"
	"time"
)

// BodyRequest is our self-made struct to process JSON request from Client
type BodyRequest struct {
	Name string `json:"type"`
}

// BodyResponse is our self-made struct to build response for Client
type BodyResponse struct {
	Eat string `json:"eat"`
	Time time.Time `json:"time"`
}

// Handler function Using AWS Lambda Proxy Request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	data := `五十嵐,珍煮丹,可不可,麻古,迷客夏,老賴,萬波`
	r := csv.NewReader(strings.NewReader(data))

	stores := []string{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
		}

		stores = record
	}

	// We will build the BodyResponse and send it back in json form
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(len(stores))
	ans := stores[randNum]
	bodyResponse := BodyResponse{
		Eat: ans,
		Time: time.Now(),
	}

	fmt.Println(bodyResponse)

	// Marshal the response into json bytes, if error return 404
	response, err := json.Marshal(&bodyResponse)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	//Returning response with AWS Lambda Proxy Response

	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
