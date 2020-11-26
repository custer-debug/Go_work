# ServerGolang
####This repository includes a backend and a frontend server. 

    
I use golang as my programming language. For the frontend, I used bootstrap. 
It is not in this repository, but you can download it from the link [bootstrap](https://bootstrap-4.ru/)
Also, to test this API, you need to install some packages. 
Before installing [fiber](https://github.com/gofiber/fiber/blob/master/.github/README_ru.md) , I advise you to read the documentation. 

    go get -u github.com/gofiber/fiber/v2

This is the link to work with the database, and the link to install 
[MySql for Golang](https://github.com/go-sql-driver/mysql):

    go get -u github.com/go-sql-driver/mysql

Communication is established using [XHR-requests](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest) 


####At the moment, this project has completed:
1. _Login page._
The problem of **different directory transitions** was solved with the help of `Cookies`. 
Here cookies are a json string like this:

        {
        ID: int
        FirstName: string
        LastName: string
        }

    ID makes it easy to navigate the database. First and last name for detailed user identification. 
    When you go to any other page of the site without authorization, the script will return to the login page.
    If there is a response from the backend that the password is incorrect, then a red notification appears, 
    and you need to try again.
    
2.  *Sign Up Page.* 
    There is a link under the submit button that leads to the registration form. 
    When filling in all the fields, a request sent with the addition of the user to the database.
    If all the data is correct, then the user moved to the login page.
    
3.  _Welcome Page._
    If authentication is successful, the user redirected to the welcome page.
     >In fact, I do not know yet what exactly the site will be directed to, 
     so I just made a greeting like:
     Hello, FirstName LastName.

    All pages use [templates](https://docs.gofiber.io/guide/templates).

4.  _The navigation bar_ has:<br/>
    4.1      `Packy` Site logo<br/>
    4.2     `Dropdown menu`, in which there are two cells, but this is still in development.<br/>
        There are two buttons in the right corner:<br/>
    4.3     `Settings` (Read on about this)<br/>
    4.4     `Log out` This button is responsible for logging out and logging off the session. 
        By clicking on this button, the cookies cleared, and the user returned to the login page.
 
5.  _Settings._ There are three tabs here:<br/>
    5.1 `Profile`Here the user can edit his data. Namely, the data that is in the database:
                 Name and surname, date of birth, phone number, login.
                 Upon successful editing, a positive response comes from the server<br/>
    5.2 `Private`Here the user can change the password. 
    The request will not be sent until the user enters the correct data, 
    namely the old password and the new password twice<br/>
    5.3 `Sites` This section is still under construction.



















