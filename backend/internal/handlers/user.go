package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cleanupDev/verbose-pancake/backend/internal/models"
	"github.com/cleanupDev/verbose-pancake/backend/internal/repositories"
	"github.com/gin-gonic/gin"
)

func ShowUsers(c *gin.Context) {
	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database is not connected!",
			"error":   err.Error(),
		})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database is not connected!",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Database is not connected!",
				"error":   err.Error(),
			})
			return
		}
		users = append(users, user)
	}

	usersJSON, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database is not connected!",
			"error":   err.Error(),
		})
		return
	}

	c.Data(http.StatusOK, "application/json", usersJSON)
}

func CreateUser(c *gin.Context) {
	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database is not connected!",
			"error":   err.Error(),
		})
		return
	}
	defer db.Close()

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid JSON!",
		})
		return
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", user.Username).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database is not connected!",
			"error":   err.Error(),
		})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Username already exists",
			"message": "Username already exists!",
		})
		return
	}

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", user.Email).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database is not connected!",
			"error":   err.Error(),
		})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Email address already exists",
			"message": "Email address already exists!",
		})
		return
	}

	hashedPassword, err := repositories.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password, email, first_name, last_name) VALUES (?, ?, ?, ?, ?)", user.Username, hashedPassword, user.Email, user.FirstName, user.LastName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created!",
	})
}

func LoginUser(c *gin.Context) {
	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database is not connected!",
			"error":   err.Error(),
		})
		return
	}
	defer db.Close()

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var hashedPassword string

	err = db.QueryRow("SELECT id, password, email FROM users WHERE username = ?", user.Username).Scan(&user.ID, &hashedPassword, &user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "User not found!",
		})
		return
	}

	err = repositories.ComparePasswords(hashedPassword, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Wrong password!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "models.User is logged in!",
		"data":    user,
	})
}

func GetUser(c *gin.Context) {
	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database is not connected!",
			"error":   err.Error(),
		})
		return
	}
	defer db.Close()

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = db.QueryRow("SELECT username, email, first_name, last_name FROM users WHERE id = ?", user.ID).Scan(&user.Username, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "User not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "This is your user!",
		"data":    user,
	})

}
