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
