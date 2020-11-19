package in_out_function

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Birthday  string
	Gender    string
	Phone     string
	Login     string
	Password  string
	WebSite   []string
}

var user = new(User)

func GetUser() *User {
	return user
}

func Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie()

	return ctx.Redirect("/login")

}

func GetLogin(c *fiber.Ctx) error {

	return c.SendFile("./html/Login.html")
}

func SetCookie(c *fiber.Ctx) {

	json, _ := json2.Marshal(struct {
		ID        int
		FirstName string
		LastName  string
	}{
		ID:        user.ID,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
	})

	cookie := new(fiber.Cookie)
	cookie.Name = "user"
	cookie.Value = string(json)
	c.Cookie(cookie)
	fmt.Println("Set Cookie")
}

//Getting general info about user
func getUserData(l string, p string, db *sql.DB) error {

	row := db.QueryRow("select * from dataofusers where login = ? and password = ?;",
		l, p)

	var err = row.Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Login,
		&user.Password,
		&user.Gender,
		&user.Birthday,
		&user.Phone,
	)
	getWebSites(db)
	return err
}

func GetUserDataID(ID int, db *sql.DB) error {
	row := db.QueryRow("select * from dataofusers"+
		" where ID = ?",
		ID,
	)

	var err = row.Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Login,
		&user.Password,
		&user.Gender,
		&user.Birthday,
		&user.Phone,
	)
	getWebSites(db)
	return err
}

func getWebSites(db *sql.DB) {
	rows, err := db.Query("SELECT Link from websites where websites.ID_user = ? ",
		user.ID)

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var tmp string
		if err = rows.Scan(&tmp); err != nil {
			log.Println(err)
			return
		}

		user.WebSite = append(user.WebSite, tmp)
	}
}

func PostLogin(c *fiber.Ctx) error {
	login := c.FormValue("login")
	password := c.FormValue("pass")
	db, _ := sql.Open("mysql", "root:Systemofadown2011@tcp(:8080)/user")

	var err = getUserData(login, password, db)
	db.Close()

	//Check correct password
	if err != nil {
		log.Println("Incorrect login or password", err)
		c.Redirect("/login")

	} else {
		fmt.Println("Welcome, " + user.Firstname + user.Lastname)
		SetCookie(c)
		c.Redirect("/welcome")
	}

	return nil

}
