package main

import (
	cruse "custer-debug/createuser"
	iof "custer-debug/in-out-function"
	"custer-debug/serverConst"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Static("/StaticFiles", "./StaticFiles")

	app.Get(serverConst.UrlLogin, iof.GetLogin)
	app.Post(serverConst.UrlLogin, iof.PostLogin)

	app.Get(serverConst.UrlProfile, GetWelcomeHandler)

	app.Get("/logout", iof.Logout)

	app.Get(serverConst.UrlCreate, cruse.GetCreateHandler)
	app.Post(serverConst.UrlCreate, cruse.PostCreateHandler)

	app.Get(serverConst.UrlSettings, HandlerGetSettings)
	app.Post(serverConst.UrlSettings, HandlerPostSettings)

	app.Get("/delete", DeleteUser)

	log.Fatal(app.Listen(":8800"))
}
