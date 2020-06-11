package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/csv"
	"io"
	"strings"
	"math/rand"
	"time"
)

type BodyResponse struct {
	Eat string `json:"eat"`
	Time time.Time `json:"time"`
}

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


	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(len(stores))
	ans := stores[randNum]
	bodyResponse := BodyResponse{
		Eat: ans,
		Time: time.Now(),
	}

	response, err := json.Marshal(&bodyResponse)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
