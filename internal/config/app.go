package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log/slog"
	"reflect"
	"strings"
	"waizly/internal/delivery/http/route"
)

type App struct {
	GinConfig *GinConfig
	Route     *route.ConfigRoute
	Config    *viper.Viper
	DB        *gorm.DB
	Logger    *slog.Logger
}

func NewApp(ginConfig *GinConfig, route *route.ConfigRoute, config *viper.Viper, db *gorm.DB, logger *slog.Logger) *App {
	return &App{GinConfig: ginConfig, Route: route, Config: config, DB: db, Logger: logger}
}

func (a *App) Serve() error {
	var app = a.GinConfig.Setup()

	// manually recover from panic
	app.Use(gin.Recovery())

	// register validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// register default error tag name to use json tag name not struct field
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			// skip if tag key says it should be ignored
			if name == "-" {
				return ""
			}
			return name
		})
	}

	a.Route.Setup(app)

	appPort := a.Config.GetInt("APP_PORT")
	a.Logger.Info("server running on port " + fmt.Sprintf("%d", appPort))

	return app.Run(fmt.Sprintf(":%d", appPort))
}
