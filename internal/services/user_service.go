package services

import (
	"github.com/henny129/nexcraft-gocore-starter-kit/database"
	"github.com/henny129/nexcraft-gocore-starter-kit/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *models.User) error {
	_, err := database.DB.NamedExec(`
		INSERT INTO users (name, email) 
		VALUES (:name, :email)`,
		user,
	)
	return err
}

func GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := database.DB.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id int, user *models.User) error {
	_, err := database.DB.Exec(
		"UPDATE users SET name=$1, email=$2 WHERE id=$3",
		user.Name, user.Email, id,
	)
	return err
}

func DeleteUser(id int) error {
	_, err := database.DB.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
