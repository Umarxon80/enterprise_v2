package dto

import "time"

type InputBill struct {
	UserId string `json:"user_id"`  
	Amount     string `json:"amount" `
	Deadline  time.Time `json:"password" `
}

type OutputBill struct {
	Id string `json:"id"`
	UserId string `json:"user_id" `
	Invoice  string `json:"invoice"`
	Deadline     time.Time `json:"deadline" `
	Amount  string `json:"amount" `
	Status  string `json:"status" `
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
