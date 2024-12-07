package controllers

import (
	model "GoArticle/app/model"
	database "GoArticle/config"
	helpers "GoArticle/helpers"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Handler untuk mengembalikan JSON
func GetAPI(c *fiber.Ctx) error {
	response := fiber.Map{
		"message": "Selamat Api Sudah Berjalan Normal Aja",
		"status":  "success",
	}
	return c.JSON(response)
}

func Register(c *fiber.Ctx) error {
	var req model.UserRegister

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if errors := helpers.ValidateStruct(req); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors,
		})
	}

	passwordEns, _ := helpers.HashPassword(req.Password)
	reqData := model.User{
		Username:          req.Username,
		Email:             req.Email,
		DisplayName:       req.DisplayName,
		Bio:               req.Bio,
		Password:          passwordEns,
		ProfilePictureUrl: req.ProfilePictureUrl,
	}

	database.DB.Create(&reqData)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    reqData,
	})
}

func Login(c *fiber.Ctx) error {
	var req model.UserLogin

	// Parsing body JSON ke struct
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	// Validasi input
	if errors := helpers.ValidateStruct(req); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors,
		})
	}

	var DataUser model.User
	result := database.DB.Where("email = ?", req.Email).First(&DataUser)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found, please check your email and password",
			})
		}
		// Error lainnya
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong, please try again later",
		})
	}

	// Validasi password
	if !helpers.CheckPasswordHash(req.Password, DataUser.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	// Membuat klaim token JWT
	claims := jwt.MapClaims{
		"user_id": DataUser.Id,                          // ID pengguna
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Token kadaluarsa 1 jam
	}

	// Mendapatkan secret key
	key := os.Getenv("SC_KEY")
	if key == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Server misconfiguration: missing secret key",
		})
	}

	// Membuat token JWT
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Gunakan HS256 jika kunci simetris
	Token, err := t.SignedString([]byte(key))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Menyimpan token ke database
	tokenData := model.Token{UserId: DataUser.Id, Token: Token}
	if err := database.DB.Create(&tokenData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save token",
		})
	}

	// Response berhasil
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"user":    DataUser,
		"token":   tokenData.Token,
	})
}
