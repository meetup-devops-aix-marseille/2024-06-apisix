package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/users", getUsers)
	app.Get("/users/:id", getUser)
	app.Post("/users", createUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)

	app.Listen(":3000")
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	for _, user := range users {
		if id == user.ID {
			return c.JSON(user)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
}

func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	users = append(users, *user)
	return c.JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	for i, u := range users {
		if id == u.ID {
			users[i] = *user
			return c.JSON(user)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
}

func deleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	for i, user := range users {
		if id == user.ID {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(fiber.Map{"status": "User deleted"})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
}
