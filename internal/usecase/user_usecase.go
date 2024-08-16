package usecase

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"time"
	"waizly/internal/entity"
	"waizly/internal/helper"
	"waizly/internal/model"
	"waizly/internal/repository"
	"waizly/pkg/apperrors"
)

type UserUseCase struct {
	DB             *gorm.DB
	UserRepository *repository.UserRepository
	Jwt            *model.JwtWrapper
	Validate       *validator.Validate
}

func NewUserUseCase(DB *gorm.DB, userRepository *repository.UserRepository, jwt *model.JwtWrapper, validate *validator.Validate) *UserUseCase {
	return &UserUseCase{DB: DB, UserRepository: userRepository, Jwt: jwt, Validate: validate}
}

func (u *UserUseCase) Register(ctx context.Context, req *model.RegisterUserRequest) (*model.UserResponse, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var user entity.User

	if err := u.UserRepository.FindByEmail(tx, &user, req.Email); err == nil {
		return nil, apperrors.NewAppError(http.StatusConflict, model.ErrorDetail{
			Field:     "email",
			ErrorCode: string(apperrors.EmailExists),
			Param:     "",
			Message:   "email already exists",
		})
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Password = helper.HashPassword(req.Password)

	if err := u.UserRepository.Create(tx, &user); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not create user",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not create user",
		})
	}

	token, err := u.Jwt.GenerateToken(user)
	if err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.TokenError),
			Param:     "",
			Message:   "can not create token",
		})
	}

	return &model.UserResponse{
		User:  &user,
		Token: token,
	}, nil
}

func (u *UserUseCase) Login(ctx *gin.Context, req *model.LoginUserRequest) (*model.UserResponse, *apperrors.AppError) {
	var user entity.User

	if err := u.UserRepository.FindByEmail(u.DB, &user, req.Email); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "email",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "email not found",
		})
	}

	match := helper.ComparePasswordHash(req.Password, user.Password)
	if !match {
		return nil, apperrors.NewAppError(http.StatusBadRequest, model.ErrorDetail{
			Field:     "password",
			ErrorCode: string(apperrors.InvalidPassword),
			Param:     "",
			Message:   "invalid password",
		})
	}

	token, err := u.Jwt.GenerateToken(user)
	if err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.TokenError),
			Param:     "",
			Message:   "can not create token",
		})
	}

	ctx.SetCookie("token", token, int(u.Jwt.ExpirationHours)*int(time.Hour), "/", "localhost", false, true)
	return &model.UserResponse{
		User:  &user,
		Token: token,
	}, nil
}

func (u *UserUseCase) Verify(ctx *gin.Context, req *model.ValidateUserRequest) (*model.UserResponse, *apperrors.AppError) {
	claims, err := u.Jwt.ValidateToken(req.Token)
	if err != nil {
		return nil, apperrors.NewAppError(http.StatusUnauthorized, model.ErrorDetail{
			Field:     "token",
			ErrorCode: string(apperrors.InvalidToken),
			Param:     "",
			Message:   err.Error(),
		})
	}

	var user entity.User

	if err := u.UserRepository.FindByEmail(u.DB, &user, claims.Email); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "user not found",
		})
	}

	return &model.UserResponse{
		User: &user,
	}, nil
}

func (u *UserUseCase) GetUser(ctx *gin.Context, req *model.GetUserRequest) (*model.UserResponse, *apperrors.AppError) {
	var user entity.User
	if err := u.UserRepository.FindById(u.DB, &user, req.ID); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "id",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "user not found",
		})
	}

	return &model.UserResponse{
		User: &user,
	}, nil
}
