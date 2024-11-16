package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Login(r *gin.Context) {}

func (h *Handler) Register(r *gin.Context) {}

func (h *Handler) VerifyEmail(r *gin.Context) {}

func (h *Handler) ForgotPassword(r *gin.Context) {}

func (h *Handler) ChangeEmail(r *gin.Context) {}

func (h *Handler) ChangePassword(r *gin.Context) {}
