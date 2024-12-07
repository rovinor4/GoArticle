package main

import (
	model "GoArticle/app/model"
	database "GoArticle/config"
	"GoArticle/helpers"
	"GoArticle/route"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	database.ConnectDatabase()

	err := database.DB.AutoMigrate(&model.User{}, &model.Category{}, &model.Article{}, &model.Token{})

	if err != nil {
		fmt.Println("Migration failed:", err)
	}

	helpers.RegisterCustomValidations()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})
	route.Api(app)

	_err := app.Listen(":5050")
	if _err != nil {
		fmt.Println("Silahkan ubah port karena sudah pakai pada main.go")
	}

}
