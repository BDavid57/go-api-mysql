package main

import (
	"log"

	"github.com/BDavid57/go-api-mysql/gorm/config"
	"github.com/BDavid57/go-api-mysql/gorm/controllers"
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	config.Connect()

    app.Get("/dogs", controllers.GetDogs)
    app.Get("/dogs/:id", controllers.GetDog)
    app.Post("/dogs", controllers.AddDog)
    app.Put("/dogs/:id", controllers.UpdateDog)
    app.Delete("/dogs/:id", controllers.RemoveDog)

	log.Fatal(app.Listen(":3000"))
}
