package main

import (
	cruse "custer-debug/createuser"
	iof "custer-debug/in-out-function"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Static("/StaticFiles", "./StaticFiles")

	app.Get("/login", iof.GetLogin)
	app.Post("/login", iof.PostLogin)

	app.Get("/welcome", GetWelcomeHandler)

	app.Get("/logout", iof.Logout)

	app.Get("/create", cruse.GetCreateHandler)
	app.Post("/create", cruse.PostCreateHandler)

	app.Get("/settings", HandlerGetSettings)
	app.Post("/settings", HandlerPostSettings)

	app.Get("/delete", DeleteUser)

	log.Fatal(app.Listen(":8800"))
}
