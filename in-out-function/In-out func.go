package in_out_function

import (
	"database/sql"
	json2 "encoding/json"
	"errors"
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

var (
	_user  *User
	cookie = new(fiber.Cookie)
)

func Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie()
	_user = nil
	return ctx.Redirect("/login")

}

func CheckCookie(c *fiber.Ctx) (fiber.Cookie, error) {
	if cookie.Value == "" || _user == nil {
		c.Redirect("/login")
		//goland:noinspection ALL
		return fiber.Cookie{}, errors.New("Cookies is empty")
	}

	return *cookie, nil
}

func GetLogin(c *fiber.Ctx) error {

	return c.SendFile("./html/PageAut.html")
}

func PostLogin(c *fiber.Ctx) error {

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
		_user = user
		json, _ := json2.Marshal(user)
		cookie.Name = "user"
		cookie.Value = string(json)
		c.Cookie(cookie)
		fmt.Println("PostLogin", cookie)
		c.Redirect("/welcome")
	}

	return nil

}
