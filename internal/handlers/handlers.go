package handlers

import "guitar-go/internal/services"

type Handlers struct {
	authService services.AuthService
	userService services.UserService
}

func NewHandlers(authSvc services.AuthService, userSvc services.UserService) *Handlers {
	return &Handlers{
		authService: authSvc,
		userService: userSvc,
	}
}
