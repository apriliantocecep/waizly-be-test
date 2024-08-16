package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"waizly/internal/helper"
	"waizly/internal/model"
	"waizly/internal/usecase"
)

type AuthMiddleware struct {
	UserUseCase *usecase.UserUseCase
}

func NewAuthMiddleware(userUseCase *usecase.UserUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		UserUseCase: userUseCase,
	}
}

func (a *AuthMiddleware) TokenAuthorization(ctx *gin.Context) {
	var errOut []model.ErrorDetail
	authorization := ctx.GetHeader("Authorization")

	if authorization == "" {
		helper.ErrorResponse[any, string](ctx, http.StatusUnauthorized, "Unauthorized")
		ctx.Abort()
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		helper.ErrorResponse[any, string](ctx, http.StatusUnauthorized, "Unauthorized")
		ctx.Abort()
		return
	}

	res, err := a.UserUseCase.Verify(ctx, &model.ValidateUserRequest{Token: token[1]})
	if err != nil {
		errOut = append(errOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errOut)
		ctx.Abort()
		return
	}

	ctx.Set("userId", res.User.ID)
	ctx.Next()
}
