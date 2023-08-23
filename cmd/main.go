package main

import (
	"github.com/vgekko/anistream-gateway/config"
	"github.com/vgekko/anistream-gateway/internal/app"
)

func main() {
	cfg := config.Load()

	app.RunHttpProxy(cfg)
}
