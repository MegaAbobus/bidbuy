package handler

import (
	"bidbuy/internal/entities"
	"bidbuy/internal/order"
	"bidbuy/internal/presenter"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(service order.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Order
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.OrderErrorResponse(err))
		}

		result, err := service.InsertOrder(context.Background(), &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.OrderErrorResponse(err))
		}
		return c.JSON(presenter.OrderSuccessResponse(result))
	}
}

func UpdateOrder(service order.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Order
		err := c.ParamsParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.OrderErrorResponse(err))
		}

		err = c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.OrderErrorResponse(err))
		}

		result, err := service.UpdateOrder(context.Background(), &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.OrderErrorResponse(err))
		}
		return c.JSON(presenter.OrderSuccessResponse(result))
	}
}

func DeleteOrder(service order.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestParam entities.DeleteRequest
		err := c.ParamsParser(&requestParam)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.OrderErrorResponse(err))
		}

		err = service.RemoveOrder(context.Background(), requestParam.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.OrderErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func ListOrders(service order.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchOrders(context.Background())
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.OrderErrorResponse(err))
		}
		return c.JSON(presenter.OrdersSuccessResponse(fetched))
	}
}

func GetOrder(service order.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestParam entities.GetRequest
		err := c.ParamsParser(&requestParam)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.OrderErrorResponse(err))
		}

		result, err := service.GetOrder(context.Background(), requestParam.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.OrderErrorResponse(err))
		}
		return c.JSON(presenter.OrderSuccessResponse(result))
	}
}
