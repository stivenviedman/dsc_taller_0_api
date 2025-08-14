package middlewares

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("KEY_TOKEN"))

func AutValidation(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Falta Token",
		})
	}

	parts := strings.Split(token, " ")
	var StringTok string
	if len(parts) == 2 && parts[0] == "Bearer" {
		StringTok = parts[1]
	}

	Token, err := jwt.Parse(StringTok, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	fmt.Println(Token)
	if claims, ok := Token.Claims.(jwt.MapClaims); ok && Token.Valid {
		fmt.Println("Claims:", claims)
		fmt.Println("Usuario:", claims["username"])
	}

	if err != nil || !Token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token inválido",
		})
	}
	return c.Next()
}

func GenerarToken(username string) (string, error) {

	datos := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, datos)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
