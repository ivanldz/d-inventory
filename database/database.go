package database

import (
	"context"
	"fmt"

	"github.com/ivanldz/d-inventory/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable password=%s host=%s",
		s.DB.User,
		s.DB.Name,
		s.DB.Password,
		s.DB.Host,
	)

	return sqlx.ConnectContext(ctx, "postgres", connectionString)
}
