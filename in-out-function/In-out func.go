package in_out_function

import (
	sv "custer-debug/serverConst"
	"database/sql"
	json2 "encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

//Function end session
func Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie()

	return ctx.Redirect(sv.UrlLogin)

}

//Function set cookie
func SetCookie(c *fiber.Ctx) {

	json, _ := json2.Marshal(struct {
		ID        int
		FirstName string
		LastName  string
	}{
		ID:        sv.GetUser.ID,
		FirstName: sv.GetUser.Firstname,
		LastName:  sv.GetUser.Lastname,
	})

	cookie := new(fiber.Cookie)
	cookie.Name = "user"
	cookie.Value = string(json)
	c.Cookie(cookie)
	fmt.Println("Set Cookie")
}

//Getting user data using the first and last name
func getUserData(l string, p string, db *sql.DB) error {

	row := db.QueryRow("select * from dataofusers where login = ? and password = ?;",
		l, p)

	var err = row.Scan(
		&sv.GetUser.ID,
		&sv.GetUser.Firstname,
		&sv.GetUser.Lastname,
		&sv.GetUser.Login,
		&sv.GetUser.Password,
		&sv.GetUser.Gender,
		&sv.GetUser.Birthday,
		&sv.GetUser.Phone,
	)
	getWebSites(db)
	return err
}

//Getting user data using the ID
func GetUserDataID(ID int, db *sql.DB) error {
	row := db.QueryRow("select * from dataofusers"+
		" where ID = ?",
		ID,
	)

	var err = row.Scan(
		&sv.GetUser.ID,
		&sv.GetUser.Firstname,
		&sv.GetUser.Lastname,
		&sv.GetUser.Login,
		&sv.GetUser.Password,
		&sv.GetUser.Gender,
		&sv.GetUser.Birthday,
		&sv.GetUser.Phone,
	)
	getWebSites(db)
	return err
}

//This function is not yet used
func getWebSites(db *sql.DB) {
	rows, err := db.Query("SELECT Link from websites where websites.ID_user = ? ",
		sv.GetUser.ID)

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var tmp string
		if err = rows.Scan(&tmp); err != nil {
			log.Println(err)
			return
		}

		sv.GetUser.WebSite = append(sv.GetUser.WebSite, tmp)
	}
}

//Handler GET-request on url: "login"
func GetLogin(c *fiber.Ctx) error {

	return c.SendFile(sv.HtmlLogin)
}

//Handler POST-request on url: "login"
func PostLogin(c *fiber.Ctx) error {

	login := c.FormValue("login")
	password := c.FormValue("password")
	println(login + password)
	db, _ := sql.Open("mysql", "root:Systemofadown2011@tcp(:8080)/user")

	var err = getUserData(login, password, db)
	db.Close()

	//Check correct password
	if err != nil {

		c.SendString("Error") //If incorrect login or password

	} else {

		fmt.Println("Welcome, " + sv.GetUser.Firstname + sv.GetUser.Lastname)
		SetCookie(c)
		c.SendString("Success")
	}

	return nil

}
