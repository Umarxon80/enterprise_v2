package db

import (
	"context"

	"github.com/gofiber/fiber/v3/log"
)

func createCompanyUserTable() error {
	_, err := DbConnection.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS company_user(
	user_id UUID REFERENCES "user"(id),
	company_id UUID REFERENCES company(id) 
	)`)
	if err != nil {
		return err
	}
	return nil
}

func CreateCompanyUser(user_id,company_id string) (error) {
	row := DbConnection.QueryRow(context.Background(), `
		INSERT INTO company_user (
			user_id,company_id
		) VALUES (
			$1,$2
		)
	`, user_id,company_id)
	if err := row.Scan(); err != nil && err.Error() != "no rows in result set" {
		return err
	}
	log.Info("company_user created successfully")
	return nil
}