package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ecommerceUser/awsgo"
	"github.com/ecommerceUser/db"
	"github.com/ecommerceUser/models"
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

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			fmt.Println("Email = " + att)
			data.UserEmail = att
		case "sub":
			fmt.Println("UserUUID = " + att)
			data.UserUUID = att
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error reading the secret: " + err.Error())
		return event, err
	}
	return event, nil
}

func ValidateParameters() bool {
	var hasParameter bool
	_, hasParameter = os.LookupEnv("SecretName")
	return hasParameter
}
