package main

import (
	"fmt"
	"log"
	"net/http"
)

//Функция которая отвечает за создание пользователя и запись в БД, полученных со страницы данных
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_ = r.ParseForm()

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		age := r.FormValue("age")
		login := r.FormValue("login")
		password := r.FormValue("password")

		fmt.Printf("Name:%s\nSurname:%s\n", firstname, lastname)
		var _, err = database.Exec(
			"insert into user.dataofusers (firstname, lastname, age, login, password) values (?,?,?,?,?)",
			firstname, lastname, age, login, password)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/auth/", 301)
	} else {

		http.ServeFile(w, r, "html/PageRegister.html")
	}

}
