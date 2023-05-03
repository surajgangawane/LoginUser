package routers

import (
	"LoginUser/login"
	"LoginUser/register"

	"github.com/gorilla/mux"
)

type Route struct {
	Register register.RegisterInt
	Login    login.LoginUser
}

func (r Route) Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user/register", r.Register.Register).Methods("POST")
	router.HandleFunc("/user/register", r.Login.LoginUser).Methods("POST")
	return router
}
