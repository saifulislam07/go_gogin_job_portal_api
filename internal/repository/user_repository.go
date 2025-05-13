package repository

import (
	"database/sql"
	"job_portal/internal/models"
	"log"
)

func CreateUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", user.Username, user.Password, user.Email)
	if err != nil {
		log.Printf("SQL insert error: %v", err)
	}
	return err
}

func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	var user models.User
	var profilePicture sql.NullString
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &profilePicture)
	if err != nil {
		return nil, err
	}

	if profilePicture.Valid {
		user.ProfilePicture = &profilePicture.String
	} else {
		user.ProfilePicture = nil
	}

	return &user, err
}
