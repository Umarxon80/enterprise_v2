package db

import (
	"context"
	"enterprise_v2/dto"
	"errors"
	"time"

	"github.com/gofiber/fiber/v3/log"
)

func createOtpTable() error {
	_, err := DbConnection.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS otp(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES "user"(id) UNIQUE,
	code VARCHAR(255) NOT NULL,
	deadline TIMESTAMP NOT NULL
	)`)
	if err != nil {
		return err
	}
	return nil
}

func CreateOtp(otp dto.Otp) (error) {
	var id string

	err := DbConnection.QueryRow(context.Background(), `
		INSERT INTO otp (
			user_id,code,deadline
		) VALUES (
			$1,$2,$3
		) RETURNING id
	`, otp.UserId,otp.Code,otp.Deadline).Scan(&id)
	if err != nil {
		return err
	}
	log.Info("Otp created, id:",id)
	return nil
}
func DeleteOtp(user_id string) (error) {
	_, err := DbConnection.Exec(context.Background(), `
	DELETE from otp
	WHERE user_id=$1
	`, user_id)
	if err != nil {
		return err
	}
	log.Info("Otp deleted, user_id: ", user_id)
	return nil
}

func CheckOtp(user_id string, code string) error {
	var otp dto.Otp
	err := DbConnection.QueryRow(context.Background(), `
	SELECT  user_id, code, deadline
	FROM otp WHERE user_id=$1
`, user_id).Scan(
		&otp.UserId,
		&otp.Code,
		&otp.Deadline,
	)
	if err != nil {
		log.Errorf("Error getting otp, %v", err)
		return err
	}
	if time.Now().After(otp.Deadline) {
		return errors.New("Deadline exeeded")
	}
	if code != otp.Code {
		return errors.New("Wrong OTP")
	}
	return nil
}
