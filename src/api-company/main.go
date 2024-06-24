package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var companies = []Company{
	{ID: 1, Name: "Acme Corp", Address: "123 Elm Street"},
	{ID: 2, Name: "Globex Inc", Address: "456 Oak Street"},
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/companies", getCompanies)
	app.Get("/companies/:id", getCompany)
	app.Post("/companies", createCompany)
	app.Put("/companies/:id", updateCompany)
	app.Delete("/companies/:id", deleteCompany)

	app.Listen(":3000")
}

func getCompanies(c *fiber.Ctx) error {
	return c.JSON(companies)
}

func getCompany(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for _, company := range companies {
		if id == company.ID {
			return c.JSON(company)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Company not found"})
}

func createCompany(c *fiber.Ctx) error {
	company := new(Company)
	if err := c.BodyParser(company); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	companies = append(companies, *company)
	return c.JSON(company)
}

func updateCompany(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	company := new(Company)
	if err := c.BodyParser(company); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	for i, comp := range companies {
		if id == comp.ID {
			companies[i] = *company
			return c.JSON(company)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Company not found"})
}

func deleteCompany(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i, company := range companies {
		if id == company.ID {
			companies = append(companies[:i], companies[i+1:]...)
			return c.JSON(fiber.Map{"status": "Company deleted"})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Company not found"})
}
