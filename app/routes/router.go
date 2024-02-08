package routes

import (
	"my-tourist-ticket/app/middlewares"
	"my-tourist-ticket/utils/cloudinary"
	"my-tourist-ticket/utils/encrypts"

	ud "my-tourist-ticket/features/user/data"
	uh "my-tourist-ticket/features/user/handler"
	us "my-tourist-ticket/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()

	userData := ud.New(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService, cloudinaryUploader)

	// define routes/ endpoint USERS
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.RegisterUser)
	e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
}
