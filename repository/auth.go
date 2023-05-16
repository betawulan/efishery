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

func (a authRepo) GetUser(ctx context.Context, filter model.UserFilter) (model.User, error) {
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
		return model.User{}, err
	}

	row := a.db.QueryRowContext(ctx, query, args...)
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	err = row.Scan(
		&user.Phone,
		&user.Name,
		&user.Role,
		&user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (a authRepo) Register(ctx context.Context, user model.User) error {
	user.CreatedAt = time.Now()

	query, args, err := sq.Insert("user").
		Columns("phone",
			"name",
			"role",
			"password",
			"created_at").
		Values(user.Phone,
			user.Name,
			user.Role,
			user.Password,
			user.CreatedAt).
		ToSql()
	if err != nil {
		return err
	}

	res, err := a.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (a authRepo) Login(ctx context.Context, phone string, password string) (model.User, error) {
	query, args, err := sq.Select("phone",
		"name",
		"role",
		"created_at").
		From("user").
		Where(sq.Eq{"phone": phone}).
		Where(sq.Eq{"password": password}).
		ToSql()
	if err != nil {
		return model.User{}, nil
	}

	row := a.db.QueryRowContext(ctx, query, args...)
	var user model.User
	err = row.Scan(&user.Phone,
		&user.Name,
		&user.Role,
		&user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, error_message.NotFound{Message: "phone or password incorrect"}
		}

		return model.User{}, nil
	}

	return user, nil
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return authRepo{db: db}
}
