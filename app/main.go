package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func() (string, error) {
		fmt.Println("Hello, Lambda!!")
		return "Hello, Golang!!", nil
	})
}
