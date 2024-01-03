package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/thetinshusasi/golang_learning/database"
	"github.com/thetinshusasi/golang_learning/models"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	primaryKey := database.Database.Db.Create(&user)
	fmt.Println(primaryKey)
	responseUser := CreateResponseUser((user))
	return c.Status(200).JSON(responseUser)

}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	respUsers := []User{}
	for _, user := range users {
		respUsers = append(respUsers, CreateResponseUser(user))
	}

	return c.Status(200).JSON(respUsers)

}
