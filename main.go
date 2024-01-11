package main

import (
	"context"

	"github.com/ivanldz/d-inventory/database"
	"github.com/ivanldz/d-inventory/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		fx.Invoke(
			func(db *sqlx.DB) {

				_, err := db.Query("SELECT * FROM USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
