package dto

import "time"


type Otp struct {
	UserId  string `json:"user_id"`
	Code     string `json:"code"`
	Deadline time.Time `json:"deadline"`
}