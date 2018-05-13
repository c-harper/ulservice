package main

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type userStruct struct {
	Data []struct {
		Name           string `json:"name"`
		FavoriteColor  string `json:"favoriteColor"`
		FavoriteAnimal string `json:"favoriteAnimal"`
	} `json:"data"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var users userStruct
	err := json.Unmarshal([]byte(request.Body), &users)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "error in request body", StatusCode: 200}, err
	}

	var b strings.Builder
	for _, user := range users.Data {
		b.WriteString(user.Name + "\n")
	}
	return events.APIGatewayProxyResponse{Body: b.String(), StatusCode: 200}, nil

}

func main() {
	lambda.Start(handler)
}
