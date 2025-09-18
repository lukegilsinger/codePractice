// ===================================================================
// internal/handlers/auth.go (UPDATED with logging)
// ===================================================================
package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"todo-list-2/internal/auth"
	"todo-list-2/internal/database"
	"todo-list-2/internal/logger"
	"todo-list-2/internal/middleware"
	"todo-list-2/internal/models"
)

type AuthHandler struct {
	db     *database.DB
	logger *logger.Logger
}

func NewAuthHandler(db *database.DB, logger *logger.Logger) *AuthHandler {
	return &AuthHandler{db: db, logger: logger}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Warn("Invalid JSON in register request", "error", err.Error())
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Username == "" || req.Email == "" || req.Password == "" {
		h.logger.LogAuthFailure(req.Username, "register", "missing required fields", getClientIP(r))
		http.Error(w, "Username, email, and password are required", http.StatusBadRequest)
		return
	}

	if len(req.Password) < 6 {
		h.logger.LogAuthFailure(req.Username, "register", "password too short", getClientIP(r))
		http.Error(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	h.logger.Info("User registration attempt", "username", req.Username, "email", req.Email)

	user, err := h.db.CreateUser(req)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			h.logger.LogAuthFailure(req.Username, "register", "username already exists", getClientIP(r))
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
		h.logger.Error("Failed to create user", "username", req.Username, "error", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		h.logger.Error("Failed to generate token", "user_id", user.ID, "error", err.Error())
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	h.logger.LogAuthSuccess(user.Username, user.ID, "register")

	response := models.AuthResponse{
		User:  *user,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Warn("Invalid JSON in login request", "error", err.Error())
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		h.logger.LogAuthFailure(req.Username, "login", "missing credentials", getClientIP(r))
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	h.logger.Info("User login attempt", "username", req.Username)

	user, err := h.db.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		h.logger.LogAuthFailure(req.Username, "login", "invalid credentials", getClientIP(r))
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		h.logger.Error("Failed to generate token", "user_id", user.ID, "error", err.Error())
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	h.logger.LogAuthSuccess(user.Username, user.ID, "login")

	response := models.AuthResponse{
		User:  *user,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUserFromContext(r)
	if user == nil {
		h.logger.Error("User not found in context for /me endpoint")
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	fullUser, err := h.db.GetUserByID(user.UserID)
	if err != nil {
		h.logger.Error("Failed to fetch user by ID", "user_id", user.UserID, "error", err.Error())
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	h.logger.Debug("User info retrieved", "user_id", user.UserID, "username", user.Username)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fullUser)
}

func getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	return r.RemoteAddr
}
