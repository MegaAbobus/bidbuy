package product

import (
	"bidbuy/internal/entities"
	"bidbuy/internal/presenter"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Storage interface {
	CreateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	ReadProduct(ctx context.Context, ID uint) (*entities.Product, error)
	UpdateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	ListProduct(ctx context.Context) (*[]presenter.Product, error)
	DeleteProduct(ctx context.Context, ID uint) error
}

type storage struct {
	conn *pgx.Conn
}

func NewStorage(conn *pgx.Conn) Storage {
	return &storage{
		conn: conn,
	}
}

func (st *storage) CreateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	var (
		id          uint
		description string
		body        string
		isAvailable bool
		price       float32
		rating      float32
	)

	query := `
	insert into "product"
	(
		description, body, is_available, price, rating
	)
	values
	(
		$1, $2, $3, $4, $5
	)
	returning product_id, description, body, is_available, price, rating`

	err := st.conn.QueryRow(ctx, query, product.Description, product.Body, product.IsAvailable, price, rating).Scan(&id, &description, &body, &price, &isAvailable, &price, &rating)
	if err != nil {
		return nil, err
	}

	return &entities.Product{
		ID:          id,
		Description: description,
		Body:        body,
		IsAvailable: isAvailable,
		Price:       price,
		Rating:      rating,
	}, nil
}
func (st *storage) DeleteOrder(ctx context.Context, ID uint) error {
	query := `delete from "order" where order_id = $1`

	commandTag, err := st.conn.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}

	return nil
}
func (st *storage) ReadProduct(ctx context.Context, ID uint) (*entities.Product, error) {
	var (
		id          uint
		description string
		body        string
		isAvailable bool
		price       float32
		rating      float32
	)

	query := `
	select * from "product" where product_id = $1
	`

	err := st.conn.QueryRow(ctx, query, ID).Scan(&id, &description, &body, &price, &isAvailable, &price, &rating)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}

		return nil, err
	}

	return &entities.Product{
		ID:          id,
		Description: description,
		Body:        body,
		IsAvailable: isAvailable,
		Price:       price,
		Rating:      rating,
	}, nil
}

func (st *storage) UpdateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	var (
		id          uint
		description string
		body        string
		isAvailable bool
		price       float32
		rating      float32
	)

	args := make([]interface{}, 0)
	args = append(args, product.ID, product.Description, product.Body, product.IsAvailable, product.Price, product.Rating)

	for i, a := range args {
		switch v := a.(type) {
		case uint:
			if v == 0 {
				args[i] = nil
			}
		case string:
			if v == "" {
				args[i] = nil
			}
		case float32:
			if v == 0 {
				args[i] = nil
			}
		}
	}

	query := `
	update "product"
	set description = coalesce($2, description), body = coalesce($3, body), is_available = coalesce($4, is_available), price = coalesce($5, price), rating = coalesce($6, rating)
	where product_id = $1
	returning product_id, description, body, is_available, price, rating`

	err := st.conn.QueryRow(ctx, query, args...).Scan(&id, &description, &body, &price, &isAvailable, &price, &rating)
	if err != nil {
		return nil, err
	}

	return &entities.Product{
		ID:          id,
		Description: description,
		Body:        body,
		IsAvailable: isAvailable,
		Price:       price,
		Rating:      rating,
	}, nil
}

func (st *storage) ListProduct(ctx context.Context) (*[]presenter.Product, error) {
	products := make([]presenter.Product, 0)

	query := `select * from "product"`

	rows, err := st.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id          uint
			description string
			body        string
			isAvailable bool
			price       float32
			rating      float32
		)

		err := rows.Scan(&id, &description, &body, &isAvailable, &price, &rating)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return &products, nil
			}

			return nil, err
		}

		products = append(products, presenter.Product{
			ID:          id,
			Description: description,
			Body:        body,
			IsAvailable: isAvailable,
			Price:       price,
			Rating:      rating,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &products, nil
}

func (st *storage) DeleteProduct(ctx context.Context, ID uint) error {
	query := `delete from "product" where product_id = $1`

	commandTag, err := st.conn.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}

	return nil
}
