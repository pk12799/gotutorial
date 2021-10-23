package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	age   int    `json :"age"`
	email string `json:"email"`
}

func main() {
	s, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang")

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println()
	defer s.Close()
	insert, err := s.Query("insert into users(name, age, email) values('parvez', 22, 'parvze@gmail.ciom')")
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()
	d := s.
	read, err := s.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	for read.Next() {
		var tag Tag
		err = read.Scan(&tag.ID, &tag.Name, &tag.age, &tag.email)
		if err != nil {
			log.Panic(err.Error())
		}
		log.Println(tag.ID, tag.Name, tag.age, tag.email)
	}
	defer read.Close()
}
