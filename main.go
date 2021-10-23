package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string	 `json:"name"`
	Age	 int	 `json:"age"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Hello World, my name is %s, i am %d years old!!", request.Name, request.Age),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

