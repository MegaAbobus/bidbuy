package main

import (
	"bidbuy/internal/handler"
	"bidbuy/internal/order"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	orderStorage := order.NewStorage(conn)
	orderService := order.NewService(orderStorage)

	app := fiber.New()
	app.Use(cors.New())

	initRoutes(app, orderService)

	app.Listen(":8080")
}

func connect(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres:changeme@localhost:5432/bidbuy?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}

func initRoutes(app *fiber.App, service order.Service) {
	api := app.Group("/api/v1")

	orders := api.Group("/orders")
	orders.Post("/", handler.CreateOrder(service))
	orders.Get("/", handler.ListOrders(service))
	orders.Delete("/:id", handler.DeleteOrder(service))
	orders.Patch("/:id", handler.UpdateOrder(service))
	orders.Get("/:id", handler.GetOrder(service))
}
