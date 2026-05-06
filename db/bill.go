package db

import (
	"context"
	"enterprise_v2/dto"
	"errors"
	"time"
	"github.com/gofiber/fiber/v3/log"
)

// command for enum type
//CREATE TYPE IF NOT EXISTS bill_statuses AS ENUM ('unpaid', 'paid', 'deadline_exeeded');

func createBillTable() error {
	_, err := DbConnection.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS bill(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES "user"(id),
	invoice UUID DEFAULT gen_random_uuid(),
	deadline TIMESTAMP NOT NULL,
	amount VARCHAR(255),
	status bill_statuses NOT NULL DEFAULT 'unpaid',
	reason VARCHAR DEFAULT 'user_accaunt'
	);`)
	if err != nil {
		return err
	}
	return nil
}

func CreateBill(bill dto.InputBill) (string,error) {
	var invoice string
    
	err := DbConnection.QueryRow(context.Background(), `
		INSERT INTO bill (
			user_id,deadline,amount
		) VALUES (
			$1,$2,$3
		) RETURNING invoice 
	`, bill.UserId,bill.Deadline,bill.Amount).Scan(&invoice)
	if err != nil {
		return "",err
	}
	log.Info("Bill created, invoice:",invoice)
	return invoice,nil
}

func CreateTaskBill(bill dto.InputBill) (string,error) {
	var invoice string
    
	err := DbConnection.QueryRow(context.Background(), `
		INSERT INTO bill (
			user_id,deadline,amount,reason
		) VALUES (
			$1,$2,$3,$4
		) RETURNING invoice 
	`, bill.UserId,bill.Deadline,bill.Amount,"task").Scan(&invoice)
	if err != nil {
		return "",err
	}
	log.Info("Bill created, invoice:",invoice)
	return invoice,nil
}

func DeleteBill(user_id string) (error) {
	_, err := DbConnection.Exec(context.Background(), `
	DELETE from bill
	WHERE user_id=$1 AND reason=$2
	`, user_id)
	if err != nil {
		return err
	}
	log.Info("Bill deleted, user_id: ", user_id)
	return nil
}

func PayBill(invoice,user_id string) error {
	var bill dto.OutputBill
	err := DbConnection.QueryRow(context.Background(), `
	SELECT  deadline
	FROM bill WHERE user_id=$1 AND invoice=$2
`, user_id,invoice).Scan(
		&bill.Deadline,
	)
	if err != nil {
		log.Errorf("Error paying bill, %v", err)
		return err
	}
	if time.Now().After(bill.Deadline) {
		deadlineBill(invoice)
		return errors.New("Deadline exeeded")
	}
	payidBill(invoice)
	return nil
}

func deadlineBill(invoice string) {
	_,_ = DbConnection.Exec(context.Background(), `
	UPDATE bill
	SET  status=$1
	WHERE invoice=$2
	`, "deadline_exeeded",invoice)
}
func payidBill(invoice string) {
	_,_ = DbConnection.Exec(context.Background(), `
	UPDATE bill
	SET  status=$1
	WHERE invoice=$2
	`, "paid",invoice)
}
