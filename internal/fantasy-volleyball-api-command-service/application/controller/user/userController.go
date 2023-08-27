package user

import (
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/controller/user/request"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/handler/user/create"
	logger "fantasy-volleyball-api/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type IUserController interface {
	CreateUser(ctx *gin.Context)
}

type userController struct {
	commandHandler create.IUserCreateCommandHandler
	logger         logger.Logger
}

func NewUserController(commandHandler create.IUserCreateCommandHandler) IUserController {
	return &userController{
		commandHandler: commandHandler,
		logger:         logger.GetLogger(reflect.TypeOf((*userController)(nil))),
	}
}

func (controller *userController) CreateUser(ctx *gin.Context) {
	var createUserRequest request.CreateUserRequest

	err := ctx.BindJSON(&createUserRequest)
	if err != nil {
		fmt.Println("Request json bind error occurred")
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	controller.logger.InfoWithContext(ctx, "userController.CreateUser INFO - Started with request: %#v", createUserRequest)

	err = controller.commandHandler.Handle(ctx, createUserRequest.ToCommand())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}
