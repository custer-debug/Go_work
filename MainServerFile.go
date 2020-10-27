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

//Главная страница, которая открывается после авторизации и при нажатии на значок Packy
func MainPage(w http.ResponseWriter, _ *http.Request) {
	tmpl, _ := template.ParseFiles("html/SuccessfullEnter.html")
	_ = tmpl.Execute(w, GlobalUser)
}

//Функция которая отвечает за редактирование данных из БД
func SettingHandleFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		if len(r.FormValue("delete")) != 0 {

			var _, err = database.Exec(
				"DELETE FROM dataofusers WHERE ID = ? and firstname = ? and lastname = ?",
				GlobalUser.ID, GlobalUser.FirstName, GlobalUser.LastName)

			if err != nil {
				log.Println(err)
			}
			fmt.Printf("Delete user:%s %s\n", GlobalUser.FirstName, GlobalUser.LastName)
			//LogoutFunc(w, r)
			GlobalUser = nil

		} else {
			fmt.Printf("Name:\t%s -> %s\n", GlobalUser.FirstName, r.FormValue("NewName"))
			fmt.Printf("Name:\t%s -> %s\n", GlobalUser.LastName, r.FormValue("NewSurname"))

			GlobalUser.FirstName = r.FormValue("NewName")
			GlobalUser.LastName = r.FormValue("NewSurname")

			var _, err = database.Exec(
				"update dataofusers set firstname = ?, lastname = ? where id = ?",
				GlobalUser.FirstName, GlobalUser.LastName, GlobalUser.ID)
			if err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/settings/", 301)
		}

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

	//var router = mux.NewRouter()

	http.HandleFunc("/auth/", LoginFunc)
	http.HandleFunc("/logout/", LogoutFunc)

	http.HandleFunc("/create/", CreateUserHandler)
	http.HandleFunc("/settings/", SettingHandleFunc)
	http.HandleFunc("/main_page/", MainPage)

	//	http.Handle("/", router)
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("./CSS/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

}
