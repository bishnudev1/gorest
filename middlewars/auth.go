package middlewars

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *fiber.Ctx) error {
	cookies := c.Cookies("gorest")

	if cookies == "" {
		return c.Status(401).JSON(fiber.Map{"status": 401, "message": "Unauthorized"})
	} else {
		_, err := jwt.Parse(cookies, func(token *jwt.Token) (interface{}, error) {
			c.Locals("gorest", token)
			return []byte("bishnudevkhutiasecretkey"), nil
		})

		if err != nil {
			return c.Status(401).JSON(fiber.Map{"status": 401, "message": "You are not authorized"})
		}
		return c.Next()
	}
}
