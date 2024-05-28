package application

import (
	"blog/config"
	"blog/controller"
	"blog/entities"
	"blog/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Run(cfg *config.Config) {
	// l, zl := logger.New("INFO")
	// go l.Run()

	httpApp := fiber.New()
	httpApp.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	db, err := gorm.Open(sqlite.Open("database.db"))
	if err !=nil{
		panic(err)
	}
	fmt.Println("Database connected")
	db.AutoMigrate(&entities.User{})
	fmt.Println("Database migrated")
	appController := controller.New()
	httpApp.Mount("/api/v1/app", routes.AppRoutes(appController))
	httpApp.Listen("127.0.0.1:9090")
}
