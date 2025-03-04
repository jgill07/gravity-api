package main

import (
	"github.com/jgill07/gravity-api/internal/api"
	"github.com/jgill07/gravity-api/internal/api/server"
	"github.com/jgill07/gravity-api/internal/config"
	"github.com/jgill07/gravity-api/internal/service"
	"github.com/jgill07/gravity-api/internal/store"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	var (
		tstore = store.NewMemoryStore()
		svc    = service.NewService(cfg, tstore)
	)
	server.Run(api.SetupRouter(svc), cfg.ApiConfig.Port)
}
