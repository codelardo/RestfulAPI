package routes

import (
	"github.com/TheGolurk/RestfulAPI/user"
	"github.com/labstack/echo"
)

// StartRoutes Inicializa las rutas
func StartRoutes(e *echo.Echo) {
	e.POST("api/v1/user", user.Create) //CREATE
	e.GET("api/v1/user", user.GetAll)  //GetAll Users
}