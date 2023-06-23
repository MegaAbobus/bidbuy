package server

import (
	"bidbuy/internal/app/order"
	"bidbuy/internal/app/product"
	"bidbuy/internal/app/user"
	"bidbuy/internal/app/worker"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App

	userSvc    user.Service
	productSvc product.Service
	orderSvc   order.Service
	workerSvc  worker.Service
}

func New(
	app *fiber.App,
	userSvc user.Service,
	productSvc product.Service,
	orderSvc order.Service,
	workerSvc worker.Service,
) *Server {
	return &Server{
		app:        app,
		userSvc:    userSvc,
		productSvc: productSvc,
		orderSvc:   orderSvc,
		workerSvc:  workerSvc,
	}
}

func (s *Server) InitRoutes() {
	api := s.app.Group("/api/v1")

	orders := api.Group("/orders")
	orders.Post("/", s.createOrder(s.orderSvc))
	orders.Get("/", s.listOrders(s.orderSvc))
	orders.Delete("/:id", s.deleteOrder(s.orderSvc))
	orders.Patch("/:id", s.updateOrder(s.orderSvc))
	orders.Get("/:id", s.getOrder(s.orderSvc))
}

func (s *Server) Listen(addr string) error {
	err := s.app.Listen(":8080")
	if err != nil {
		return err
	}

	return nil
}
