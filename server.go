package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type User struct{
	FirstName string
	SecondName string
	Age int
}

var database *sql.DB



func CreateHandler(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST"{
	err := r.ParseForm()
	if err != nil{
		log.Println("err")

		log.Println(err)
		return
	}
		firstname := r.FormValue("firstname")
		age := r.FormValue("age")
		secondname := r.FormValue("secondname")


		_,e := database.Exec("insert into mydb.mydb(first_name,age,second_name) values (?,?,?)",
			firstname,age,secondname)
			if e != nil{
				log.Println(e)
				return
			}

		http.Redirect(w,r,"/",301)

	}else{
		http.ServeFile(w,r,"index.html")
	}

}




func IndexHandler(w http.ResponseWriter, _ *http.Request){
	rows, err := database.Query("select * from mydb.mydb")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var user []User
	for rows.Next(){
		u:= User{}
		err1 := rows.Scan(&u.FirstName,&u.Age,&u.SecondName)
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
	http.HandleFunc("/create",CreateHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181",nil)





}