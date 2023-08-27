package create

import (
	"context"
	"errors"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/repository/user"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/domain"
	logger "fantasy-volleyball-api/pkg/log"
	"fantasy-volleyball-api/pkg/utils"
	"reflect"
)

type IUserCreateCommandHandler interface {
	Handle(ctx context.Context, command Command) error
}

type userCreateCommandHandler struct {
	userRepository repository.IUserRepository
	logger         logger.Logger
}

func NewUserCreateCommandHandler(userRepository repository.IUserRepository) IUserCreateCommandHandler {
	return &userCreateCommandHandler{
		userRepository: userRepository,
		logger:         logger.GetLogger(reflect.TypeOf((*userCreateCommandHandler)(nil))),
	}
}

func (handler *userCreateCommandHandler) Handle(ctx context.Context, command Command) error {
	handler.logger.InfoWithContext(ctx, "userCreateCommandHandler.Handle INFO - Started command: %#v", command)
	err := handler.userRepository.Upsert(ctx, handler.buildEntityFromCommand(command))

	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (handler *userCreateCommandHandler) buildEntityFromCommand(command Command) *domain.User {
	epochNow, _ := utils.EpochNow()
	entity := domain.User{
		FirstName:  command.FirstName,
		LastName:   command.LastName,
		Email:      command.Email,
		Password:   command.Password,
		CreatedAt:  epochNow,
		ModifiedAt: epochNow,
	}
	return &entity

}
