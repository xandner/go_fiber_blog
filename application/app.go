package application

import (
	"blog/config"
	"blog/controller"
	"blog/entities"
	"blog/repo"
	"blog/routes"
	"blog/usecase"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
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
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
	db.AutoMigrate(&entities.User{})
	fmt.Println("Database migrated")
	userRepo := repo.NewUserRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	appController := controller.New(userUsecase)
	httpApp.Get("/swagger/*", swagger.HandlerDefault)
	httpApp.Mount("/api/v1/app", routes.AppRoutes(appController))
	httpApp.Listen("127.0.0.1:9090")
}
