package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)


type User struct{
	ID int
	FirstName string
	LastName string
	Age int
	Login string
	Password string
}


var database *sql.DB


func menu(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		r.ParseForm()
		login := r.FormValue("login")
		Password := r.FormValue("pass")
		u := User{}
		rows := database.QueryRow("select * from owndata where login = ? and password = ?;", login, Password)

		rows.Scan(&u.ID,&u.FirstName, &u.LastName,&u.Age,&u.Login,&u.Login)


			if len(u.FirstName) > 0 && len(u.LastName) > 0{
				fmt.Printf("First Name:\t%s\nLast Name:\t%s\n",
					u.FirstName, u.LastName)

				tmpl, _ := template.ParseFiles("html/template.html")
				
				tmpl.Execute(w,u)
			}else{
				http.Redirect(w,r,"/",301)
			}



	}else{
		http.ServeFile(w, r, "html/menu.html")
	}


}


func main() {
	db, err := sql.Open("mysql","root:Systemofadown2011@tcp(:8080)/user")
	if err != nil {
		fmt.Println(err)
		return
	}

	database = db
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/",menu)

	http.Handle("/",router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181",nil)


}
