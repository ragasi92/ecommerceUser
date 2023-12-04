package db

import (
	"os"

	"github.com/ecommerceUser/models"
	secretmanager "github.com/ecommerceUser/secret-manager"
)

var SecretModel models.SecretRDSJson

func ReadSecret() error {
	var err error
	SecretModel, err = secretmanager.GetSecret(os.Getenv("SecretName"))
	return err
}
