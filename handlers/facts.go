package handlers

import (
	"github.com/10wpressure/simple-rest-api/database"
	"github.com/10wpressure/simple-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&fact)

	return ConfirmationView(c)
}

func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Fact added successfully",
		"Subtitle": "Add more wonderful facts to the list!",
	})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Fact",
		"Subtitle": "Add a cool fact!",
	})
}

func ListFacts(c *fiber.Ctx) error {
	var facts []models.Fact

	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title":    "Simple Rest-API app",
		"Subtitle": "Facts for fun times with friends",
		"Facts":    facts,
	})
}

func DeleteFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")
	database.DB.Db.First(&fact, id)
	if fact.Question == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "No fact found with ID"})
	}
	database.DB.Db.Delete(&id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Fact successfully deleted"})
}
