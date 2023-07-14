package controllers

import (
	"golangGin/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

// grouping the route for anything that has prefix /user routes
func (uc *UserController) RegisterUserRoutes(registerGroup *gin.RouterGroup) {
	userRoute := registerGroup.Group("/user")

	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:name", uc.GetUser)
	userRoute.GET("/get", uc.GetAll)
	userRoute.POST("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:name", uc.DeleteUser)
}
