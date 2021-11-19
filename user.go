package main

import (
	"fmt"
	"log"
	"net/http"
)

func insert(w http.ResponseWriter, r *http.Request, data map[string]interface{}) bool {
	st := fmt.Sprintf("insert into users(name, email, password) values('%s' , '%s' , '%s')", data["name"], data["email"], data["password"])
	fmt.Println(st)
	//insert into database
	insert, err := conn.Query(st)
	if err != nil {
		log.Println(err.Error())
	}
	insert.Close()
	return true
}

// func delete(w http.ResponseWriter, r *http.Request) {

// }

// func update(w http.ResponseWriter, r *http.Request) {

// }
// func read(w http.ResponseWriter, r *http.Request) {

// }
