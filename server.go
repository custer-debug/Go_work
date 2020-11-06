package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func MainPageHandler(ctx *fiber.Ctx)error{
	//url = "/welcome"
	if  err := CheckCookie(ctx); err != nil{
		log.Println(err)
		return nil
	}

	return ctx.Render("./html/Welcome.html",_user)

}






func main() {
	app := fiber.New()
	app.Static("/StaticFiles", "./StaticFiles")

	app.Get("/login", GetLogin)
	app.Post("/login", PostLogin)

	app.Get("/welcome", MainPageHandler)
	app.Get("/logout",Logout)

	log.Fatal(app.Listen(":80"))
}
