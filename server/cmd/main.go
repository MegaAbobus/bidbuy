package main

import (
	"bidbuy/internal/app/order"
	"bidbuy/internal/app/product"
	"bidbuy/internal/app/user"
	"bidbuy/internal/app/worker"
	"bidbuy/internal/server"

	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5"
)

type serverOpts struct {
	userSvc    user.Service
	productSvc product.Service
	orderSvc   order.Service
	workerSvc  worker.Service
}

func main() {
	ctx := context.Background()

	conn, err := initDatabase(ctx)
	if err != nil {
		log.Fatalf("database fatal %s", err)
	}

	app := fiber.New()
	app.Use(cors.New())

	serverOpts := createServices(conn)
	server := server.New(
		app,
		serverOpts.userSvc,
		serverOpts.productSvc,
		serverOpts.orderSvc,
		serverOpts.workerSvc,
	)

	server.InitRoutes()
	if err := server.Listen(":8080"); err != nil {
		log.Fatalf("fatal %s", err)
	}
}

func initDatabase(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres:changeme@localhost:5432/bidbuy?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}

func createServices(conn *pgx.Conn) *serverOpts {
	userStorage := user.NewStorage(conn)
	userService := user.NewService(userStorage)

	productStorage := product.NewStorage(conn)
	productService := product.NewService(productStorage)

	orderStorage := order.NewStorage(conn)
	orderService := order.NewService(orderStorage)

	workerStorage := worker.NewStorage(conn)
	workerService := worker.NewService(workerStorage)

	return &serverOpts{
		userSvc:    userService,
		productSvc: productService,
		orderSvc:   orderService,
		workerSvc:  workerService,
	}
}
