package db

import (
	"context"
	
	"github.com/gofiber/fiber/v3/log"
)

func createTaskTable() error {
	_, err := DbConnection.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS task (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID REFERENCES "user"(id),
		company_id UUID REFERENCES company(id) ,
		business_plan_id UUID REFERENCES business_plan(id) ,
		node_id INTEGER REFERENCES node(id) DEFAULT 1,
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now()
	)`)
	if err != nil {
		return err
	}
	return nil
}

func CreateTask(user_id,company_id,business_plan_id string) (string, error) {
	var id string

	err := DbConnection.QueryRow(context.Background(), `
		INSERT INTO task (
			user_id,company_id,business_plan_id
		) VALUES (
			$1,$2,$3
		) RETURNING id
	`, user_id, company_id, business_plan_id).Scan(&id)
	if err != nil {
		return  "",err
	}
	log.Info("Task created, id:", id)
	return id,nil
}