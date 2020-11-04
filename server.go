package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

type User struct {
	Firstname	string `json:"firstname"`
	Lastname 	string `json:"lastname"`
	Age 		int 	`json:"age"`
	Login 		string `json:"login"`
	Password 	string `json:"password"`

}

var(
	user *User
	database   *sql.DB


)


func GetLogin(c *fiber.Ctx) error {
	user = nil
	fmt.Println("GetLogin")
	c.SendFile("./html/PageAut.html")

	return nil
}



func PostLogin(c *fiber.Ctx) error {

	err := c.BodyParser(&user)
	if err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "cannot parse json on struct",
		})
		return err
	}
	db, _ := sql.Open("mysql", "root:Systemofadown2011@tcp(:8080)/user")
	database = db
	defer db.Close()
	rows := database.QueryRow("select firstname,lastname,age from dataofusers where login = ? and password = ?;", user.Login, user.Password)
	rows.Scan(&user.Firstname,&user.Lastname,&user.Age)
	if user.Firstname != "" && user.Lastname != "" && user.Age > 0{
		fmt.Println("Welcome, " + user.Firstname)
		c.Redirect("/Welcome")
	}else{
		user = nil
		fmt.Println("Incorrect login or password")

	}

	return nil
}




func main() {
	app := fiber.New()
	app.Static("/CSS", "./CSS")


	app.Get("/login", GetLogin)
	app.Post("/login", PostLogin)


	log.Fatal(app.Listen(":80"))
}
