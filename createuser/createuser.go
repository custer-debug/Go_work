package createuser

import (
	iof "custer-debug/in-out-function"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func GetCreateHandler(ctx *fiber.Ctx) error {
	fmt.Println("GetCreate")

	return ctx.SendFile("./html/CreateUser.html")
}

func PostCreateHandler(ctx *fiber.Ctx) error {
	fmt.Println("PostCreate")
	var u = new(iof.User)
	if err := ctx.BodyParser(u); err != nil {
		log.Println(err)
		return err
	}
	db, _ := sql.Open("mysql", "root:Systemofadown2011@tcp(:8080)/user")
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
		log.Println(err)
		return err
	}

	fmt.Println("Create User", u)
	return ctx.Redirect("/login")
}
