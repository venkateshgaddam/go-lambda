package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"What is your Name?"`
	Age  int32  `json:"Your Age"`
}

type MyResponse struct {
	Message string `json:"Answer"`
}

func Handler(event MyEvent) (MyResponse, error) {
	result := fmt.Sprintf("%s is %d years old!", event.Name, event.Age)
	fmt.Println(result)
	return MyResponse{Message: result}, nil
}

func main() {

	fmt.Println("Lambda Function is working!!!")
	lambda.Start(Handler)
}
