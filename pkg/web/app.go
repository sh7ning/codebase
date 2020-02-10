package web

import (
	"app/pkg/cfg"
	"app/pkg/web/middlewares/auth"
	"app/pkg/web/middlewares/errors"
	"app/pkg/web/middlewares/logger"
	"app/pkg/web/routes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func New() *http.Server {
	engine := getEngine()

	return &http.Server{
		Addr: cfg.AppConfig.HttpServer.Address,
		//Handler:        http.TimeoutHandler(engine, 60*time.Second, "request timeout"),
		Handler:      engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		//IdleTimeout:	30 * time.Second,
	}
}

func getEngine() *gin.Engine {
	//mode: debug | release | test
	if cfg.AppConfig.AppDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(logger.Logger(), errors.Recovery(), auth.Check(), cors.AllowAll())
	engine.NoRoute(errors.NoFound())

	//handle static file
	engine.Static("/static", "public/static")
	engine.StaticFile("/favicon.ico", "public/favicon.ico")
	engine.StaticFile("/", "public/index.html")

	routes.Routes(engine)

	return engine
}
