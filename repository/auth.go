package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/efishery/error_message"
	"github.com/betawulan/efishery/model"
)

type authRepo struct {
	db *sql.DB
}

func (a authRepo) GetUser(ctx context.Context, filter model.RegisterFilter) (model.Register, error) {
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

	row := a.db.QueryRowContext(ctx, query, args...)
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

func (a authRepo) Register(ctx context.Context, register model.Register) error {
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

	res, err := a.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	register.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (a authRepo) Login(ctx context.Context, phone string, password string) (model.Register, error) {
	query, args, err := sq.Select("phone",
		"name",
		"role",
		"created_at").
		From("user").
		Where(sq.Eq{"phone": phone}).
		Where(sq.Eq{"password": password}).
		ToSql()
	if err != nil {
		return model.Register{}, nil
	}

	row := a.db.QueryRowContext(ctx, query, args...)
	var user model.Register
	err = row.Scan(&user.Phone,
		&user.Name,
		&user.Role,
		&user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Register{}, error_message.NotFound{Message: "phone or password incorrect"}
		}

		return model.Register{}, nil
	}

	return user, nil
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return authRepo{db: db}
}
