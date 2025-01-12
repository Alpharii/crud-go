package controllers

import (
	"crud-go/model"
	"fmt"
	"crud-go/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateProduct(c *fiber.Ctx) error{
	product:= model.Product{}
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	config.DB.Create(&product)
	if product.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Sprintf("Failed to create product"),
		})
	}
	return c.JSON(product)
}

func GetAllProduct(c *fiber.Ctx) error{
	var product []model.Product
	if err := config.DB.Find(&product).Error; err != nil {
		return err
	}
	config.DB.Find(&product)
	if len(product) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("No products found"),
		})
	}
	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error{
	id := c.Params("id")
	var product model.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"message": fmt.Sprintf("Product with id %s not found", id),
			})
		}
		return err
	}
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error{
	id:= c.Params("id")
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	if err := config.DB.Model(&product).Where("id = ?", id).Updates(&product).Error; err != nil {
		return err
	}
	return c.JSON(product)
}

func DeleteProduct(c*fiber.Ctx) error{
	id:= c.Params("id")
	var product model.Product
	if err := config.DB.Where("id = ?", id).Delete(&product).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Product with id %s deleted", id),
	})
}