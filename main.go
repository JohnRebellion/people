package main

import (
	"people/person"

	"github.com/JohnRebellion/go-utils/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.SQLiteConnect("people.db")
	database.DBConn.AutoMigrate(&person.Person{})
	app := fiber.New()

	apiEndpoint := app.Group("/api")
	v1Endpoint := apiEndpoint.Group("/v1")

	personEndpoint := v1Endpoint.Group("/person")
	personEndpoint.Get("/", person.GetPeople)
	personEndpoint.Get("/:id", person.GetPerson)
	personEndpoint.Delete("/:id", person.DeletePerson)
	personEndpoint.Post("/", person.NewPerson)
	personEndpoint.Put("/", person.UpdatePerson)

	app.Listen(":8888")
}
