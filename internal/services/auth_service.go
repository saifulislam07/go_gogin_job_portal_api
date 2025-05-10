package services

import (
	"database/sql"
	"job_portal/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, user *models.User) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashPassword)
}
