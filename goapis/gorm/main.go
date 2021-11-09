package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

var DB *gorm.DB
var err error

const DNS = "root@tcp(localhost:3306)/golang?parseTime=true"

func inmig() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("canonot")
	}
	DB.AutoMigrate(&User{})
}

func Getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user []User
	DB.Find(&user)
	json.NewEncoder(w).Encode(user)
}
func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user User
	params := mux.Vars(r)
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}
func Deleteuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Delete(&user)

}
func Createuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
func Updateuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)

}
func initilize() {
	r := mux.NewRouter()

	r.HandleFunc("/user", Getusers).Methods("GET")
	r.HandleFunc("/user/{id}", Updateuser).Methods("PUT")
	r.HandleFunc("/user/{id}", Deleteuser).Methods("DELETE")
	r.HandleFunc("/create", Createuser).Methods("POST")
	r.HandleFunc("/user/{id}", user).Methods("GET")

	log.Fatal(http.ListenAndServe(":9001", r))
}

func main() {

	inmig()
	initilize()
}
