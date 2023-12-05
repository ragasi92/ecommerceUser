package db

import (
	"fmt"

	"github.com/ecommerceUser/models"
	"github.com/ecommerceUser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Start user signup")
	err := MysqlDBconnect()
	if err != nil {
		return err
	}
	defer DB.Close()

	sqlSentene := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.MysqlNowDate() + "')"
	fmt.Println(sqlSentene)
	_, err = DB.Exec(sqlSentene)
	if err != nil {
		return err
	}

	fmt.Println("SignUp > Succesfuly!")
	return nil
}
