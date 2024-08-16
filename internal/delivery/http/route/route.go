package route

import (
	"github.com/gin-gonic/gin"
	"waizly/internal/delivery/http/controller"
	"waizly/internal/delivery/http/middleware"
)

type ConfigRoute struct {
	WelcomeController  *controller.WelcomeController
	AuthController     *controller.AuthController
	AuthMiddleware     *middleware.AuthMiddleware
	TaxController      *controller.TaxController
	CurrencyController *controller.CurrencyController
	ItemController     *controller.ItemController
	CustomerController *controller.CustomerController
	InvoiceController  *controller.InvoiceController
	UserController     *controller.UserController
}

func apiGroup(app *gin.Engine, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return app.Group("/api", handlers...)
}

func NewConfigRoute(
	welcomeController *controller.WelcomeController,
	authController *controller.AuthController,
	authMiddleware *middleware.AuthMiddleware,
	taxController *controller.TaxController,
	currencyController *controller.CurrencyController,
	itemController *controller.ItemController,
	customerController *controller.CustomerController,
	invoiceController *controller.InvoiceController,
	userController *controller.UserController,
) *ConfigRoute {
	return &ConfigRoute{
		WelcomeController:  welcomeController,
		AuthController:     authController,
		AuthMiddleware:     authMiddleware,
		TaxController:      taxController,
		CurrencyController: currencyController,
		ItemController:     itemController,
		CustomerController: customerController,
		InvoiceController:  invoiceController,
		UserController:     userController,
	}
}

func (c *ConfigRoute) Setup(app *gin.Engine) {
	c.guestRoute(app)
	c.protectedRoute(app)
}

func (c *ConfigRoute) guestRoute(app *gin.Engine) {
	api := apiGroup(app)
	{
		api.GET("/", c.WelcomeController.Index)
		api.POST("/login", c.AuthController.Login)
		api.POST("/register", c.AuthController.Register)
	}
}

func (c *ConfigRoute) protectedRoute(app *gin.Engine) {
	api := apiGroup(app, c.AuthMiddleware.TokenAuthorization)
	{
		userGroup := api.Group("/user")
		{
			userGroup.GET("/me", c.UserController.GetUser)
		}

		taxGroup := api.Group("/tax")
		{
			taxGroup.GET("/", c.TaxController.Index)
			taxGroup.POST("/", c.TaxController.Create)
			taxGroup.GET("/:id", c.TaxController.Show)
			taxGroup.PUT("/:id", c.TaxController.Update)
			taxGroup.DELETE("/:id", c.TaxController.Delete)
		}

		currencyGroup := api.Group("/currency")
		{
			currencyGroup.GET("/", c.CurrencyController.Index)
			currencyGroup.POST("/", c.CurrencyController.Create)
			currencyGroup.GET("/:id", c.CurrencyController.Show)
			currencyGroup.PUT("/:id", c.CurrencyController.Update)
			currencyGroup.DELETE("/:id", c.CurrencyController.Delete)
		}

		itemGroup := api.Group("/item")
		{
			itemGroup.GET("/", c.ItemController.Index)
			itemGroup.POST("/", c.ItemController.Create)
			itemGroup.GET("/:id", c.ItemController.Show)
			itemGroup.PUT("/:id", c.ItemController.Update)
			itemGroup.DELETE("/:id", c.ItemController.Delete)
		}

		customerGroup := api.Group("/customer")
		{
			customerGroup.GET("/", c.CustomerController.Index)
			customerGroup.POST("/", c.CustomerController.Create)
			customerGroup.GET("/:id", c.CustomerController.Show)
			customerGroup.PUT("/:id", c.CustomerController.Update)
			customerGroup.DELETE("/:id", c.CustomerController.Delete)
		}

		invoiceGroup := api.Group("/invoice")
		{
			invoiceGroup.GET("/", c.InvoiceController.Index)
			invoiceGroup.POST("/", c.InvoiceController.Create)
			invoiceGroup.GET("/:id", c.InvoiceController.Show)
			invoiceGroup.PUT("/:id", c.InvoiceController.Update)
			invoiceGroup.DELETE("/:id", c.InvoiceController.Delete)
		}
	}
}
