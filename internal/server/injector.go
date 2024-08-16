//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"waizly/internal/config"
	"waizly/internal/delivery/http/controller"
	"waizly/internal/delivery/http/middleware"
	"waizly/internal/delivery/http/route"
	"waizly/internal/repository"
	"waizly/internal/usecase"
)

var configSet = wire.NewSet(config.NewViper, config.NewLogger, config.NewValidator, config.NewGin, config.NewApp, config.NewDatabase, config.NewJwtWrapper)
var controllerSet = wire.NewSet(
	controller.NewWelcomeController,
	controller.NewAuthController,
	controller.NewTaxController,
	controller.NewCurrencyController,
	controller.NewItemController,
	controller.NewCustomerController,
	controller.NewInvoiceController,
	controller.NewUserController,
)
var repositorySet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewCustomerRepository,
	repository.NewInvoiceRepository,
	repository.NewInvoiceItemRepository,
	repository.NewItemRepository,
	repository.NewTaxRepository,
	repository.NewCurrencyRepository,
)
var usecaseSet = wire.NewSet(
	usecase.NewUserUseCase,
	usecase.NewTaxUseCase,
	usecase.NewCurrencyUseCase,
	usecase.NewItemUseCase,
	usecase.NewCustomerUseCase,
	usecase.NewInvoiceUseCase,
)
var middlewareSet = wire.NewSet(middleware.NewAuthMiddleware)

func InitializeServer() *config.App {
	wire.Build(configSet, controllerSet, middlewareSet, repositorySet, usecaseSet, route.NewConfigRoute)
	return nil
}
