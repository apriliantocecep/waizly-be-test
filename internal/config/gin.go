package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type GinConfig struct {
	Config *viper.Viper
}

func NewGin(config *viper.Viper) *GinConfig {
	return &GinConfig{
		Config: config,
	}
}

func (g *GinConfig) Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
