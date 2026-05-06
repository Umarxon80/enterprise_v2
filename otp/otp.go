package otp

import (
	"enterprise_v2/db"
	"enterprise_v2/dto"
	"enterprise_v2/email"
	"fmt"
	"math/rand"
	"time"
)

func GenerateOtp() (string, time.Time) {
	return fmt.Sprintf("%06d", rand.Intn(1000000)), time.Now().Add(30 * time.Second)
}

func MakeAndSendOtp(user_id string,user_email string)(error){
	code,deadline:=GenerateOtp()
	otp:=dto.Otp{
		UserId: user_id,
		Code: code,
		Deadline: deadline,
	}
	if err:=db.DeleteOtp(user_id);err!=nil{
		return fmt.Errorf("Error deleting old_otps %w",err)
	}
	if err:=db.CreateOtp(otp);err!=nil{
		return fmt.Errorf("Error creating otp %w",err)
	}
	if err:=email.SendOtp(user_email,"Bob","Email verification","Your email was registered on enterprise.us. If it was you please click botton below and varify your email ","http://localhost:8080/verify_email/"+otp.Code);err!=nil {
		return fmt.Errorf("Error sending otp %w",err)
	}
	return nil
}
