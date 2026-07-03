package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"job-applications-cli/internal/application"
	"job-applications-cli/internal/repository"

	"github.com/joho/godotenv"
)

func main() {
	// Conect to the database
	godotenv.Load()

	conn, err := repository.Connect(
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	fmt.Println("Connected!")

	// Make the query
	rows, err := conn.Query(context.Background(), "SELECT id, company, role, status FROM applications") // Only these columns for now.
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Read the results.
	applications := []application.JobApplication{}

	for rows.Next() {
		var app application.JobApplication
		if err := rows.Scan(&app.Id, &app.Company, &app.Role, &app.Status); err != nil {
			panic(err)
		}
		applications = append(applications, app)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	application.InitializeApplications(applications)

	// Handle commands
	args := os.Args

	// 1. Validate input
	if len(args) < 2 {
		fmt.Println("usage: japp <command>")
		return
	}

	command := args[1]

	// 2. Routing
	switch command {
	case "get":
		if len(args) < 3 {
			fmt.Println("usage: japp get <id>")
			return
		}

		idInt, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
			return
		}

		app, err := application.GetApplication(idInt)
		if err != nil {
			fmt.Println("not found")
			return
		}
		fmt.Println(app)

	case "add":
		// TODO: Implement add command
	default:
		fmt.Println("unknown command")
	}
}
