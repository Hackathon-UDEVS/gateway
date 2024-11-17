package handler

import (
	"context"
	"fmt"
	token2 "github.com/Hackaton-UDEVS/gateway/api/token"
	"github.com/Hackaton-UDEVS/gateway/internal/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"strings"

	"github.com/Hackaton-UDEVS/gateway/internal/genproto/auth"
)

// Login godoc
// @Summary Login user
// @Description Logs in a user and returns user data
// @Accept json
// @Tags user
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
		log.Error("Error while parsing login data", zap.Error(err))
		r.JSON(400, gin.H{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}
	if req.Email == "" && req.Password == "" {
		log.Error("Error while parsing login data", zap.String("email", req.Email))
		r.JSON(400, gin.H{"message": "Username and password are required"})
		return
	}
	fmt.Println(req.Email, "-------", req.Password)

	// Validate fields
	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Password) == "" {
		log.Warn("Validation error - Missing required fields")
		r.JSON(400, gin.H{
			"message": "Email and password are required",
		})
		return
	}

	// Call the user service to authenticate the user
	resp, err := h.Clients.User.Login(context.Background(), &req)
	if err != nil {
		if strings.Contains(err.Error(), "invalid email or password") {
			log.Warn("Invalid email or password", zap.String("email", req.Email))
			r.JSON(400, gin.H{
				"message": "Invalid email or password",
			})
			return
		}

		if strings.Contains(err.Error(), "user not found") {
			log.Warn("User not found", zap.String("email", req.Email))
			r.JSON(400, gin.H{
				"message": "User not found",
			})
			return
		}

		log.Error("Error while logging in", zap.Error(err))
		r.JSON(400, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	// Generate a JWT token for the user
	token := token2.GenereteJWTTokenForUser(resp)
	log.Info("Login success", zap.String("email", req.Email))
	r.JSON(http.StatusOK, gin.H{
		"message": "Success login user",
		"token":   token,
	})
}

// Register godoc
// @Summary Register user
// @Description Registers a new user and returns user data
// @Accept json
// @Produce json
// @Tags user
// @Param register body auth.RegisterUserReq true "Register request"
// @Success 200 {object} auth.RegisterUserRes "Registration successful"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /register [post]
func (h *Handler) Register(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req auth.RegisterUserReq
	if err := r.ShouldBindJSON(&req); err != nil {
		log.Error("Error while parsing data", zap.Error(err))
		r.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data", "error": err.Error()})
		return
	}

	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Password) == "" || strings.TrimSpace(req.Role) == "" {
		log.Warn("Validation error - Missing required fields", zap.Any("request", req))
		r.JSON(http.StatusBadRequest, gin.H{"message": "username or email cannot be empty"})
		return
	}

	if !isValidEmail(req.Email) {
		log.Warn("Validation error - Invalid email format", zap.String("email", req.Email))
		r.JSON(http.StatusBadRequest, gin.H{"message": "invalid email format"})
		return
	}

	resp, err := h.Clients.User.RegisterUser(context.Background(), &req)
	if err != nil {
		log.Error("Error while registering user", zap.Error(err))
		r.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user", "error": err.Error()})
		return
	}
	if resp.Message == "Duplicate email" {
		log.Warn("Validation error - Invalid email format", zap.String("email", req.Email))
		r.JSON(400, gin.H{"message": "Email already exists"})
		return
	}

	logRes := auth.LoginRes{
		UserRes: &auth.UserModel{
			Email: req.Email,
			Role:  req.Role,
		},
	}
	token := token2.GenereteJWTTokenForUser(&logRes)

	log.Info("Register success", zap.String("email", req.Email))
	r.JSON(http.StatusCreated, gin.H{"message": resp.Message, "token": token})
}

// Validate email format using regex
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

//
//// VerifyEmail godoc
//// @Summary Verify user email
//// @Description Verifies the user's email address
//// @Accept json
//// @Produce json
//// @Param verify body auth.VerifyUserReq true "Verify request"
//// @Success 200 {object} auth.VerifyUserRes "Email verification successful"
//// @Failure 400 {object} string "Bad request"
//// @Failure 500 {object} string "Internal server error"
//// @Router /verify-email [post]
//func (h *Handler) VerifyEmail(r *gin.Context) {
//	log, _ := logs.NewLogger()
//
//	var req auth.VerifyUserReq
//
//	if err := r.ShouldBindJSON(&req); err != nil {
//		log.Error("Error while parsing data")
//		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	resp, err := h.Clients.User.VerifyUser(context.Background(), &req)
//	if err != nil {
//		log.Error("Error while verify")
//		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	log.Info("Verify success")
//	r.JSON(http.StatusOK, gin.H{"data": resp})
//}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Retrieves user data by user ID
// @Accept json
// @Produce json
// @Tags user
// @Param id path string true "User ID"
// @Success 200 {object} auth.GetUserByIDRes "User data retrieved successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /user/get-user/{id} [get]
func (h *Handler) GetUserByID(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req auth.GetUserByIDReq

	id := r.Param("id")
	req.Userid = id
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
// @Tags user
// @Param query body auth.GetAllUserReq true "Query parameters"
// @Success 200 {object} auth.GetAllUserRes "List of users retrieved successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /user/getAll-users [get]
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
// @Tags user
// @Param id path string true "User ID"
// @Param update body auth.UpdateUserReq true "Update request"
// @Success 200 {object} auth.UpdateUserRes "User updated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /user/update-user/{id} [put]
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
