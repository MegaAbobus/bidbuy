package presenter

import (
	"bidbuy/internal/entities"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID           uint    `json:"id"`
	PhoneNumber  string  `json:"phoneNumber"`
	Email        string  `json:"email"`
	AmountRansom float32 `json:"amountOfRansom"`
}

func UserSuccessResponse(data *entities.User) *fiber.Map {
	user := User{
		ID:           data.ID,
		PhoneNumber:  data.PhoneNumber,
		Email:        data.Email,
		AmountRansom: data.AmountRansom,
	}

	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

func UsersSuccessResponse(data *[]User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
