package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func RequireAuth(ctx fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		log.Error("No token provided")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "no token provided 1"})
	}

	devidedAuth := strings.SplitN(authHeader, " ", 2)
	if len(devidedAuth) != 2 || devidedAuth[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token format"})
	}
	claims, err := checkToken(devidedAuth[1])
	if err != nil {
		log.Error("No token provided")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "no token provided 2", "err": err.Error()})
	}

	ctx.Locals("userid", claims["id"])
	ctx.Locals("email", claims["email"])
	ctx.Locals("role", claims["role"])
	return ctx.Next()
}
func checkToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil

	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
