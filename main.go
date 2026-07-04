package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"job-applications-cli/internal/application"
	"job-applications-cli/internal/repository"

	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// Conect to the database
	godotenv.Load()

	conn, err := repository.Connect(os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	fmt.Println("Connected!")

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
	case "list":
		applications, err := repository.GetApplications(conn)
		if err != nil {
			panic(err)
		}

		for _, app := range applications {
			fmt.Printf(
				"Company: %s, Role: %s, Status: %s\n, Channel: %s\n",
				app.Company,
				app.Role,
				app.Status,
				app.Channel,
			)
		}
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

		app, err := repository.GetApplication(conn, idInt)
		if err != nil {
			fmt.Println("not found")
			return
		}
		fmt.Println(app)

	case "add":
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Company: ")
		company, _ := reader.ReadString('\n')

		fmt.Print("Role: ")
		role, _ := reader.ReadString('\n')

		fmt.Print("Status: ")
		status, _ := reader.ReadString('\n')

		fmt.Print("Channel: ")
		channel, _ := reader.ReadString('\n')

		newApp := application.JobApplication{
			Company: strings.TrimSpace(company),
			Role:    strings.TrimSpace(role),
			Status:  application.ApplicationStatus(strings.TrimSpace(status)),
			Channel: strings.TrimSpace(channel),
		}

		err := repository.CreateApplication(conn, newApp)
		if err != nil {
			fmt.Println("failed to create application because of error:", err)
			return
		}

		fmt.Println("application created")
	case "update":
		if len(args) < 3 {
			fmt.Println("usage: japp update <id>")
			return
		}

		idInt, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("invalid id")
			return
		}

		fields := []string{
			"company",
			"role",
			"status",
			"channel",
		}

		prompt := promptui.Select{
			Label: "Field to update",
			Items: fields,
		}

		_, field, err := prompt.Run()
		if err != nil {
			fmt.Println("selection cancelled")
			return
		}

		caser := cases.Title(language.English)

		valuePrompt := promptui.Prompt{
			Label: caser.String(field),
		}

		newValue, err := valuePrompt.Run()
		if err != nil {
			fmt.Println("input cancelled")
			return
		}

		err = repository.UpdateApplication(conn, idInt, field, newValue)
		if err != nil {
			fmt.Println("failed to update application because of error:", err)
			return
		}

		fmt.Printf("application %d updated\n", idInt)

	case "delete":
		if len(args) < 3 {
			fmt.Println("usage: japp delete <id>")
			return
		}

		idInt, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
			return
		}

		err = repository.DeleteApplication(conn, idInt)
		if err != nil {
			fmt.Println("failed to delete application because of error:", err)
			return
		}

		fmt.Println("application deleted")
	default:
		fmt.Println("unknown command")
	}
}
