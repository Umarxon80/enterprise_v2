package db

import (
	"context"
	
	"github.com/gofiber/fiber/v3/log"
)

func createTaskLogTable() error {
	_, err := DbConnection.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS task_log (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		task_id UUID REFERENCES task(id),
		from_node_id INTEGER REFERENCES node(id) ,
		to_node_id INTEGER REFERENCES node(id) ,
		details VARCHAR(255) DEFAULT NULL
	)`)
	if err != nil {
		return err
	}
	return nil
}

func CreateTaskLog(task_id string,from_node_id int,to_node_id int,details string) ( error) {
	var id string
	err := DbConnection.QueryRow(context.Background(), `
		INSERT INTO task_log (
			task_id,from_node_id,to_node_id,details
		) VALUES (
			$1,$2,$3,$4
		) RETURNING id
	`,task_id,from_node_id,to_node_id,details).Scan(&id)
	if err != nil {
		return  err
	}
	log.Info("TaskLog created, id:", id)
	return nil
}