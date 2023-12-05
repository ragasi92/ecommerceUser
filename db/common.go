package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ecommerceUser/models"
	secretmanager "github.com/ecommerceUser/secret-manager"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var DB *sql.DB

func ReadSecret() error {
	var err error
	SecretModel, err = secretmanager.GetSecret(os.Getenv("SecretName"))
	return err
}

func MysqlDBconnect() error {
	fmt.Println("Connecting to MySql DB")
	var err error
	DB, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("DB connected succesfully")
	return nil
}

func ConnStr(key models.SecretRDSJson) string {
	var dbuser, authtoken, dbEndpoint, dbName string
	dbuser = key.Username
	authtoken = key.Password
	dbEndpoint = key.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbuser, authtoken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
