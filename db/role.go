package db

import (
	"context"
	"enterprise_v2/dto"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v5"
)

func createRoleTable() error {
	_, err := DbConnection.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS role(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(255),
	)`)
	if err != nil {
		return err
	}
	return nil
}

func CreateRole(role dto.InputRole) (string,error) {
	var id string

	err := DbConnection.QueryRow(context.Background(), `
		INSERT INTO role (
			name
		) VALUES (
			$1
		) RETURNING id
	`, role.Name).Scan(&id)
	if err != nil {
		return "",err
	}
	log.Info("Role created, id:",id)
	return id,nil
}
func GetRoles() ([]dto.OutputRole,error) {
	rows, err := DbConnection.Query(context.Background(), `
	SELECT * FROM role
	`)
	if err != nil {
		return nil,err
	}
	companies, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.OutputRole])
	if err != nil {
		return nil,err
	}
	log.Info("Returning all companies")
	return companies,nil
}
func GetOneRole(id string) (dto.OutputRole,error) {
	var role dto.OutputRole

	err := DbConnection.QueryRow(context.Background(), `
	SELECT id, name
	FROM role WHERE id=$1
`, id).Scan(
		&role.Id,
		&role.Name,
	)
	if err != nil {
		return dto.OutputRole{},err
	}

	log.Infof("Returning role id: %d", id)
	return role,nil
}
func PatchRole(id string, role dto.InputRole) (string,error) {
	ch, err := DbConnection.Exec(context.Background(), `
	UPDATE role
	SET  name=$1
	WHERE id=$2
	`, role.Name, id)
	if err != nil {
		return "",err
	}
	if ch.RowsAffected() < 1 {
		return "",fiber.ErrNotFound
	}

	log.Info("Role updated, id: ", id)
	return id,nil
}
func DeleteRole(id string) (string,error) {
	ch, err := DbConnection.Exec(context.Background(), `
	DELETE from role
	WHERE id=$1
	`, id)
	if err != nil {
		return "",err
	}
	if ch.RowsAffected() < 1 {
		return "",fiber.ErrNotFound
	}
	log.Info("Role deleted, id: ", id)
	return id,nil
}
