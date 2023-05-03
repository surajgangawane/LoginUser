package main

import (
	"LoginUser/config"
	"LoginUser/login"
	"LoginUser/register"
	"LoginUser/repository"
	"LoginUser/routers.go"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadConfig()
	appConfig := config.GetAppConfig()
	conn, err := getMySqlDbConnection(appConfig)
	if err != nil {
		log.Fatal("Failed to connect db", err)
		panic(err)
	}

	conn.SetMaxOpenConns(1)
	conn.SetMaxIdleConns(2)
	conn.SetConnMaxLifetime(time.Minute * 5)

	defer conn.Close()

	dbClient := repository.NewDbClient(conn, appConfig)
	register := register.NewRegister(dbClient)
	login := login.NewLoginUser(dbClient)

	var routers routers.Route
	routers.Register = register
	routers.Login = login
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", routers.Init())
}

func getMySqlDbConnection(appConfg config.AppConfig) (*sql.DB, error) {
	mySqlConnectionUrl := appConfg.GetAppConfigUrl()
	return sql.Open("mysql", mySqlConnectionUrl)
}
