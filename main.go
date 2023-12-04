package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ecommerceUser/awsgo"
)

func main() {

	lambda.Start(RunLambda)
}

func RunLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitAws()
	if !ValidateParameters() {
		fmt.Println("Parameters error: Missing Secret Name")
		err := errors.New("Missing Secret Name in parameters")
		return event, err
	}
	return events.CognitoEventUserPoolsPostConfirmation{}, nil
}

func ValidateParameters() bool {
	var hasParameter bool
	_, hasParameter = os.LookupEnv("SecretName")
	return hasParameter
}
