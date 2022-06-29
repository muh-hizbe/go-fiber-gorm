package route

import (
	"go-fiber-gorm/config"
	"go-fiber-gorm/handler"
	"go-fiber-gorm/middleware"
	"go-fiber-gorm/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", "public")

	r.Post("/login", handler.LoginHandler)

	r.Get("/user", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Delete("/user/:id", handler.UserHandlerDelete)

	r.Post("/book", utils.HandleSingleFile, handler.BookHandlerCreate)

	r.Post("/gallery", utils.HandleMultipleFile, handler.PhotoHandlerCreate)
	r.Delete("/gallery/:id", utils.HandleMultipleFile, handler.PhotoHandlerDelete)
}
