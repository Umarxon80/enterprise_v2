package db

import (
	"context"
// 	"enterprise_v2/dto"
// 	"time"

// 	"github.com/gofiber/fiber/v3"
// 	"github.com/gofiber/fiber/v3/log"
// 	"github.com/jackc/pgx/v5"
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



// func CreateUser(user dto.InputUser) (string, error) {
// 	var id string

// 	err := DbConnection.QueryRow(context.Background(), `
// 		INSERT INTO "user" (
// 			first_name,last_name,email,password
// 		) VALUES (
// 			$1,$2,$3,$4
// 		) RETURNING id
// 	`, user.FirstName, user.LastName, user.Email, user.Password).Scan(&id)
// 	if err != nil {
// 		return "", err
// 	}
// 	log.Info("User created, id:", id)
// 	return id, nil
// }
// func GetUsers() ([]dto.OutputUser, error) {
// 	rows, err := DbConnection.Query(context.Background(), `
// 	SELECT 
// 	u.id, u.first_name, u.last_name, u.email, u.password, u.is_active, u.created_at, u.updated_at,
// 	r.id AS role_id, r.name AS role_name
// 	FROM "user" u 
// 	LEFT JOIN role r ON u.role_id = r.id  
// 	`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	users, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (dto.OutputUser, error) {
// 		var u dto.OutputUser
// 		err := row.Scan(
// 			&u.Id,
// 			&u.FirstName,
// 			&u.LastName,
// 			&u.Email,
// 			&u.Password,
// 			&u.IsActive,
// 			&u.CreatedAt,
// 			&u.UpdatedAt,
// 			&u.Role.Id,
// 			&u.Role.Name,
// 		)
// 		return u, err
// 	})

// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Info("Returning all users")
// 	return users, nil
// }
// func GetOneUser(id string) (dto.OutputUser, error) {
// 	var user dto.OutputUser

// 	err := DbConnection.QueryRow(context.Background(), `
// 	SELECT 
// 		u.id, 
// 		u.first_name, 
// 		u.last_name, 
// 		u.email, 
// 		u.password, 
// 		r.id AS role, 
// 		r.name AS role, 
// 		u.is_active, 
// 		u.created_at, 
// 		u.updated_at
// 	FROM "user" u 
// 	LEFT JOIN role r ON u.role_id = r.id 
// 	WHERE u.id = $1
// 	`, id).Scan(
// 		&user.Id,
// 		&user.FirstName,
// 		&user.LastName,
// 		&user.Email,
// 		&user.Password,
// 		&user.Role.Id,
// 		&user.Role.Name,
// 		&user.IsActive,
// 		&user.CreatedAt,
// 		&user.UpdatedAt,
// 	)
// 	if err != nil {
// 		return dto.OutputUser{}, err
// 	}

// 	log.Infof("Returning user id: %s", id)
// 	return user, nil
// }

// func GetOneByEmailUser(email string) (dto.OutputUser, error) {
// 	var user dto.OutputUser

// 	err := DbConnection.QueryRow(context.Background(), `
// 	SELECT 
// 		u.id, 
// 		u.first_name, 
// 		u.last_name, 
// 		u.email, 
// 		u.password, 
// 		r.id AS role, 
// 		r.name AS role, 
// 		u.is_active, 
// 		u.created_at, 
// 		u.updated_at
// 	FROM "user" u 
// 	LEFT JOIN role r ON u.role_id = r.id 
// 	WHERE u.email = $1
// 	`, email).Scan(
// 		&user.Id,
// 		&user.FirstName,
// 		&user.LastName,
// 		&user.Email,
// 		&user.Password,
// 		&user.Role.Id,
// 		&user.Role.Name,
// 		&user.IsActive,
// 		&user.CreatedAt,
// 		&user.UpdatedAt,
// 	)
// 	if err != nil {
// 		return dto.OutputUser{}, err
// 	}

// 	log.Infof("Returning user email: %s", email)
// 	return user, nil
// }

// func PatchUser(id string, user dto.InputUser) (string, error) {
// 	ch, err := DbConnection.Exec(context.Background(), `
// 	UPDATE "user"
// 	SET  first_name=$1,last_name=$2,email=$3,password=$4,updated_at=$5
// 	WHERE id=$6
// 	`, user.FirstName, user.LastName, user.Email, user.Password,time.Now(), id)
// 	if err != nil {
// 		return "", err
// 	}
// 	if ch.RowsAffected() < 1 {
// 		return "", fiber.ErrNotFound
// 	}

// 	log.Info("User updated, id: ", id)
// 	return id, nil
// }
// func DeleteUser(id string) (string, error) {
// 	ch, err := DbConnection.Exec(context.Background(), `
// 	DELETE from "user"
// 	WHERE id=$1
// 	`, id)
// 	if err != nil {
// 		return "", err
// 	}
// 	if ch.RowsAffected() < 1 {
// 		return "", fiber.ErrNotFound
// 	}
// 	log.Info("User deleted, id: ", id)
// 	return id, nil
// }

// func ActivateUser(id string) ( error) {
// 	ch, err := DbConnection.Exec(context.Background(), `
// 	UPDATE "user"
// 	SET  is_active=$1, updated_at=$2
// 	WHERE id=$3
// 	`, true,time.Now(),id)
// 	if err != nil {
// 		return  err
// 	}
// 	if ch.RowsAffected() < 1 {
// 		return  fiber.ErrNotFound
// 	}

// 	log.Info("User updated, id: ", id)
// 	return nil
// }
