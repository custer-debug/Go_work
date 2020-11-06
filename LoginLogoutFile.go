package main

import (
	"database/sql"
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)


type User struct {
	ID 			int
	Firstname	string
	Lastname 	string
	Age 		int
	Login 		string
	Password 	string
}

var	(
	_user* User
	cookie = new(fiber.Cookie)
	)



func Logout(ctx *fiber.Ctx)error{
	ctx.ClearCookie()
	if cookie.Value == ""{
		fmt.Println("Cookie is empty")
	}
	_user = nil

	ctx.Redirect("/login")
	return nil
}



func CheckCookie(c *fiber.Ctx)error{
	if cookie.Value == "" || _user == nil{
		c.Redirect("/login")
		//goland:noinspection ALL
		return errors.New("Cookies is empty")
	}

	return nil
}




func GetLogin(c *fiber.Ctx) error {
	fmt.Println("GetLogin")

	return c.SendFile("./html/PageAut.html")
}



func PostLogin(c *fiber.Ctx) error {

	login := c.FormValue("login")
	password := c.FormValue("pass")

	db, _ := sql.Open("mysql", "root:Systemofadown2011@tcp(:8080)/user")
	var user = new(User)
	defer db.Close()
	rows := db.QueryRow("select * from dataofusers where login = ? and password = ?;", login, password)

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

	}else{
		fmt.Println("Welcome, " + user.Firstname + user.Lastname)
		_user = user
		json , _ := json2.Marshal(user)
		cookie.Name = "user"
		cookie.Value = string(json)
		c.Cookie(cookie)
		fmt.Println("PostLogin",cookie)
		c.Redirect("/welcome")
	}

	return nil

}
