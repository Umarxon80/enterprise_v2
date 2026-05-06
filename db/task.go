package db

import (
	"context"
	"enterprise_v2/dto"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v5"
)

func createTaskTable() error {
	_, err := DbConnection.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS task (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID REFERENCES "user"(id),
		company_id UUID REFERENCES company(id) ,
		business_plan_id UUID REFERENCES business_plan(id) ,
		node_id INTEGER REFERENCES node(id) DEFAULT 2,
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

func GetTaskByUserId(user_id string) (dto.OutputTask, error) {
	var task dto.OutputTask

	err := DbConnection.QueryRow(context.Background(), `
	SELECT 
		t.id,
		u.id,
		u.first_name,
		u.last_name,
		c.id,
		c.name,
		bp.id,
		n.id,
		n.title
	FROM task t 
	LEFT JOIN "user" u ON t.user_id = u.id 
	LEFT JOIN business_plan bp ON t.business_plan_id = bp.id 
	LEFT JOIN company c ON t.company_id = c.id 
	LEFT JOIN node n ON t.node_id = n.id
	WHERE t.user_id = $1
	`, user_id).Scan(
		&task.Id,
		&task.User.Id,
		&task.User.FirstName,
		&task.User.LastName,
		&task.Company.Id,
		&task.Company.Name,
		&task.BusinessPlan.Id,
		&task.Node.Id,
		&task.Node.Title,
	)

	if err != nil {
		return dto.OutputTask{}, err
	}

	log.Infof("Returning task id: %s", task.Id)
	return task, nil
}
func TaskUpdateNode(task_id string,node_id int) error {
	ch, err := DbConnection.Exec(context.Background(), `
	UPDATE task
	SET  node_id=$1
	WHERE id=$2
	`, node_id,task_id)
	if err != nil {
		return  err
	}
	if ch.RowsAffected() < 1 {
		return  fiber.ErrNotFound
	}

	log.Info("User updated, id: ", task_id)
	return nil
}


func GetTasks(node_title string) ([]dto.OutputTask,error) {
	rows, err := DbConnection.Query(context.Background(), `
	SELECT 
		t.id,
		u.id,
		u.first_name,
		u.last_name,
		c.id,
		c.name,
		bp.id,
		n.id,
		n.title
	FROM task t 
	LEFT JOIN "user" u ON t.user_id = u.id 
	LEFT JOIN business_plan bp ON t.business_plan_id = bp.id 
	LEFT JOIN company c ON t.company_id = c.id 
	LEFT JOIN node n ON t.node_id = n.id  
	where n.title=$1
	`,node_title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (dto.OutputTask, error) {
		var task dto.OutputTask
		err := row.Scan(
		&task.Id,
		&task.User.Id,
		&task.User.FirstName,
		&task.User.LastName,
		&task.Company.Id,
		&task.Company.Name,
		&task.BusinessPlan.Id,
		&task.Node.Id,
		&task.Node.Title,
		)
		return task, err
	})

	if err != nil {
		return nil, err
	}
	log.Info("Returning all users")
	return tasks, nil
}