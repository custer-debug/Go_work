package main

import (
	crus "custer-debug/createuser"
	iof "custer-debug/in-out-function"
	"database/sql"
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

func DeleteUser(ctx *fiber.Ctx) error {
	db, _ := sql.Open("mysql", "root:Systemofadown2011@tcp(:8080)/user")
	defer db.Close()
	var user = getCookies(ctx)
	_, err := db.Exec("delete from dataofusers where id = ?", user.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return ctx.Redirect("/logout")
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

	app.Get("/delete", DeleteUser)

}

func main() {
	app := fiber.New()

	MainHandler(app)

	log.Fatal(app.Listen(":80"))
}
