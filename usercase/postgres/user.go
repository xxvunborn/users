package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/unclebob/user/entities"
)

// UserUsercase is the Usercase for a user
type UserUsercase struct {
	Store *sqlx.DB
}

// Get the data from one user
func (uu *UserUsercase) Get(q *entities.UserQuery) (*entities.User, error) {
	query := squirrel.Select("*").From("users").Where("deleted_at is null")

	if q.ID != "" {
		query = query.Where("id = ?", q.ID)
	}

	if q.FullName != "" {
		query = query.Where("fullName = ?", q.FullName)
	}

	if q.Email != "" {
		query = query.Where("email = ?", q.Email)
	}

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := uu.Store.QueryRowx(sql, args...)

	c := &entities.User{}
	if err := row.StructScan(c); err != nil {
		return nil, err
	}

	return c, nil
}
