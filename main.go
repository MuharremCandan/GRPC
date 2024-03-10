package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"test-grpc-project/pkg/config"
	"test-grpc-project/pkg/db"
	"test-grpc-project/pkg/gapi"
	"test-grpc-project/pkg/model"
	"test-grpc-project/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {

	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("error loading config file %s:", err)
	}

	db := db.ConnectDb(&config)

	gapi.NewServer(&config, db).NewgRpcServer()

}

func runHttpServer(db *gorm.DB, config config.Config) {
	app := fiber.New()

	app.Get("/greet", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Post("/createuser", func(c *fiber.Ctx) error {
		var user model.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		user.Password = string(hashedPassword)
		if err := db.Create(&user).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				map[string]string{
					"error": err.Error(),
				},
			)
		}
		user.CreatedAt = time.Now()

		return c.Status(200).JSON(user)
	})

	app.Delete("/deleteuser:id", func(c *fiber.Ctx) error {

		id := c.Params("id")
		idToi, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				map[string]string{
					"error": err.Error(),
				},
			)
		}

		if err := db.Where("id = ?", idToi).Delete(&model.User{}).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				map[string]string{
					"error": err.Error(),
				},
			)
		}

		return c.Status(200).SendString("User deleted!")
	})

	app.Listen(net.JoinHostPort(config.HttpServer.Host, config.HttpServer.Port))
}
