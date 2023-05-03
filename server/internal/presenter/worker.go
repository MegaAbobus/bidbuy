package presenter

import (
	"bidbuy/internal/entities"

	"github.com/gofiber/fiber/v2"
)

type Worker struct {
	ID                  uint   `json:"id"`
	Name                string `json:"name"`
	PasportSeriesNumber string `json:"pasportSeriesNumber"`
	PasportIssued       string `json:"pasportIssued"`
	PhoneNumber         string `json:"phoneNumber"`
}

func WorkerSuccessResponse(data *entities.Worker) *fiber.Map {
	worker := Worker{
		ID:                  data.ID,
		Name:                data.Name,
		PasportSeriesNumber: data.PasportSeriesNumber,
		PasportIssued:       data.PasportIssued,
		PhoneNumber:         data.PhoneNumber,
	}

	return &fiber.Map{
		"status": true,
		"data":   worker,
		"error":  nil,
	}
}

func WorkersSuccessResponse(data *[]Worker) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func WorkerErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
