package db

import (
	"context"
	"enterprise_v2/dto"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v5"
)

func createCompaniesTable() error {
	_, err := DbConnection.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS company(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(255),
	financial_history TEXT,
	industry VARCHAR(255)
	)`)
	if err != nil {
		return err
	}
	return nil
}

func CreateCompany(company dto.InputCompany) (string,error) {
	var id string

	err := DbConnection.QueryRow(context.Background(), `
		INSERT INTO company (
			name, financial_history, industry
		) VALUES (
			$1, $2, $3
		) RETURNING id
	`, company.Name, company.FinancialHistory, company.Industry).Scan(&id)
	if err != nil {
		return "",err
	}
	log.Info("Company created, id:",id)
	return id,nil
}
func GetCompanies() ([]dto.OutputCompany,error) {
	rows, err := DbConnection.Query(context.Background(), `
	SELECT * FROM company
	`)
	if err != nil {
		return nil,err
	}
	companies, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.OutputCompany])
	if err != nil {
		return nil,err
	}
	log.Info("Returning all companies")
	return companies,nil
}
func GetOneCompany(id string) (dto.OutputCompany,error) {
	var company dto.OutputCompany

	err := DbConnection.QueryRow(context.Background(), `
	SELECT id, name, financial_history, industry
	FROM company WHERE id=$1
`, id).Scan(
		&company.Id,
		&company.Name,
		&company.FinancialHistory,
		&company.Industry,
	)
	if err != nil {
		return dto.OutputCompany{},err
	}

	log.Infof("Returning company id: %d", id)
	return company,nil
}
func PatchCompany(id string, company dto.InputCompany) (string,error) {
	ch, err := DbConnection.Exec(context.Background(), `
	UPDATE company
	SET  name=$1, financial_history=$2, industry=$3
	WHERE id=$4
	`, company.Name, company.FinancialHistory, company.Industry, id)
	if err != nil {
		return "",err
	}
	if ch.RowsAffected() < 1 {
		return "",fiber.ErrNotFound
	}

	log.Info("Company updated, id: ", id)
	return id,nil
}
func DeleteCompany(id string) (string,error) {
	ch, err := DbConnection.Exec(context.Background(), `
	DELETE from company
	WHERE id=$1
	`, id)
	if err != nil {
		return "",err
	}
	if ch.RowsAffected() < 1 {
		return "",fiber.ErrNotFound
	}
	log.Info("Company deleted, id: ", id)
	return id,nil
}
