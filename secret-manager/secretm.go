package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/ecommerceUser/awsgo"
	"github.com/ecommerceUser/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var data models.SecretRDSJson
	fmt.Println("> getting secret: " + secretName)

	smClient := secretsmanager.NewFromConfig(awsgo.Cfg)

	clave, err := smClient.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}

	err = json.Unmarshal([]byte(*clave.SecretString), &data)
	if err != nil {
		fmt.Println("Error unmarshal secret: " + err.Error())
		return data, err
	}

	return data, nil

}
