package main

import (
	iof "custer-debug/in-out-function"
	"database/sql"
	json2 "encoding/json"
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

	return ctx.Render("./html/welcome.html", iof.GetUser())
}

func HandlerGetSettings(ctx *fiber.Ctx) error {

	return ctx.Render("./html/settings.html", iof.GetUser())
}

func connectToDatabase() *sql.DB {
	db, _ := sql.Open("mysql",
		"root:Systemofadown2011@tcp(:8080)/user")
	return db
}

//Function for change profile information
func changeUserData(b []byte) string {
	fmt.Println("changeUserData")

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
		return "Trouble with DB"
	}

	return "OK"
}

func msgForSite(status string, body string) []byte {
	type Msg struct {
		Status string
		Body   string
	}
	var msg = new(Msg)
	msg.Status = status
	msg.Body = body

	res, _ := json2.Marshal(msg)
	return res
}

func changePassword(b []byte) []byte {
	fmt.Println("changePassword")

	type Tmp struct {
		OldPassword  string `json:"oldPassword"`
		NewPassword  string `json:"newPassword1"`
		NewPassword1 string `json:"newPassword2"`
	}
	fmt.Println(string(b))

	var tmp = new(Tmp)
	var user = iof.GetUser()
	json2.Unmarshal(b, tmp)

	if tmp.OldPassword != user.Password {
		return msgForSite("Error", "Incorrect password")
	}

	user.Password = tmp.NewPassword

	var db = connectToDatabase()

	_, err := db.Exec("update dataofusers set password = ? where ID = ?",
		user.Password,
		user.ID,
	)
	db.Close()
	if err != nil {
		return msgForSite("Error", "Trouble with DB")
	}

	return msgForSite("OK", "Password changed successfully")
}

func HandlerPostSettings(ctx *fiber.Ctx) error {

	var data = string(ctx.Body())
	if strings.Contains(data, "firstname") {
		iof.SetCookie(ctx)
		ctx.SendString(changeUserData(ctx.Body()))

	} else if strings.Contains(data, "oldPassword") {
		ctx.Send(changePassword(ctx.Body()))

	} else {
		log.Println("Bad request")
	}
	return nil
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
