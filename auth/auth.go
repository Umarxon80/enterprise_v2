package auth

import (
	"enterprise_v2/db"
	"enterprise_v2/dto"
	"enterprise_v2/helper"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt/v5"
)



var secretKey = []byte(os.Getenv("JWT_SECRET"))

func LogIn(ctx fiber.Ctx) error {
	var auth dto.Auth
	if err:=ctx.Bind().JSON(&auth);err!=nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validation
	if err:=helper.Validate(auth);err!=nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Finding user
	user,err:=db.GetOneByEmailUser(auth.Email)
	if err!=nil {
		log.Error("Error getting user ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Checking password
	if err:=helper.CheckPassword(auth.Password,user.Password); err!=nil {
		log.Error("Wrong password ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong email or password"})
	}
	token, err := generateToken(user.Id, user.Email,user.Role.Name)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not generate token"})
	}

	return ctx.JSON(fiber.Map{"token": token})
	
}

func generateToken(userid,email,role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    userid,
		"email": email,
		"role": role,
		"exp":   time.Now().Add(time.Hour * 10).Unix(),
		"iat":   time.Now().Unix(),
	})
	return token.SignedString(secretKey)
}