package main

import (
	"github.com/bem-filkom/web-bem-backend/internal/config"
)

func main() {
	fiber := config.NewFiber()
	app := config.NewServer(fiber)
	app.RegisterHandlers()
	app.Run()
}
