package service

import (
	"database/sql"
	"fiet-backend/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Println("Error scanning row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data error"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Println("Error scanning row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data error"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
