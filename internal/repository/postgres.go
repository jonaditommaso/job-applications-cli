package repository

import (
	"context"

	"job-applications-cli/internal/application"

	"github.com/jackc/pgx/v5"
)

func Connect(connectionString string) (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), connectionString)
}

func CreateApplication(
	conn *pgx.Conn,
	app application.JobApplication,
) error {
	_, err := conn.Exec(
		context.Background(),
		`
		INSERT INTO applications (
			company,
			role,
			status,
			channel
		)
		VALUES ($1, $2, $3, $4)
		`,
		app.Company,
		app.Role,
		app.Status,
		app.Channel,
	)

	return err
}
