package db

import (
	"context"
	"database/sql"
	"enterprise_v2/dto"
	"errors"

	"github.com/gofiber/fiber/v3/log"
)

// INSERT INTO node (title, role_id, prev_node_id) VALUES ('submit',    'f873d964-c2fc-478d-865c-5329a05fa5c9', NULL);
// INSERT INTO node (title, role_id, prev_node_id) VALUES ('payment',   'f873d964-c2fc-478d-865c-5329a05fa5c9', 1);
// INSERT INTO node (title, role_id, prev_node_id) VALUES ('review',    'a18c656e-a0ce-4748-a923-2e1e0f502f62', 2);
// INSERT INTO node (title, role_id, prev_node_id) VALUES ('committee', 'd5942d80-fe87-460b-ae0d-3354307a74c8', 3);

func createNodeTable() error {
    _, err := DbConnection.Exec(context.Background(), `
    CREATE TABLE IF NOT EXISTS node (
        id              SERIAL PRIMARY KEY,
        title           VARCHAR(255) NOT NULL,
        role_id         UUID REFERENCES role(id),
        prev_node_id    INTEGER REFERENCES node(id)
    ) `)
    if err != nil {
        return err
    }
    return nil
}

func GetOneNode(id int) (dto.OutputNode, error) {
	var node dto.OutputNode

	err := DbConnection.QueryRow(context.Background(), `
	SELECT 
		u.id, 
		u.title, 
		u.prev_node_id,
		r.name AS role, 
	FROM node n 
	LEFT JOIN role r ON n.role_id = r.id 
	WHERE u.id = $1
	`, id).Scan(
		&node.Id,
		&node.Title,
		&node.PrevNodeID,
		&node.Role,
	)
	if err != nil {
		return dto.OutputNode{}, err
	}

	log.Infof("Returning node, id: %s", id)
	return node, nil
}

// 0 if node finished
//
// 0+ int if any more nodes
func GetNextNode(prev_node_id int) (int, error) {
	var node dto.OutputNode

	err := DbConnection.QueryRow(context.Background(), `
	SELECT 
		n.id, 
		n.title, 
		n.prev_node_id,
		COALESCE(r.name, '') 
	FROM node n 
	LEFT JOIN role r ON n.role_id = r.id 
	WHERE n.prev_node_id = $1
	`, prev_node_id).Scan(
		&node.Id,
		&node.Title,
		&node.PrevNodeID,
		&node.Role,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	log.Infof("Returning node id: %d", node.Id)
	return node.Id, nil
}