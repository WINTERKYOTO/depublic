package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/depublic/depublic/internal/config"
	"github.com/depublic/depublic/internal/repository"
	"github.com/depublic/depublic/internal/util"
	"github.com/dgrijalva/jwt-go"
)

// UserHandler is a struct that holds the handlers for user-related requests.
type UserHandler struct {
	config *config.Config
	repo   repository.Repository
}

// NewUserHandler returns a new UserHandler instance.
func NewUserHandler(config *config.Config, repo repository.Repository) *UserHandler {
	return &UserHandler{
		config: config,
		repo:   repo,
	}
}

// GetUserHandler handles the `GET /users/:id` request.
func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request path
	userID := util.ParseUint(r.URL.PathParam("id"))
	if userID == 0 {
		util.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Get the user
	user, err := h.repo.GetUser(userID)
	if err != nil {
		if err == repository.ErrNotFound {
			util.Error(w, http.StatusNotFound, "User not found")
		} else {
			util.Error(w, http.StatusInternalServerError, "Failed to get user: %v", err)
		}
		return
	}

	// Return the user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUserHandler handles the `PUT /users/:id` request.
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request path
	userID := util.ParseUint(r.URL.PathParam("id"))
	if userID == 0 {
		util.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Validate the request body
	var updateRequest struct {
		FullName string `json:"full_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		util.Error(w, http.StatusBadRequest, "Invalid request body: %v", err)
		return
	}

	// Get the user
	user, err := h.repo.GetUser(userID)
	if err != nil {
		if err == repository.ErrNotFound {
			util.Error(w, http.StatusNotFound, "User not found")
		} else {
			util.Error(w, http.StatusInternalServerError, "Failed to get user: %v", err)
		}
		return
	}

	// Update the user
	user.FullName = updateRequest.FullName
	if err := h.repo.UpdateUser(user); err != nil {
		util.Error(w, http.StatusInternalServerError, "Failed to update user: %v", err)
		return
	}

	// Return the updated user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUserHandler handles the `DELETE /users/:id` request.
func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request path
	userID := util.ParseUint(r.URL.PathParam("id"))
	if userID == 0 {
		util.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Delete the user
	if err := h.repo.DeleteUser(userID); err != nil {
	
