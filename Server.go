package main

import (
	iof "custer-debug/in-out-function"
	json2 "encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func MainPageHandler(ctx *fiber.Ctx) error {
	c, err := iof.CheckCookie(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}
	var user iof.User
	err = json2.Unmarshal([]byte(c.Value), &user)
	return ctx.Render("./html/Welcome.html", user)

}

func main() {
	app := fiber.New()

	app.Static("", "./StaticFiles")
	app.Get("/login", iof.GetLogin)
	app.Post("/login", iof.PostLogin)

	app.Get("/welcome", MainPageHandler)
	app.Get("/logout", iof.Logout)

	log.Fatal(app.Listen(":80"))
}
