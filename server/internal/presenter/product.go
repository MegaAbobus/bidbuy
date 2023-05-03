package presenter

import (
	"bidbuy/internal/entities"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID          uint    `json:"id"`
	Description string  `json:"description"`
	Body        string  `json:"body"`
	IsAvailable bool    `json:"isAvailable"`
	Price       float32 `json:"price"`
	Rating      float32 `json:"rating"`
}

func ProductSuccessResponse(data *entities.Product) *fiber.Map {
	product := Product{
		ID:          data.ID,
		Description: data.Description,
		Body:        data.Body,
		IsAvailable: data.IsAvailable,
		Price:       data.Price,
		Rating:      data.Rating,
	}

	return &fiber.Map{
		"status": true,
		"data":   product,
		"error":  nil,
	}
}

func ProductsSuccessResponse(data *[]Product) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func ProductErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
