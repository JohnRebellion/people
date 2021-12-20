package person

import (
	"github.com/JohnRebellion/go-utils/database"
	fiberUtils "github.com/JohnRebellion/go-utils/fiber"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Address    string `json:"address"`
}

func GetPeople(c *fiber.Ctx) error {
	people := []Person{}
	err := database.DBConn.Find(&people).Error

	if err == nil {
		return c.JSON(people)
	}

	return nil
}

func GetPerson(c *fiber.Ctx) error {
	person := new(Person)
	database.DBConn.Find(&person, c.Params("id"))
	return c.JSON(person)
}

func DeletePerson(c *fiber.Ctx) error {
	database.DBConn.Delete(&Person{}, c.Params("id"))
	return nil
}

func NewPerson(c *fiber.Ctx) error {
	fiberUtils.Ctx.New(c)
	person := new(Person)
	fiberUtils.ParseBody(&person)
	database.DBConn.Create(&person)
	return c.JSON(fiber.Map{"status": "success", "message": "Nailagay na sa database si" + person.Name})
}

func UpdatePerson(c *fiber.Ctx) error {
	fiberUtils.Ctx.New(c)
	person := new(Person)
	fiberUtils.ParseBody(&person)
	database.DBConn.Updates(&person)
	return c.JSON(fiber.Map{"status": "success", "message": "Naiba na sa database si" + person.Name})
}
