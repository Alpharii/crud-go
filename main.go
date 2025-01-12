package main

import (
	"crud-go/config"
	"crud-go/model"
	"crud-go/routes"
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main(){
	//load env
	if err:= godotenv.Load(); err != nil{
		log.Fatal("Error loading .env file")
	}

	//koneksi database
	config.ConnectDatabase()

    m := gormigrate.New(config.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
        {
            ID: "20240112_create_products_table",
            Migrate: func(tx *gorm.DB) error {
                return tx.AutoMigrate(&model.Product{})
            },
            Rollback: func(tx *gorm.DB) error {
                return tx.Migrator().DropTable(&model.Product{})
            },
        },
    })

	if err := m.Migrate(); err != nil {
        log.Fatalf("Could not migrate: %v", err)
    }
    fmt.Println("Migration successful")

	//inisialisasi fiber
	app := fiber.New()

	//cors middleware
	app.Use(cors.New())

	//routes
	routes.ProductRoutes(app)

	//jalankan server
	fmt.Println("Server started on port 8080")
	log.Fatal(app.Listen(":8080"))
}