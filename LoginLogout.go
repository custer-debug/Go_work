package main

import (
	"fmt"
	"net/http"
)

func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/auth/", 301)
	GlobalUser = nil
}

//Функция которая отвечает за вход
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	var u = new(User)

	if r.Method == "POST" {

		_ = r.ParseForm()
		login := r.FormValue("login")
		Password := r.FormValue("pass")
		rows := database.QueryRow("select * from dataofusers where login = ? and password = ?;", login, Password)

		_ = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.Login, &u.Password)
		GlobalUser = u
		if len(u.FirstName) > 0 && len(u.LastName) > 0 {
			fmt.Printf("First Name:\t%s\nLast Name:\t%s\n",
				u.FirstName, u.LastName)

			http.Redirect(w, r, "/main_page/", 301)
			u = nil
		} else {

			http.Redirect(w, r, "/auth/", 301)
		}
	} else {

		http.ServeFile(w, r, "html/PageAut.html")
	}

}
