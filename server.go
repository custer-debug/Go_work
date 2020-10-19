package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type User struct{
	Name string
	Age int
}

var database *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request){
	rows, err := database.Query("select * from mydb.mydb")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer rows.Close()

	user := []User{}
	for rows.Next(){
		u:= User{}
		err1 := rows.Scan(&u.Name,&u.Age)
		if err1 != nil{
			fmt.Print(err1)
			continue
		}
		user = append(user,u)
	}

tmpl, err1 := template.ParseFiles("template.html")
if err1 != nil{
	fmt.Println(err1)
	return
}
tmpl.Execute(w,user)
}


func main() {
	db, err := sql.Open("mysql","root:Systemofadown2011@tcp(:8080)/mydb")
	if err != nil {
		fmt.Println(err)
		return
	}

	database = db
	defer db.Close()

	http.HandleFunc("/",IndexHandler)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181",nil)





}