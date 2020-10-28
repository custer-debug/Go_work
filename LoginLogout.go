package main

import (
	"fmt"
	"net/http"
)

func ClearCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ClearCookie")

	c.Name = ""
	c.Value = ""
	c.Path = ""
	c.MaxAge = -1
	http.SetCookie(w, c)

}

func LogoutFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.FormValue("logout") == "")

	//ClearCookie(w,r)
	//GlobalUser = nil

}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c = &http.Cookie{
		Name:   "Name",
		Value:  GlobalUser.FirstName,
		Path:   "/auth/",
		MaxAge: 1,
	}
	http.SetCookie(w, c)
	fmt.Println("SetCookie: " + c.String())

}

func CheckCookie(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("CheckCookie")

	if c.Name == "" {
		return false
	}

	fmt.Println(c.String())
	return true

}

//Функция которая отвечает за вход
func LoginFuncGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginFuncGet")
	if c.Name != "" {
		//MainPage(w,r)
		http.Redirect(w, r, "/main_page/", 301)
		fmt.Println(c.String())
	} else {

		http.ServeFile(w, r, "html/PageAut.html")
	}

}

func LoginFuncPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginFuncPost")
	var u = new(User)
	login := r.FormValue("login")
	Password := r.FormValue("pass")
	rows := database.QueryRow("select * from dataofusers where login = ? and password = ?;", login, Password)
	url := "/auth/"
	_ = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.Login, &u.Password)
	GlobalUser = u
	setCookie(w, r)

	if len(u.FirstName) > 0 && len(u.LastName) > 0 {
		fmt.Printf("User: %s %s enter on site\n",
			u.FirstName, u.LastName)
		url = "/main_page/"
		u = nil
	}

	http.Redirect(w, r, url, 301)
}
