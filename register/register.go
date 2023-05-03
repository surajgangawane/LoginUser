package register

import (
	"LoginUser/models"
	"LoginUser/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"context"
)

type register struct {
	dbClient repository.Repository
}

func (reg register) Register(rw http.ResponseWriter, r *http.Request) {
	var (
		registerRequest models.RegisterRequest
		errorString     string
	)

	var ctx = context.Background()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error while reading request ", err)
		errorString = fmt.Sprintf("Bad Request: %s", err.Error())
		rw.WriteHeader(401)
		rw.Write([]byte(errorString))
	}

	err = json.Unmarshal(requestByte, &registerRequest)
	if err != nil {
		log.Fatal("Error while unmarshaling request ", err)
		errorString = fmt.Sprintf("Bad Request: %s", err.Error())
		rw.WriteHeader(401)
		rw.Write([]byte(errorString))
	}

	userExists, err := reg.dbClient.UserAlreadyRegistered(ctx, registerRequest.UserName)
	if err != nil {
		log.Fatal("Error while registering user", err)
		errorString = fmt.Sprintf("Failed to register, please try again !")
		rw.WriteHeader(500)
		rw.Write([]byte(errorString))
	}

	if userExists {
		log.Fatal("User is already registered")
		errorString = fmt.Sprintf("User %s is already registered", registerRequest.UserName)
		rw.WriteHeader(200)
		rw.Write([]byte(errorString))
	}

	err = reg.dbClient.RegisterNewUser(ctx, registerRequest)
	if err != nil {
		log.Fatal("Failed to register new user in db")
		errorString = fmt.Sprintf("Failed to register, please try again")
		rw.WriteHeader(500)
		rw.Write([]byte(errorString))
	}

	successMessage := "A verification mail has been sent to your registered mail"
	rw.WriteHeader(200)
	rw.Write([]byte(successMessage))

}

func NewRegister(dbClient repository.Repository) RegisterInt {
	return register{
		dbClient: dbClient,
	}
}
