package handlers

import (
	"context"
	"time"

	"GoTST/internal/userService"
	"GoTST/internal/web/users"
)

type UserHandler struct {
	service userService.Service
}

func NewUserHandler(service userService.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(ctx context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	usersList, err := h.service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := make([]users.User, 0) // 👈 создаём пустой слайс, не nil

	for _, u := range usersList {
		id := int(u.ID)
		createdAt := u.CreatedAt
		updatedAt := u.UpdatedAt
		var deletedAt *time.Time
		if u.DeletedAt.Valid {
			deletedAt = &u.DeletedAt.Time
		}

		response = append(response, users.User{
			Id:        &id,
			Email:     &u.Email,
			Password:  &u.Password,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
			DeletedAt: deletedAt,
		})
	}

	return users.GetUsers200JSONResponse(response), nil
}


func (h *UserHandler) PostUser(ctx context.Context, req users.PostUserRequestObject) (users.PostUserResponseObject, error) {
	input := req.Body
	user := &userService.User{
		Email:    input.Email,
		Password: input.Password,
	}

	if err := h.service.CreateUser(user); err != nil {
		return nil, err
	}

	return users.PostUser201Response{}, nil
}

func (h *UserHandler) PatchUserByID(ctx context.Context, req users.PatchUserByIDRequestObject) (users.PatchUserByIDResponseObject, error) {
	input := req.Body
	user := &userService.User{
		ID:       uint(req.Id),
		Email:    input.Email,
		Password: input.Password,
	}

	if err := h.service.UpdateUser(user); err != nil {
		return nil, err
	}

	id := int(user.ID)
	createdAt := user.CreatedAt
	updatedAt := user.UpdatedAt
	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}

	return users.PatchUserByID200JSONResponse{
		Id:        &id,
		Email:     &user.Email,
		Password:  &user.Password,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		DeletedAt: deletedAt,
	}, nil
}

func (h *UserHandler) DeleteUserByID(ctx context.Context, req users.DeleteUserByIDRequestObject) (users.DeleteUserByIDResponseObject, error) {
	if err := h.service.DeleteUser(uint(req.Id)); err != nil {
		return nil, err
	}
	return users.DeleteUserByID204Response{}, nil
}
