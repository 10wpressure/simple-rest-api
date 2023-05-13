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
		return NewFactView(c)
	}
	result := database.DB.Db.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}

	return ListFacts(c)
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

func ShowFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact":  fact,
	})
}

func EditFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")

	result := database.DB.Db.Where("id =?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}
	return c.Render("edit", fiber.Map{
		"Title":    "Edit Fact",
		"Subtitle": "Editing your interesting fact",
		"Fact":     fact,
	})
}

func UpdateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}
	result := database.DB.Db.Model(&fact).Where("id = ?", id).Updates(fact)
	if result.Error != nil {
		return EditFact(c)
	}
	return ShowFact(c)
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
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
