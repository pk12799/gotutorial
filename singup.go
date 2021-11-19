package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

func singup(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "sdghdfhgf")

	if r.Method != "POST" {

		errs := singupst{
			name:  "parvez",
			email: "pkp@gmail.com",
		}
		myvar := map[string]interface{}{"MyVar": "Foo Bar Baz", "Name": []int{1, 2, 3}, "struct": errs}

		outputHTML(w, "templ/singup.html", myvar)

	}

}

func singupsub(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Method)
	pass := r.FormValue("password")
	cpass := r.FormValue("cpass")
	//compare to values'

	re := reflect.DeepEqual(pass, cpass)
	if re {
		name := r.FormValue("name")
		passwo, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err.Error())
		}
		ema := r.FormValue("email")
		stat := r.FormValue("status")
		if stat == "1" {

			user := map[string]interface{}{"name": name, "email": ema, "password": passwo}
			insert(w, r, user)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Fprint(w, "plese accept the policy")
		}

	}
	fmt.Println("return singup")

}
