package routes

import (
	"my-tourist-ticket/app/middlewares"
	"my-tourist-ticket/utils/cloudinary"
	"my-tourist-ticket/utils/encrypts"

	ud "my-tourist-ticket/features/user/data"
	uh "my-tourist-ticket/features/user/handler"
	us "my-tourist-ticket/features/user/service"

	cd "my-tourist-ticket/features/city/data"
	ch "my-tourist-ticket/features/city/handler"
	cs "my-tourist-ticket/features/city/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()

	userData := ud.New(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService, cloudinaryUploader)

	cityData := cd.NewCity(db, cloudinaryUploader)
	cityService := cs.NewCity(cityData)
	cityHandlerAPI := ch.NewCity(cityService)

	// define routes/ endpoint USERS
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.RegisterUser)
	e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())
	e.GET("/users/admin", userHandlerAPI.GetAdminUserData, middlewares.JWTMiddleware())
	e.PUT("/users/admin/:id", userHandlerAPI.UpdateUserPengelolaById, middlewares.JWTMiddleware())

	//define routes/ endpoint CITY
	e.POST("/citys", cityHandlerAPI.CreateCity, middlewares.JWTMiddleware())
	e.PUT("/citys/:city_id", cityHandlerAPI.UpdateCity, middlewares.JWTMiddleware())
	e.GET("/citys/:city_id", cityHandlerAPI.GetCityById)
}
