package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var conn *sql.DB
var err error

func setupRoutes() {
	conn, err = sql.Open("mysql", "root@tcp(localhost:3306)/fourm_golang")
	if err != nil {
		log.Println("database not connected")
	}
	r := mux.NewRouter()
	r.HandleFunc("/", login)
	r.HandleFunc("/login", loginsubmit).Methods("POST")
	r.HandleFunc("/singup", singup)
	r.HandleFunc("/singupsub", singupsub)

	r.HandleFunc("/home", homepage)
	r.HandleFunc("/email", email)
	log.Fatal(http.ListenAndServe(":8000", r))

}

var key = []byte(os.Getenv("SESSION_KEY"))
var Store = sessions.NewCookieStore(key)

func main() {

	setupRoutes()
}
