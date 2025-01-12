package routes

import(
	"crud-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App){
	app.Get("/api/v1/products", controllers.GetAllProduct)
	app.Get("/api/v1/products/:id", controllers.GetProduct)
	app.Post("/api/v1/products", controllers.CreateProduct)
	app.Patch("/api/v1/products/:id", controllers.UpdateProduct)
	app.Delete("/api/v1/products/:id", controllers.DeleteProduct)
}