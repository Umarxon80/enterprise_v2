package db

import (
	"context"
	"enterprise_v2/dto"

	"github.com/gofiber/fiber/v3/log"
)

func createBusinessPlanTable() error {
	_, err := DbConnection.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS business_plan (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		field VARCHAR(255),
		address VARCHAR(255),
		investment VARCHAR(255),
		document VARCHAR(255)
	)`)
	if err != nil {
		return err
	}
	return nil
}

func CreateBusinessPlan(business dto.InputBusiness) (string, error) {
	var id string

	err := DbConnection.QueryRow(context.Background(), `
		INSERT INTO business_plan(
			field,address,investment,document
		) VALUES (
			$1,$2,$3,$4
		) RETURNING id
	`, business.BusinessField, business.Address, business.Investment,business.Document).Scan(&id)
	if err != nil {
		return "", err
	}
	log.Info("businessPlan created, id:", id)
	return id, nil
}