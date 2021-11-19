package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "hello")
	outputHTML(w, "templ/index.html", nil)
	// insert(w, r)
}

func loginsubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")
	password := r.FormValue("password")

	// if err != nil {
	// 	return
	// }

	st := fmt.Sprintf("select password from users where email= '%s' ", email)
	fmt.Println(st)

	pass, err := conn.Query(st)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer pass.Close()
	var passd Password
	for pass.Next() {

		err = pass.Scan(&passd.password)
		if err != nil {
			log.Println(err.Error())
		}
	}
	fmt.Println(passd.password)
	passwd := passd.password
	con := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(password))
	if con == nil {

		fmt.Println(email, password, con)

		session, err := Store.Get(r, "ne_session")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		session.Options.Path = "/home"
		session.Values["user"] = email
		session.Values["login"] = false
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	} else {
		fmt.Println("err")
	}
}
