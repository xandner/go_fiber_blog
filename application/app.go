package application

import (
	"blog/config"
	"blog/controller"
	"blog/entities"
	"blog/repo"
	"blog/routes"
	"blog/usecase"
	"fmt"
	_ "blog/docs"

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

	// initialize database
	db, err := gorm.Open(sqlite.Open("database.db"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")

	//migrate database
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Article{})
	fmt.Println("Database migrated")

	//initialize repo
	userRepo := repo.NewUserRepo(db)
	articleRepo:= repo.NewArticleRepo(db)

	//initialize usecase
	userUsecase := usecase.NewUserUsecase(userRepo)
	authUsecase:= usecase.NewAuthUsecase(userRepo)
	articleUsecase:= usecase.NewArticleUsecase(articleRepo)

	//initialize controller
	appController := controller.New(userUsecase,articleUsecase,cfg)
	authController:= controller.NewAuthController(authUsecase)
	httpApp.Get("/swagger/*", swagger.HandlerDefault)
	httpApp.Mount("/api/v1/app", routes.AppRoutes(appController,authController))
	httpApp.Listen("127.0.0.1:9090")
}
