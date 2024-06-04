package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ninosistemas10/gambituser/models"
	"github.com/ninosistemas10/gambituser/tools"
)

func SignUp(sig models.SingUp) error {
	fmt.Println("Comienza Registro")
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()
	sentencia := "insert into users(User_Email, User_UUID, User_DateAdd) VALUE( '" + sig.UserEmail + "', '" + sig.UserUUID + "', '" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("> SignUp > Ejecucion exitosa")
	return nil

}
