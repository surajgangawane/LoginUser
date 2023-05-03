package login

import (
	"LoginUser/models"
	"LoginUser/repository"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type loginUser struct {
	dbClient repository.Repository
}

func (lu loginUser) LoginUser(rw http.ResponseWriter, r *http.Request) {
	var (
		loginRequest  models.LoginRequest
		errorString   string
		loginResponse models.LoginResponse
	)

	ctx := context.Background()

	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error while reading login request ", err)
		errorString = fmt.Sprintf("Bad Request: %s", err.Error())
		rw.WriteHeader(401)
		rw.Write([]byte(errorString))
	}

	err = json.Unmarshal(requestByte, &loginRequest)
	if err != nil {
		log.Fatal("Error while unmarshaling login request ", err)
		errorString = fmt.Sprintf("Bad Request: %s", err.Error())
		rw.WriteHeader(401)
		rw.Write([]byte(errorString))
	}

	userDetails, err := lu.dbClient.GetUserDetails(ctx, loginRequest.UserName)
	if err != nil {
		log.Fatal("Error while login user", err)
		errorString = fmt.Sprintf("Error while login user %s, please try again !", loginRequest.UserName)
		rw.WriteHeader(500)
		rw.Write([]byte(errorString))
	}

	if userDetails == (models.UserData{}) {
		log.Fatal("User is not registered")
		errorString = fmt.Sprintf("User %s is not registered. Please register", loginRequest.UserName)
		rw.WriteHeader(200)
		rw.Write([]byte(errorString))
	}

	if userDetails.Password != loginRequest.Password {
		log.Fatal("Invalid password")
		errorString = fmt.Sprintf("Invalid Passoword. %s please enter correct password", loginRequest.UserName)
		rw.WriteHeader(200)
		rw.Write([]byte(errorString))
	}

	token := lu.createToken(ctx, userDetails)
	loginResponse.Token = token
	loginResponse.User.UserName = userDetails.UserName
	loginResponse.User.FirstName = userDetails.FirstName
	loginResponse.User.LastName = userDetails.LastName
	loginResponse.User.Email = userDetails.EmailId

	resByte, err := json.Marshal(loginResponse)
	if err != nil {
		log.Fatal("Error while marshalling login response", err)
		errorString = fmt.Sprintf("Error while login user %s, please try again !", loginRequest.UserName)
		rw.WriteHeader(500)
		rw.Write([]byte(errorString))
	}

	rw.WriteHeader(200)
	rw.Write(resByte)

}

func (lu loginUser) createToken(ctx context.Context, userDetails models.UserData) string {
	claims := models.JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(10)).Unix(),
		},
		CustomClaims: map[string]string{
			"user_name":  userDetails.UserName,
			"first_name": userDetails.FirstName,
			"last_name":  userDetails.LastName,
			"email":      userDetails.EmailId,
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenString.SignedString(tokenString)
	return token
}

func NewLoginUser(dbClient repository.Repository) LoginUser {
	return loginUser{
		dbClient: dbClient,
	}
}
