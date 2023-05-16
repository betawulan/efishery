package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/efishery/model"
)

type registerRepo struct {
	db *sql.DB
}

func (r registerRepo) GetUser(ctx context.Context, filter model.RegisterFilter) (model.Register, error) {
	querySelect := sq.Select("phone",
		"name",
		"role",
		"created_at").
		From("user")

	if filter.Phone != "" {
		querySelect = querySelect.Where(sq.Eq{"phone": filter.Phone})
	}

	if filter.Name != "" {
		querySelect = querySelect.Where(sq.Eq{"name": filter.Name})
	}

	query, args, err := querySelect.ToSql()
	if err != nil {
		return model.Register{}, err
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if err != nil {
		return model.Register{}, err
	}

	var user model.Register
	err = row.Scan(
		&user.Phone,
		&user.Name,
		&user.Role,
		&user.CreatedAt)
	if err != nil {
		return model.Register{}, err
	}

	return user, nil
}

func (r registerRepo) Register(ctx context.Context, register model.Register) error {
	register.CreatedAt = time.Now()

	query, args, err := sq.Insert("user").
		Columns("phone",
			"name",
			"role",
			"password",
			"created_at").
		Values(register.Phone,
			register.Name,
			register.Role,
			register.Password,
			register.CreatedAt).
		ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	register.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func NewRegisterRepository(db *sql.DB) RegisterRepository {
	return registerRepo{
		db: db,
	}
}
