package bootstrap

import (
	"log"
	"waizly/internal/config"
)

type Bootstrap struct {
	App *config.App
}

func NewBootstrap(app *config.App) *Bootstrap {
	return &Bootstrap{App: app}
}

func (b *Bootstrap) StartServer() {
	err := b.App.Serve()

	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
