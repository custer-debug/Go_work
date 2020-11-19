package main

import (
	iof "custer-debug/in-out-function"
	"database/sql"
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
)

func getCookies(ctx *fiber.Ctx) *iof.User {
	var user = new(iof.User)
	err := json2.Unmarshal([]byte(ctx.Cookies("user")), user)
	if err != nil {
		log.Println("Cookie is empty")

	}
	return user
}

func GetWelcomeHandler(ctx *fiber.Ctx) error {

	if iof.GetUser().ID == 0 {
		fmt.Println("Id is zero")
		var u = getCookies(ctx)
		var db = connectToDatabase()
		iof.GetUserDataID(u.ID, db)
		db.Close()
	}

	return ctx.Render("./html/Welcome.html", iof.GetUser())
}

func HandlerGetSettings(ctx *fiber.Ctx) error {

	return ctx.Render("./html/Settings.html", iof.GetUser())
}

func connectToDatabase() *sql.DB {
	db, _ := sql.Open("mysql",
		"root:Systemofadown2011@tcp(:8080)/user")
	return db
}

//Function for change profile information
func changeUserData(b []byte) error {

	type Tmp struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Phone     string `json:"phone"`
		Birthday  string `json:"birthday"`
	}

	var tmp = new(Tmp)
	var user = iof.GetUser()
	json2.Unmarshal(b, tmp)
	user.Firstname = tmp.Firstname
	user.Lastname = tmp.Lastname
	user.Phone = tmp.Phone
	user.Birthday = tmp.Birthday

	var db = connectToDatabase()

	_, err := db.Exec("update dataofusers set firstname = ?,lastname = ?, Phone = ?, Birthday = ? "+
		"where ID = ?",
		user.Firstname,
		user.Lastname,
		user.Phone,
		user.Birthday,
		user.ID,
	)
	db.Close()
	if err != nil {

		return err
	}

	return nil
}

func changePassword(b []byte) error {

	type Tmp struct {
		OldPassword  string `json:"oldPassword"`
		NewPassword  string `json:"newPassword1"`
		NewPassword1 string `json:"newPassword2"`
	}
	fmt.Println(string(b))

	var tmp = new(Tmp)
	var user = iof.GetUser()
	json2.Unmarshal(b, tmp)
	fmt.Println(tmp.OldPassword)
	fmt.Println(user.Password)

	if tmp.OldPassword != user.Password {
		return errors.New("Incorrect password")
	}
	if tmp.NewPassword != tmp.NewPassword1 {
		return errors.New("Passwords do not match")
	}

	user.Password = tmp.NewPassword

	var db = connectToDatabase()

	_, err := db.Exec("update dataofusers set password = ? where ID = ?",
		user.Password,
		user.ID,
	)
	db.Close()
	if err != nil {
		return err
	}

	return nil
}

func HandlerPostSettings(ctx *fiber.Ctx) error {

	var find = "oldPassword"
	var data = string(ctx.Body())
	println(strings.Contains(data, find))
	if !strings.Contains(data, find) {
		log.Println(changeUserData(ctx.Body()))
		iof.SetCookie(ctx)

	} else {
		log.Println(changePassword(ctx.Body()))
	}
	return ctx.Redirect("/welcome")
}

func DeleteUser(ctx *fiber.Ctx) error {
	var db = connectToDatabase()

	var user = getCookies(ctx)
	_, err := db.Exec("delete from dataofusers where id = ?", user.ID)
	db.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	return ctx.Redirect("/logout")
}
