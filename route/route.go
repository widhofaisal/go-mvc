package route

import (
	"golang-echo-gorm/controller"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", controller.GetHelloWorld)

	// Route user
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	// Route blog
	e.GET("/blogs", controller.GetBlogsController)
	e.GET("/blogs/:id", controller.GetBlogController)
	e.POST("/blogs", controller.CreateBlogController)
	e.PUT("/blogs/:id", controller.UpdateBlogController)
	e.DELETE("/blogs/:id", controller.DeleteBlogController)

	return e
}
