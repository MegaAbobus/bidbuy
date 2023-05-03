package worker

import (
	"bidbuy/internal/entities"
	"bidbuy/internal/presenter"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Storage interface {
	CreateWorker(ctx context.Context, worker *entities.Worker) (*entities.Worker, error)
	ReadWorker(ctx context.Context, ID uint) (*entities.Worker, error)
	UpdateWorker(ctx context.Context, worker *entities.Worker) (*entities.Worker, error)
	ListWorker(ctx context.Context) (*[]presenter.Worker, error)
	DeleteWorker(ctx context.Context, ID uint) error
}

type storage struct {
	conn *pgx.Conn
}

func NewStorage(conn *pgx.Conn) Storage {
	return &storage{
		conn: conn,
	}
}

func (st *storage) CreateWorker(ctx context.Context, worker *entities.Worker) (*entities.Worker, error) {
	var (
		id                  uint
		name                string
		pasportSeriesNumber string
		pasportIssued       string
		phoneNumber         string
	)

	query := `
	insert into "worker"
	(
		name, pasport_series_number, pasport_issued, phone_number
	)
	values
	(
		$1, $2, $3, $4
	)
	returning worker_id, name, pasport_series_number, pasport_issued, phone_number`

	err := st.conn.QueryRow(ctx, query, worker.Name, worker.PasportSeriesNumber, worker.PasportIssued, worker.PhoneNumber).Scan(&id, &name, &pasportSeriesNumber, &pasportIssued, &phoneNumber)
	if err != nil {
		return nil, err
	}

	return &entities.Worker{
		ID:                  id,
		Name:                name,
		PasportSeriesNumber: pasportSeriesNumber,
		PasportIssued:       pasportIssued,
		PhoneNumber:         phoneNumber,
	}, nil
}

func (st *storage) ReadWorker(ctx context.Context, ID uint) (*entities.Worker, error) {
	var (
		id                  uint
		name                string
		pasportSeriesNumber string
		pasportIssued       string
		phoneNumber         string
	)

	query := `
	select * from "worker" where worker_id = $1
	`

	err := st.conn.QueryRow(ctx, query, ID).Scan(&id, &name, &pasportSeriesNumber, &pasportIssued, &phoneNumber)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}

		return nil, err
	}

	return &entities.Worker{
		ID:                  id,
		Name:                name,
		PasportSeriesNumber: pasportSeriesNumber,
		PasportIssued:       pasportIssued,
		PhoneNumber:         phoneNumber,
	}, nil
}

func (st *storage) UpdateWorker(ctx context.Context, worker *entities.Worker) (*entities.Worker, error) {
	var (
		id                  uint
		name                string
		pasportSeriesNumber string
		pasportIssued       string
		phoneNumber         string
	)

	args := make([]interface{}, 0)
	args = append(args, worker.ID, worker.Name, worker.PasportSeriesNumber, worker.PasportIssued, worker.PhoneNumber)

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
		}
	}

	query := `
	update "worker"
	set name = coalesce($2, name), pasport_series_number = coalesce($3, pasport_series_number), pasport_issued = coalesce($4, pasport_issued), phone_number = coalesce($5, phone_number)
	where worker_id = $1
	returning worker_id, name, pasport_series_number, pasport_issued, phone_number`

	err := st.conn.QueryRow(ctx, query, args...).Scan(&id, &name, &pasportSeriesNumber, &pasportIssued, &phoneNumber)
	if err != nil {
		return nil, err
	}

	return &entities.Worker{
		ID:                  id,
		Name:                name,
		PasportSeriesNumber: pasportSeriesNumber,
		PasportIssued:       pasportIssued,
		PhoneNumber:         phoneNumber,
	}, nil
}

func (st *storage) ListWorker(ctx context.Context) (*[]presenter.Worker, error) {
	workers := make([]presenter.Worker, 0)

	query := `select * from "worker"`

	rows, err := st.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id                  uint
			name                string
			pasportSeriesNumber string
			pasportIssued       string
			phoneNumber         string
		)

		err := rows.Scan(&id, &name, &pasportSeriesNumber, &pasportIssued, &phoneNumber)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return &workers, nil
			}

			return nil, err
		}

		workers = append(workers, presenter.Worker{
			ID:                  id,
			Name:                name,
			PasportSeriesNumber: pasportSeriesNumber,
			PasportIssued:       pasportIssued,
			PhoneNumber:         phoneNumber,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &workers, nil
}

func (st *storage) DeleteWorker(ctx context.Context, ID uint) error {
	query := `delete from "worker" where worker_id = $1`

	commandTag, err := st.conn.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}

	return nil
}
