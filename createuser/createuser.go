package createuser

import (
	"custer-debug/emailPkg"
	sv "custer-debug/serverConst"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetCreateHandler(ctx *fiber.Ctx) error {
	fmt.Println("GetCreate")
	return ctx.SendFile(sv.HtmlCreateUser)
}

func PostCreateHandler(ctx *fiber.Ctx) error {

	fmt.Println("PostCreate")

	if len(ctx.FormValue("code")) != 0 {
		PostCheckCode(ctx)
	} else {
		parseBodyUser(ctx)
	}

	return nil
}

func parseBodyUser(ctx *fiber.Ctx) {
	var u = new(sv.User)
	if err := ctx.BodyParser(u); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u)
	sv.GetUser = u
	emailPkg.SendEmailCreateUser(u.Login)

}

func insertUser() {

	if sv.GetUser == nil {
		return
	}

	db, _ := sql.Open(sv.DataBase, sv.DataBaseSource)
	defer db.Close()

	var _, err = db.Exec("insert into "+
		"dataofusers(firstname, lastname, Birthday,Gender,Phone,login, password) "+
		"values(?,?,?,?,?,?,?);",
		sv.GetUser.Firstname,
		sv.GetUser.Lastname,
		sv.GetUser.Birthday,
		sv.GetUser.Gender,
		sv.GetUser.Phone,
		sv.GetUser.Login,
		sv.GetUser.Password,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func PostCheckCode(ctx *fiber.Ctx) error {

	code := ctx.FormValue("code")
	fmt.Println(code, sv.Code)

	if sv.Code == code {
		insertUser()
		sv.GetUser = nil
		msg, _ := json.Marshal(struct {
			Status string
		}{
			Status: "Success",
		})

		return ctx.Send(msg)
	}

	sv.GetUser = nil
	msg, _ := json.Marshal(struct {
		Status  string
		Subject string
		Body    string
	}{
		Status:  "Error",
		Subject: "Invalid code",
		Body:    "You entered the wrong code",
	})
	fmt.Println(msg)
	return ctx.Send(msg)

}
