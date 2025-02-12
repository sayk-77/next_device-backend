package tools

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
)

func GetImageProduct(app *fiber.App) {
	imageDir := "./images_product"

	app.Get("/api/images/product/:name", func(c *fiber.Ctx) error {
		filename := c.Params("name")

		imagePath := filepath.Join(imageDir, filename)

		absPath, err := filepath.Abs(imagePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		fmt.Println("Looking for image at:", absPath)

		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).SendString("Image not found")
		}

		return c.SendFile(absPath)
	})
}

func GetImageBrand(app *fiber.App) {
	imageDir := "./images_brand"

	app.Get("/api/images/brand/:name", func(c *fiber.Ctx) error {
		filename := c.Params("name")

		imagePath := filepath.Join(imageDir, filename)

		absPath, err := filepath.Abs(imagePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		fmt.Println("Looking for image at:", absPath)

		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).SendString("Image not found")
		}

		return c.SendFile(absPath)
	})
}

func GetImageBanner(app *fiber.App) {
	imageDir := "./images_banners"

	app.Get("/api/images/banner/:name", func(c *fiber.Ctx) error {
		filename := c.Params("name")

		imagePath := filepath.Join(imageDir, filename)

		absPath, err := filepath.Abs(imagePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		fmt.Println("Looking for image at:", absPath)

		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).SendString("Image not found")
		}

		return c.SendFile(absPath)
	})
}

func GetImageCategory(app *fiber.App) {
	imageDir := "./images_category"

	app.Get("/api/images/category/:name", func(c *fiber.Ctx) error {
		filename := c.Params("name")

		imagePath := filepath.Join(imageDir, filename)

		absPath, err := filepath.Abs(imagePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		fmt.Println("Looking for image at:", absPath)

		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).SendString("Image not found")
		}

		return c.SendFile(absPath)
	})
}

func GetReviewImage(app *fiber.App) {
	imageDir := "./images_review"

	app.Get("/api/images/review/:name", func(c *fiber.Ctx) error {
		filename := c.Params("name")

		imagePath := filepath.Join(imageDir, filename)

		absPath, err := filepath.Abs(imagePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		fmt.Println("Looking for image at:", absPath)

		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).SendString("Image not found")
		}

		return c.SendFile(absPath)
	})
}
