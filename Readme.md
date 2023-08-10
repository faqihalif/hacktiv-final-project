# FINAL PROJECT

## Development Environtment

- Protobuf + GRPC

## Features

- Create Patient
- Get List Patient
- Filter Patient



## Run Locally

Clone the project

```bash
  git clone https://github.com/faqihalif/hacktiv-final-project.git
```

Migrate Database

```sql
    execute sql at directory docs/dbMigration/db_Migration.sql
```

Update setting in .env file


Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go mod tidy
```

Go to the cmd/web directory

```bash
  cd cmd/grpc/
```

Start the server

```bash
  go run main.go
```

