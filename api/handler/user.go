package handler

import (
	"context"
	"gateway/internal/logs"
	"github.com/gin-gonic/gin"
	"net/http"

	"gateway/internal/genproto/auth"
)

// Login godoc
// @Summary Login user
// @Description Logs in a user and returns user data
// @Accept json
// @Produce json
// @Param login body auth.LoginReq true "Login request"
// @Success 200 {object} auth.LoginRes "Login successful"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /login [post]
func (h *Handler) Login(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req auth.LoginReq

	if err := r.ShouldBindJSON(&req); err != nil {
		log.Error("Error while parsing data")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Clients.User.Login(context.Background(), &req)
	if err != nil {
		log.Error("Error while login")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Login success")
	r.JSON(http.StatusOK, gin.H{"data": resp})
}

// Register godoc
// @Summary Register user
// @Description Registers a new user and returns user data
// @Accept json
// @Produce json
// @Param register body auth.RegisterUserReq true "Register request"
// @Success 200 {object} auth.RegisterUserRes "Registration successful"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /register [post]
func (h *Handler) Register(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req auth.RegisterUserReq

	if err := r.ShouldBindJSON(&req); err != nil {
		log.Error("Error while parsing data")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Clients.User.RegisterUser(context.Background(), &req)
	if err != nil {
		log.Error("Error while register")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Register success")
	r.JSON(http.StatusOK, gin.H{"data": resp})
}

// VerifyEmail godoc
// @Summary Verify user email
// @Description Verifies the user's email address
// @Accept json
// @Produce json
// @Param verify body auth.VerifyUserReq true "Verify request"
// @Success 200 {object} auth.VerifyUserRes "Email verification successful"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /verify-email [post]
func (h *Handler) VerifyEmail(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req auth.VerifyUserReq

	if err := r.ShouldBindJSON(&req); err != nil {
		log.Error("Error while parsing data")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Clients.User.VerifyUser(context.Background(), &req)
	if err != nil {
		log.Error("Error while verify")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Verify success")
	r.JSON(http.StatusOK, gin.H{"data": resp})
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Retrieves user data by user ID
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} auth.GetUserByIDRes "User data retrieved successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /get-user/{id} [get]
func (h *Handler) GetUserByID(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req auth.GetUserByIDReq

	if err := r.ShouldBindJSON(&req); err != nil {
		log.Error("Error while parsing data")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Clients.User.GetUserByID(context.Background(), &req)
	if err != nil {
		log.Error("Error while getUserByID")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Get User success")
	r.JSON(http.StatusOK, gin.H{"data": resp})
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieves a list of all users
// @Accept json
// @Produce json
// @Param query body auth.GetAllUserReq true "Query parameters"
// @Success 200 {object} auth.GetAllUserRes "List of users retrieved successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /getAll-users [get]
func (h *Handler) GetAllUsers(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req auth.GetAllUserReq

	if err := r.ShouldBindQuery(&req); err != nil {
		log.Error("Error while parsing data")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Clients.User.GetAllUsers(context.Background(), &req)
	if err != nil {
		log.Error("Error while getAllUsers")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Get All Users success")
	r.JSON(http.StatusOK, gin.H{"data": resp})
}

// UpdateUser godoc
// @Summary Update user
// @Description Updates user data by ID
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param update body auth.UpdateUserReq true "Update request"
// @Success 200 {object} auth.UpdateUserRes "User updated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /update-user/{id} [put]
func (h *Handler) UpdateUser(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req auth.UpdateUserReq

	if err := r.ShouldBindJSON(&req); err != nil {
		log.Error("Error while parsing data")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Clients.User.UpdateUser(context.Background(), &req)
	if err != nil {
		log.Error("Error while update")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Update success")
	r.JSON(http.StatusOK, gin.H{"data": resp})
}
