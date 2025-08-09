// user.go
package handlers

import (
	"encoding/json"
	"net/http"

	"guitar-go/internal/models"
	"guitar-go/pkg/utils"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// GetUsers godoc
// @Summary Get all users
// @Description Get list of all users (admin only)
// @Tags users
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.User
// @Failure 403 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/admin/users [get]
func (h *Handlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Error fetching users")
		return
	}
	utils.JSONResponse(w, http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User details"
// @Success 201 {object} models.User
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /register [post]
func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := h.authService.ValidatePassword(req.Password); err != nil {
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := h.authService.HashPassword(req.Password)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Error hashing password")
		return
	}

	user := &models.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     req.Role,
	}

	if err := h.userService.CreateUser(user); err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	user.Password = "" // Clear password in response
	utils.JSONResponse(w, http.StatusCreated, user)
}
