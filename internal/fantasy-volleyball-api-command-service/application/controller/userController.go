package controller

import (
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/controller/request"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/handler/user/create"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserController interface {
	CreateUser(ctx *gin.Context)
}

type userController struct {
	commandHandler create.IUserCreateCommandHandler
}

func NewUserController(commandHandler create.IUserCreateCommandHandler) IUserController {
	return userController{
		commandHandler: commandHandler,
	}
}

func (controller userController) CreateUser(ctx *gin.Context) {
	fmt.Println("hello controller")
	var createUserRequest request.CreateUserRequest

	err := ctx.BindJSON(&createUserRequest)
	if err != nil {
		fmt.Println("Request json bind error occurred")
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	err = controller.commandHandler.Handle(ctx, createUserRequest.ToCommand())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}
