package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

//Главная страница, которая открывается после авторизации и при нажатии на значок Packy
func MainPage(w http.ResponseWriter, _ *http.Request) {
	tmpl, _ := template.ParseFiles("html/SuccessfullEnter.html")
	_ = tmpl.Execute(w, GlobalUser)
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

func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/auth/", 301)
	GlobalUser = nil
}

//Функция которая отвечает за редактирование данных из БД
func SettingHandleFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		_ = r.ParseForm()
		fmt.Printf("Name:%s\nSurname:%s\n", r.FormValue("NewName"), r.FormValue("NewSurname"))
		GlobalUser.FirstName = r.FormValue("NewName")
		GlobalUser.LastName = r.FormValue("NewSurname")

		var _, err = database.Exec(
			"update dataofusers set firstname = ?, lastname = ? where id = ?",
			GlobalUser.FirstName, GlobalUser.LastName, GlobalUser.ID)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/main_page/", 301)
	} else {
		tmpl, _ := template.ParseFiles("html/Settings.html")
		_ = tmpl.Execute(w, GlobalUser)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:Systemofadown2011@tcp(:8080)/user")
	if err != nil {
		fmt.Println(err)
		return
	}

	database = db
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/auth/", LoginFunc)
	router.HandleFunc("/logout/", LogoutFunc)

	router.HandleFunc("/create/", CreateUserHandler)
	router.HandleFunc("/settings/", SettingHandleFunc)
	router.HandleFunc("/main_page/", MainPage)

	http.Handle("/", router)
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("./CSS/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

}
