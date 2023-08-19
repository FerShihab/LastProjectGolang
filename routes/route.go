package routes

import (
	"lastProject/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {

	auth := e.Group("")
	auth.GET("/cars", controllers.GetCarsController)
	auth.GET("/cars/:id", controllers.GetDetailCarController)
	auth.DELETE("/cars/:id", controllers.DeleteDetailCarController)
	auth.POST("/cars", controllers.AddCar)
}
