package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/ninosistemas10/gambituser/awsgo"
	"github.com/ninosistemas10/gambituser/db"
	"github.com/ninosistemas10/gambituser/models"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAws()
	if !ValidoParameter() {
		fmt.Println("Error en los parametros debe enviar  SecretName")
		err := errors.New("Error en los parametros debe enviar SecretName")
		return event, err
	}

	var datos models.SingUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error al aleer  el Secret" + err.Error())
		return event, err
	}

	err = db.SignUp(datos)

	return event, err
}

func ValidoParameter() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
