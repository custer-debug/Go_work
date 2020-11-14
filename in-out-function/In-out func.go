package in_out_function

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Age       int
	Login     string
	Password  string
}

func Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie()
	return ctx.Redirect("/login")

}

func GetLogin(c *fiber.Ctx) error {

	return c.SendFile("./html/Login.html")
}

func PostLogin(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	login := c.FormValue("login")
	password := c.FormValue("pass")

	db, _ := sql.Open("mysql", "root:Systemofadown2011@tcp(:8080)/user")
	var user = new(User)
	defer db.Close()
	rows := db.QueryRow("select * from dataofusers where login = ? and password = ?;",
		login, password)

	var err = rows.Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Age,
		&user.Login,
		&user.Password,
	)

	if err != nil {
		log.Println("Incorrect login or password")
		c.Redirect("/login")

	} else {
		fmt.Println("Welcome, " + user.Firstname + user.Lastname)
		json, _ := json2.Marshal(user)
		cookie.Name = "user"
		cookie.Value = string(json)
		c.Cookie(cookie)
		fmt.Println("PostLogin", cookie)

		c.Redirect("/welcome")
	}

	return nil

}
