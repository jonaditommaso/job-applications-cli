# Job Applications CLI

A small CLI application written in Go to manage job applications stored in a PostgreSQL database.

## Features

The tool allows you to manage job applications with basic information such as:

- Company
- Role
- Application status
- Application channel

## Requirements

- Go 1.25 or later
- PostgreSQL
- A `DATABASE_URL` environment variable containing the database connection string

## Installation

1. Clone this repository.
2. Download the dependencies:

```bash
go mod download
```

3. Configure the `DATABASE_URL` environment variable.

Example:

```bash
export DATABASE_URL="postgres://user:password@localhost:5432/job_applications"
```

## Usage

Run the CLI directly with Go:

```bash
go run . <command>
```

Or build a binary and use it as `japp`:

```bash
go build -o japp .
./japp <command>
```

## Commands

### List applications

```bash
go run . list
```

Displays all registered job applications.

### Get an application

```bash
go run . get 20
```

Displays the application with the specified ID.

### Add a new application

```bash
go run . add
```

The command will interactively prompt for:

- Company
- Role
- Status
- Channel

### Update an application

```bash
go run . update 20
```

You will be asked to select a field and provide a new value.

Currently supported fields:

- `company`
- `role`
- `status`
- `channel`

### Delete an application

```bash
go run . delete 20
```

Deletes the application with the specified ID.

## Supported Statuses

The following application statuses are currently supported:

- `applied`
- `interview`
- `technical_test`
- `offer`
- `rejected`
- `ghosted`

## Technical Notes

- PostgreSQL access is implemented using `pgx`.
- The application expects an `applications` table containing at least the columns used by the CLI.
- The add and update workflows are interactive and run directly in the terminal.
