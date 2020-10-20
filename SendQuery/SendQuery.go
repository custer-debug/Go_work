package SendQuery

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

//Function for add column
func AddColumn(db *sql.DB, n1 string,n2 string,a int, w http.ResponseWriter, r *http.Request){

	_,e := db.Exec("insert into user.owndata(firstname,lastname,age) " +
		"values (?,?,?)", n1,a,n2)
	if e != nil{
		log.Println(e)
		return
	}

	http.Redirect(w,r,"/",301)
}

//Function to display all columns
func ShowAll(db *sql.DB) *sql.Rows{
	rows, err := db.Query("select firstname,lastname,age from user.owndata")
	if err != nil{
		fmt.Println(err)
		panic(err)
	}

	return rows
}



