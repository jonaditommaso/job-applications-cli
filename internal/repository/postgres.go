package repository

import (
	"context"
	"fmt"

	"job-applications-cli/internal/application"

	"github.com/jackc/pgx/v5"
)

func Connect(connectionString string) (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), connectionString)
}

func GetApplications(conn *pgx.Conn) ([]application.JobApplication, error) {
	rows, err := conn.Query(
		context.Background(),
		"SELECT id, company, role, status, channel FROM applications",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applications []application.JobApplication

	for rows.Next() {
		var app application.JobApplication

		err := rows.Scan(
			&app.Id,
			&app.Company,
			&app.Role,
			&app.Status,
			&app.Channel,
		)
		if err != nil {
			return nil, err
		}

		applications = append(applications, app)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return applications, nil
}

func GetApplication(
	conn *pgx.Conn,
	appID int,
) (*application.JobApplication, error) {
	var app application.JobApplication

	err := conn.QueryRow(
		context.Background(),
		`
		SELECT id, company, role, status, channel
		FROM applications
		WHERE id = $1
		`,
		appID,
	).Scan(
		&app.Id,
		&app.Company,
		&app.Role,
		&app.Status,
		&app.Channel,
	)

	if err != nil {
		return nil, err
	}

	return &app, nil
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

func UpdateApplication(
	conn *pgx.Conn,
	appID int,
	field string,
	value string,
) error {
	allowedFields := map[string]bool{
		"company": true,
		"role":    true,
		"status":  true,
		"channel": true,
	}

	if !allowedFields[field] {
		return fmt.Errorf("invalid field")
	}

	_, err := conn.Exec(
		context.Background(),
		`
		UPDATE applications
		SET `+field+` = $1
		WHERE id = $2
		`,
		value,
		appID,
	)
	return err
}

func DeleteApplication(
	conn *pgx.Conn,
	appID int,
) error {
	_, err := conn.Exec(
		context.Background(),
		"DELETE FROM applications WHERE id = $1",
		appID,
	)
	return err
}
