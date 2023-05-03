package handler

import (
	"bidbuy/internal/app/worker"
	"bidbuy/internal/entities"
	"bidbuy/internal/presenter"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateWorker(service worker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Worker
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.WorkerErrorResponse(err))
		}

		result, err := service.InsertWorker(context.Background(), &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WorkerErrorResponse(err))
		}
		return c.JSON(presenter.WorkerSuccessResponse(result))
	}
}

func UpdateWorker(service worker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Worker
		err := c.ParamsParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.WorkerErrorResponse(err))
		}

		err = c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.WorkerErrorResponse(err))
		}

		result, err := service.UpdateWorker(context.Background(), &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WorkerErrorResponse(err))
		}
		return c.JSON(presenter.WorkerSuccessResponse(result))
	}
}

func DeleteWorker(service worker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestParam entities.DeleteRequest
		err := c.ParamsParser(&requestParam)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.WorkerErrorResponse(err))
		}

		err = service.RemoveWorker(context.Background(), requestParam.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WorkerErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func ListWorkers(service worker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchWorkers(context.Background())
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WorkerErrorResponse(err))
		}
		return c.JSON(presenter.WorkersSuccessResponse(fetched))
	}
}

func GetWorker(service worker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestParam entities.GetRequest
		err := c.ParamsParser(&requestParam)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.WorkerErrorResponse(err))
		}

		result, err := service.GetWorker(context.Background(), requestParam.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WorkerErrorResponse(err))
		}
		return c.JSON(presenter.WorkerSuccessResponse(result))
	}
}
