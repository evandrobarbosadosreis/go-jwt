package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/evandrobarbosadosreis/go-jwt/controllers"
	"github.com/evandrobarbosadosreis/go-jwt/driver"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.Connect()
	controllers := controllers.Controllers{}
	router := mux.NewRouter()

	router.HandleFunc(
		"/signup",
		controllers.Signup(db)).Methods("POST")

	router.HandleFunc(
		"/login",
		controllers.Login(db)).Methods("POST")

	router.HandleFunc(
		"/protected",
		controllers.AuthorizationMiddleware(controllers.ProtectedEndpoint())).Methods("GET")

	log.Println("Listen on port 4000...")
	log.Panic(http.ListenAndServe(":4000", router))
}
