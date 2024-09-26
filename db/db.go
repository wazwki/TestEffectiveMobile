package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB

func DBInit() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		slog.Error(err.Error())
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Успешное подключение к базе данных")

	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		slog.Error(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		slog.Error(err.Error())
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		slog.Error(err.Error())
	}
	slog.Info("Миграции успешно применены")

}
