package create

import (
	"context"
	"errors"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/repository/user"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/domain"
	"fantasy-volleyball-api/pkg/utils"
)

type IUserCreateCommandHandler interface {
	Handle(ctx context.Context, command Command) error
}

type userCreateCommandHandler struct {
	userRepository repository.IUserRepository
}

func NewUserCreateCommandHandler(userRepository repository.IUserRepository) IUserCreateCommandHandler {
	return userCreateCommandHandler{
		userRepository: userRepository,
	}
}

func (handler userCreateCommandHandler) Handle(ctx context.Context, command Command) error {
	err := handler.userRepository.Upsert(ctx, handler.buildEntityFromCommand(command))

	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (handler userCreateCommandHandler) buildEntityFromCommand(command Command) *domain.User {
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
