// auth.go
package handlers

import (
	"encoding/json"
	"net/http"

	"guitar-go/pkg/utils"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

// Login godoc
// @Summary User login
// @Description Authenticate user and get JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body Credentials true "Login credentials"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /login [post]
func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	user, err := h.userService.GetUserByUsername(creds.Username)
	if err != nil {
		utils.JSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := h.authService.ComparePassword(user.Password, creds.Password); err != nil {
		utils.JSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token, err := h.authService.GenerateToken(user)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	utils.JSONResponse(w, http.StatusOK, TokenResponse{Token: token})
}
