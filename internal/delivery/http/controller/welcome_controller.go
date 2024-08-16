package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waizly/internal/helper"
)

type WelcomeController struct {
}

func NewWelcomeController() *WelcomeController {
	return &WelcomeController{}
}

func (controller *WelcomeController) Index(ctx *gin.Context) {
	helper.SuccessResponse[string, any](ctx, http.StatusOK, "Welcome to Waizly API")
}
