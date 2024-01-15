package repository

import (
	"context"

	"github.com/ivanldz/d-inventory/internal/entity"
)

const (
	qryInsertUser = `
	INSTERT INTO USERS (email, name, password) 
	VALUES ($1, $2, $3) RETURNING id;
	`

	qryGetUserByEmail = `
	select 
		id,
		email,
		name,
		password 
	from USERS
	where email = $1;
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.GetContext(ctx, user, qryGetUserByEmail, email)

	return user, err
}
