package handlers

import (
	"database/sql"
	"job_portal/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		user, err := services.GetUserByIdHandler(db, id)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserProfleHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var userUpdate struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}

		if err := c.ShouldBindJSON(&userUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")

		if userID != id && !isAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to update other user's profile"})
			return
		}

		updatedUser, err := services.UpdateUserProfile(db, id, userUpdate.Username, userUpdate.Email)
	}
}
