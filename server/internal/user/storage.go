package user

import (
	"bidbuy/internal/entities"
	"bidbuy/internal/presenter"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Storage interface {
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	ReadUser(ctx context.Context, ID uint) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	ListUser(ctx context.Context) (*[]presenter.User, error)
	DeleteUser(ctx context.Context, ID uint) error
}

type storage struct {
	conn *pgx.Conn
}

func NewStorage(conn *pgx.Conn) Storage {
	return &storage{
		conn: conn,
	}
}

func (st *storage) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	var (
		id             uint
		phoneNumber    string
		amountOfRansom float32
		email          string
	)

	query := `
	insert into "user"
	(
		phone_number, amount_of_ransom, email
	)
	values
	(
		$1, $2, $3
	)
	returning user_id, phone_number, amount_of_ransom, email`

	err := st.conn.QueryRow(ctx, query, user.PhoneNumber, user.AmountRansom, user.Email).Scan(&id, &phoneNumber, &amountOfRansom, &email)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:           id,
		PhoneNumber:  phoneNumber,
		AmountRansom: amountOfRansom,
		Email:        email,
	}, nil
}

func (st *storage) ReadUser(ctx context.Context, ID uint) (*entities.User, error) {
	var (
		id             uint
		phoneNumber    string
		amountOfRansom float32
		email          string
	)

	query := `
	select * from "user" where user_id = $1
	`

	err := st.conn.QueryRow(ctx, query, ID).Scan(&id, &phoneNumber, &amountOfRansom, &email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}

		return nil, err
	}

	return &entities.User{
		ID:           id,
		PhoneNumber:  phoneNumber,
		AmountRansom: amountOfRansom,
		Email:        email,
	}, nil
}

func (st *storage) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	var (
		id             uint
		phoneNumber    string
		amountOfRansom float32
		email          string
	)

	args := make([]interface{}, 0)
	args = append(args, user.ID, user.PhoneNumber, user.AmountRansom, user.Email)

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
	update "user"
	set phone_number = coalesce($2, phone_number), amount_of_ransom = coalesce($3, amount_of_ransom), email = coalesce($4, email)
	where user_id = $1
	returning user_id, phone_number, amount_of_ransom, email`

	err := st.conn.QueryRow(ctx, query, args...).Scan(&id, &phoneNumber, &amountOfRansom, &email)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:           id,
		PhoneNumber:  phoneNumber,
		Email:        email,
		AmountRansom: amountOfRansom,
	}, nil
}

func (st *storage) ListUser(ctx context.Context) (*[]presenter.User, error) {
	users := make([]presenter.User, 0)

	query := `select * from "user"`

	rows, err := st.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id             uint
			phoneNumber    string
			amountOfRansom float32
			email          string
		)

		err := rows.Scan(&id, &phoneNumber, &amountOfRansom, &email)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return &users, nil
			}

			return nil, err
		}

		users = append(users, presenter.User{
			ID:           id,
			PhoneNumber:  phoneNumber,
			Email:        email,
			AmountRansom: amountOfRansom,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

func (st *storage) DeleteUser(ctx context.Context, ID uint) error {
	query := `delete from "user" where user_id = $1`

	commandTag, err := st.conn.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}

	return nil
}
