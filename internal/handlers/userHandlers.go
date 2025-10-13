package handlers

import (
	"context"
	"project_x/internal/userService"
	"project_x/internal/web/users"
)

type UserHandler struct {
	service userService.UserService
}

// GetUsersId implements users.StrictServerInterface.
func (h *UserHandler) GetUsersId(ctx context.Context, request users.GetUsersIdRequestObject) (users.GetUsersIdResponseObject, error) {
	user, err := h.service.GetUserByID(request.Id)
	if err != nil {
		return users.GetUsersId404Response{}, nil
	}

	response := users.GetUsersId200JSONResponse{
		Id:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
	return response, nil
}

func NewUserHandler(s userService.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       usr.ID,
			Email:    usr.Email,
			Password: usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	createdUser, err := h.service.CreateUser(userRequest.Email, userRequest.Password)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       createdUser.ID,
		Email:    createdUser.Email,
		Password: createdUser.Password,
	}
	return response, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userRequest := request.Body
	updatedUser, err := h.service.UpdateUser(request.Id, userRequest.Email, userRequest.Password)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       updatedUser.ID,
		Email:    updatedUser.Email,
		Password: updatedUser.Password,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.service.DeleteUser(request.Id)
	if err != nil {
		return users.DeleteUsersId404Response{}, nil
	}
	return users.DeleteUsersId204Response{}, nil
}
