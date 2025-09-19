package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware protects routes
func AuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}

// RoleMiddleware checks if user has the required role
func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		rolesClaim, ok := claims["roles"].([]interface{})
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Invalid roles claim"})
		}

		userRoles := make([]string, len(rolesClaim))
		for i, role := range rolesClaim {
			userRoles[i] = role.(string)
		}

		for _, role := range userRoles {
			if role == requiredRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden: you don't have the required role",
		})
	}
}

// Helper untuk mendapatkan ID user dari token (opsional, bisa digunakan di handler)
func GetUserIDFromToken(c *fiber.Ctx) (int, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID in token")
	}
	return int(id), nil
}
