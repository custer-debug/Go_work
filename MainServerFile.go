package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Login     string
	Password  string
}

var database *sql.DB
var GlobalUser = new(User)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) { //создание пользователя
	if r.Method == "POST" {
		_ = r.ParseForm()

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		age := r.FormValue("age")
		login := r.FormValue("login")
		password := r.FormValue("password")

		fmt.Printf("Name:%s\nSurname:%s\n", firstname, lastname)
		_, err := database.Exec("insert into user.dataofusers (firstname, lastname, age, login, password) values (?,?,?,?,?)", firstname, lastname, age, login, password)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {

		http.ServeFile(w, r, "html/PageRegister.html")
	}

}

func MainPageOfServer(w http.ResponseWriter, r *http.Request) {
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

			tmpl, _ := template.ParseFiles("html/SuccessfullEnter.html")

			_ = tmpl.Execute(w, u)
			u = nil
		} else {

			http.Redirect(w, r, "/auth/", 301)
		}

	} else {

		http.ServeFile(w, r, "html/PageAut.html")
	}

}

func SettingHandleFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		_ = r.ParseForm()

	} else {
		tmpl, _ := template.ParseFiles("html/Settings.html")
		_ = tmpl.Execute(w, GlobalUser)
	}

}

//Настройки

func main() {
	db, err := sql.Open("mysql", "root:Systemofadown2011@tcp(:8080)/user")
	if err != nil {
		fmt.Println(err)
		return
	}

	database = db
	defer db.Close()

	http.HandleFunc("/auth/", MainPageOfServer)
	http.HandleFunc("/create/", CreateUserHandler)
	http.HandleFunc("/settings/", SettingHandleFunc)

	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("./CSS/"))))

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

}
