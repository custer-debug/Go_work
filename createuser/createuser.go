package createuser

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type User struct {
	Firstname string `json:"name"`
	Lastname  string `json:"surname"`
	Age       int    `json:"age"`
	Login     string `json:"login"`
	Password  string `json:"password" `
}

func GetCreateHandler(ctx *fiber.Ctx) error {
	fmt.Println("GetCreate")

	return ctx.SendFile("./html/CreateUser.html")
}

func PostCreateHandler(ctx *fiber.Ctx) error {
	fmt.Println("PostCreate")
	var u = new(User)
	if err := ctx.BodyParser(u); err != nil {
		log.Println(err)
		return err
	}
	db, _ := sql.Open("mysql", "root:Systemofadown2011@tcp(:8080)/user")
	defer db.Close()

	var _, err = db.Exec("insert into dataofusers(firstname, lastname, age, login, password) values(?,?,?,?,?);",
		u.Firstname,
		u.Lastname,
		u.Age,
		u.Login,
		u.Password)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("Create User", u)
	return ctx.Redirect("/login")
}
