package presenter

import (
	"bidbuy/internal/entities"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID      uint
	UserId  uint
	Address string
	Price   float32
	Status  string
}

func OrderSuccessResponse(data *entities.Order) *fiber.Map {
	order := Order{
		ID:      data.ID,
		UserId:  data.UserID,
		Address: data.Address,
		Price:   data.Price,
		Status:  data.Status,
	}

	return &fiber.Map{
		"status": true,
		"data":   order,
		"error":  nil,
	}
}

func OrdersSuccessResponse(data *[]Order) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func OrderErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
