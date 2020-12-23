package web

import (
	"codebase/app/api/app/internal/cfg"
	"codebase/app/api/app/internal/web/routes"
	"codebase/pkg/defers"
	"codebase/pkg/log"
	"codebase/pkg/web"
	"codebase/pkg/web/middlewares/auth"
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"go.uber.org/zap"
)

//运行 api server
func New() *http.Server {
	engine := web.NewEngine(cfg.Config.AppDebug)
	engine.Use(auth.Check(cfg.Config.HttpServer.Token), cors.Default())
	routes.Routes(engine)
	httpServer := web.NewServer(engine, &cfg.Config.HttpServer.Config)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Error("api server ListenAndServe error, "+err.Error(), zap.Error(err))
		}
	}()

	log.Info("http server run success, listen: " + cfg.Config.HttpServer.Address)

	defers.Register(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown", zap.Error(err))
		}

		log.Info("Server exiting")
		return nil
	})

	return httpServer
}
