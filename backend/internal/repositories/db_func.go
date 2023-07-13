package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var DB_HOST string
var DB_PORT = "3306"
var dataSourceName string

func init() {
	err := godotenv.Load("./../environments/backend.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_USER = os.Getenv("USER")
	DB_PASSWORD = os.Getenv("PASSWORD")
	DB_NAME = os.Getenv("DATABASE")
	DB_HOST = os.Getenv("HOST")

	dataSourceName = DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME
}

func getDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Error opening database:", err.Error())
		return nil, err
	}

	return db, nil
}

func PingDatabase(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database is not connected!",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Database is connected!",
		})
	}
}

func InitDB(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id int(11) NOT NULL AUTO_INCREMENT,
			username varchar(255) NOT NULL,
			password varchar(255) NOT NULL,
			email varchar(255) NOT NULL,
			first_name varchar(255) NOT NULL,
			last_name varchar(255) NOT NULL,
			created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			UNIQUE KEY username (username),
			UNIQUE KEY email (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Database is initialized!",
	})
}

type User struct {
	ID        *int    `json:"id"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	CreatedAt *string `json:"created_at"`
}

func ShowUsers(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	usersJSON, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	c.Data(http.StatusOK, "application/json", usersJSON)
}

func CreateUser(c *gin.Context) {
	db, err := getDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password, email, first_name, last_name) VALUES (?, ?, ?, ?, ?)", user.Username, user.Password, user.Email, user.FirstName, user.LastName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User is created!",
	})
}
