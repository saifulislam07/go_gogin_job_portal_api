package services

import (
	"database/sql"
	"job_portal/internal/models"
	"job_portal/internal/repository"
	"job_portal/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, user *models.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashPassword)

	return repository.CreateUser(db, user)
}

func LoginUser(db *sql.DB, username, password string) (string, error) {
	user, err := repository.GetUserByUserName(db, username)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	return utils.GenerateToken(user.Username, user.ID, user.IsAdmin)
}

func UpdateProfilePicture(db *sql.DB, id int, profilePicture string) error {
	return repository.UpdateProfilePicture(db, id, profilePicture)
}
