package request

import "fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/handler/user/create"

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (request *CreateUserRequest) ToCommand() create.Command {
	return create.Command{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}
}
