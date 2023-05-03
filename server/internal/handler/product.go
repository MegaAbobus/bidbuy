package handler

import (
	"bidbuy/internal/app/product"
	"bidbuy/internal/entities"
	"bidbuy/internal/presenter"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Product
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		result, err := service.InsertProduct(context.Background(), &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductSuccessResponse(result))
	}
}

func UpdateProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Product
		err := c.ParamsParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		err = c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		result, err := service.UpdateProduct(context.Background(), &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductSuccessResponse(result))
	}
}

func DeleteProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestParam entities.DeleteRequest
		err := c.ParamsParser(&requestParam)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		err = service.RemoveProduct(context.Background(), requestParam.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func ListProducts(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchProducts(context.Background())
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductsSuccessResponse(fetched))
	}
}

func GetProduct(service product.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestParam entities.GetRequest
		err := c.ParamsParser(&requestParam)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		result, err := service.GetProduct(context.Background(), requestParam.ID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}
		return c.JSON(presenter.ProductSuccessResponse(result))
	}
}
