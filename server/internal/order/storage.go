package order

import (
	"bidbuy/internal/entities"
	"bidbuy/internal/presenter"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Storage interface {
	CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error)
	ReadOrder(ctx context.Context, ID uint) (*entities.Order, error)
	UpdateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error)
	ListOrder(ctx context.Context) (*[]presenter.Order, error)
	DeleteOrder(ctx context.Context, ID uint) error
}

type storage struct {
	conn *pgx.Conn
}

func NewStorage(conn *pgx.Conn) Storage {
	return &storage{
		conn: conn,
	}
}

func (st *storage) CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	query := `
	insert into order
	(
		user_id, address, price, status
	)
	values
	(
		$1, $2, $3, $4
	)`

	_, err := st.conn.Exec(ctx, query, order.UserID, order.Address, order.Price, order.Status)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (st *storage) ReadOrder(ctx context.Context, ID uint) (*entities.Order, error) {
	var (
		id      uint
		userId  uint
		address string
		price   float32
		status  string
	)

	query := `
	select * from "order" where id = $1
	`

	err := st.conn.QueryRow(ctx, query, ID).Scan(&id, &userId, &address, &price, &status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}

		return nil, err
	}

	return &entities.Order{
		ID:      id,
		UserID:  userId,
		Address: address,
		Price:   price,
		Status:  status,
	}, nil
}

func (st *storage) UpdateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	var (
		id      uint
		userId  uint
		address string
		price   float32
		status  string
	)

	query := `
	update order
	set user_id = $2, address = $3, price = $4, status = $5
	where id = $1
	returning id, user_id, address, price, status`

	err := st.conn.QueryRow(ctx, query, order.ID, order.UserID, order.Address, order.Price, order.Status).Scan(&id, &userId, &address, &price, &status)
	if err != nil {
		return nil, err
	}

	return &entities.Order{
		ID:      id,
		UserID:  userId,
		Address: address,
		Price:   price,
		Status:  status,
	}, nil
}

func (st *storage) ListOrder(ctx context.Context) (*[]presenter.Order, error) {
	orders := make([]presenter.Order, 0)

	query := `select * from "order"`

	rows, err := st.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id      uint
			userId  uint
			address string
			price   float32
			status  string
		)

		err := rows.Scan(&id, &userId, &address, &price, &status)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return &orders, nil
			}

			return nil, err
		}

		orders = append(orders, presenter.Order{
			ID:      id,
			UserId:  userId,
			Address: address,
			Price:   price,
			Status:  status,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &orders, nil
}

func (st *storage) DeleteOrder(ctx context.Context, ID uint) error {
	query := `delete from order where id = $1`

	commandTag, err := st.conn.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}

	return nil
}
