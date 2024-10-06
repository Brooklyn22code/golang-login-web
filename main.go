package main

import (
	"log"
	"net/http"

	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	Id       int64  `json: id`
	Username string `json: username`
	Password string `json: password`
	class    string `json: class`
}

func addUser(username string, password string) (string, bool) {

	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
		return "", false
	}
	defer rows.Close()
	for rows.Next() {
		var usernames string
		if err := rows.Scan(&usernames); err != nil {
			log.Fatal(err)
			return "", false
		}
		if username == usernames {
			return fmt.Sprintf("%v already been used!", username), false
		}
	}

	user1, err2 := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)

	if err2 != nil {
		fmt.Println(err2)
		return "", false
	}

	id, err3 := user1.LastInsertId()

	if err3 != nil {
		fmt.Println(err3)
		return "", false
	}

	st := fmt.Sprintf("you are the number %d user!", id)
	return st, true
}

func getUser(username string, password string) (string, bool) {
	var user User
	err := db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "User not found!", false
		}
		log.Println(err)
		return "Database error!", false
	}

	if user.Password == password {
		id := user.Id

		st := fmt.Sprintf("you are the number %d user!", id)

		return st, true
	} else {
		return "Password not correct!", false
	}
}

func main() {

	check := []int{}

	cfg := mysql.Config{
		User:   "root",
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "usersshark",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/statics", "./statics")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
		check = append(check, 1)
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
		check = append(check, 2)
	})

	router.POST("/submit", func(c *gin.Context) {
		username := c.PostForm("Username")
		password := c.PostForm("Password")
		fmt.Println(username, password)
		lastElement := check[len(check)-1]
		if lastElement == 1 {
			st, err := addUser(username, password)
			if !err {
				c.String(http.StatusNotAcceptable, st)
			} else {
				c.HTML(http.StatusOK, "home.html", gin.H{
					"Username": username,
					"Password": password,
					"ID":       st,
					"st":       "signup successful",
				})

			}
		}

		if lastElement == 2 {
			st, err := getUser(username, password)
			if err {
				c.HTML(http.StatusOK, "home.html", gin.H{
					"Username": username,
					"Password": password,
					"ID":       st,
					"st":       "login successful",
				})
			} else {
				c.String(http.StatusNotAcceptable, st)
			}
		}

	})

	router.Run(":8080")
}
