package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/efishery/model"
	"github.com/betawulan/efishery/packages/error_message"
)

type authRepo struct {
	db *sql.DB
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
