package createuser

import (
	sv "custer-debug/serverConst"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetCreateHandler(ctx *fiber.Ctx) error {
	fmt.Println("GetCreate")
	return ctx.SendFile(sv.HtmlCreateUser)
}

func PostCreateHandler(ctx *fiber.Ctx) error {

	fmt.Println("PostCreate")
	fmt.Println(string(ctx.Body()))
	var u = new(sv.User)
	if err := ctx.BodyParser(u); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(u)

	db, _ := sql.Open(sv.DataBase, sv.DataBaseSource)
	defer db.Close()

	var _, err = db.Exec("insert into dataofusers(firstname, lastname, Birthday,Gender,Phone,login, password) "+
		"values(?,?,?,?,?,?,?);",
		u.Firstname,
		u.Lastname,
		u.Birthday,
		u.Gender,
		u.Phone,
		u.Login,
		u.Password,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Create User", u)
	return ctx.Redirect(sv.UrlLogin)
}
