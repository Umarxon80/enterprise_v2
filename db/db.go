package db

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DbConnection *pgxpool.Pool

func Connect() {
	var err error
	connstring := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
	DbConnection, err = pgxpool.New(context.Background(), connstring)
	if err != nil {
		log.Fatalf("Error connecting db: %w", err)
	}
	if err := createCompaniesTable(); err != nil {
		log.Fatalf("Error creating companies table %v", err)
	}
	if err := createRoleTable(); err != nil {
		log.Fatalf("Error creating role table %v", err)
	}
	if err := createUserTable(); err != nil {
		log.Fatalf("Error creating user table %v", err)
	}
	if err := createOtpTable(); err != nil {
		log.Fatalf("Error creating otp table %v", err)
	}
	if err := createBillTable(); err != nil {
		log.Fatalf("Error creating otp table %v", err)
	}

}
