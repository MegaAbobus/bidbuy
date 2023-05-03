package main

import (
	"bidbuy/internal/app/order"
	"bidbuy/internal/app/product"
	"bidbuy/internal/app/user"
	"bidbuy/internal/handler"
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
}

func main() {
	ctx := context.Background()

	conn, err := connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	userStorage := user.NewStorage(conn)
	userService := user.NewService(userStorage)

	productStorage := product.NewStorage(conn)
	productService := product.NewService(productStorage)

	orderStorage := order.NewStorage(conn)
	orderService := order.NewService(orderStorage)

	app := fiber.New()
	app.Use(cors.New())

	initRoutes(app, serverOpts{
		userSvc:    userService,
		productSvc: productService,
		orderSvc:   orderService,
	})

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

func initRoutes(app *fiber.App, opts serverOpts) {
	api := app.Group("/api/v1")

	users := api.Group("/users")
	users.Post("/", handler.CreateUser(opts.userSvc))
	users.Get("/", handler.ListUsers(opts.userSvc))
	users.Delete("/:id", handler.DeleteUser(opts.userSvc))
	users.Patch("/:id", handler.UpdateUser(opts.userSvc))
	users.Get("/:id", handler.GetUser(opts.userSvc))

	products := api.Group("/products")
	products.Post("/", handler.CreateProduct(opts.productSvc))
	products.Get("/", handler.ListProducts(opts.productSvc))
	products.Delete("/:id", handler.DeleteProduct(opts.productSvc))
	products.Patch("/:id", handler.UpdateProduct(opts.productSvc))
	products.Get("/:id", handler.GetProduct(opts.productSvc))

	orders := api.Group("/orders")
	orders.Post("/", handler.CreateOrder(opts.orderSvc))
	orders.Get("/", handler.ListOrders(opts.orderSvc))
	orders.Delete("/:id", handler.DeleteOrder(opts.orderSvc))
	orders.Patch("/:id", handler.UpdateOrder(opts.orderSvc))
	orders.Get("/:id", handler.GetOrder(opts.orderSvc))
}
