package test

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"waizly/internal/server"
)

var app *gin.Engine

var db *gorm.DB

var viperConfig *viper.Viper

func init() {
	initServer := server.InitializeServer()

	viperConfig = initServer.Config
	db = initServer.DB
	app = initServer.GinConfig.Setup()
	initServer.Route.Setup(app)
}
