package main

import (
	crus "custer-debug/createuser"
	iof "custer-debug/in-out-function"
	json2 "encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func getCookies(ctx *fiber.Ctx) *iof.User {
	var user = new(iof.User)
	err := json2.Unmarshal([]byte(ctx.Cookies("user")), user)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}

func GetWelcomeHandler(ctx *fiber.Ctx) error {

	var user = getCookies(ctx)

	return ctx.Render("./html/Welcome.html", user)
}

func HandlerGetSettings(ctx *fiber.Ctx) error {

	var user = getCookies(ctx)

	return ctx.Render("./html/Settings.html", user)
}

func MainHandler(app *fiber.App) {

	app.Static("/bootstrap", "./bootstrap")
	app.Get("/login", iof.GetLogin)
	app.Post("/login", iof.PostLogin)

	app.Get("/welcome", GetWelcomeHandler)

	app.Get("/logout", iof.Logout)

	app.Get("/create", crus.GetCreateHandler)
	app.Post("/create", crus.PostCreateHandler)

	app.Get("/settings", HandlerGetSettings)

}

func main() {
	app := fiber.New()

	MainHandler(app)

	log.Fatal(app.Listen(":80"))
}
