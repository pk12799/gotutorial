package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type singupst struct {
	name  string
	email string
}
type Password struct {
	password string
}

// start session and set cockie
// var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

//homepage function
func homepage(w http.ResponseWriter, r *http.Request) {
	sess, _ := Store.Get(r, "ne_session")

	status := sess.Values["login"]
	fmt.Println(status)
	if status == nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else if status == true {

		r.ParseForm()
		fmt.Println(r.Method)
		if r.Method == "GET" {
			outputHTML(w, "templ/home.html", nil)
		}

		if r.Method == "POST" {

			ques := r.FormValue("question")
			log.Println()
			st, _ := conn.Prepare("insert into question(question,userpost) values(?, ?)")
			re, _ := st.Exec(ques, 1)
			fmt.Println(re)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
		fmt.Fprint(w, "homepage")
		// w.WriteHeader(http.StatusOK)
	}
}

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// fmt.Println(data)
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
