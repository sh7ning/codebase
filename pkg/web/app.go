package web

import (
	"codebase/pkg/web/middlewares/auth"
	"codebase/pkg/web/middlewares/errors"
	"codebase/pkg/web/middlewares/logger"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Address string `mapstructure:"address" validate:"required"`
	Token   string `mapstructure:"token"`
}

func NewEngine(debug bool, config *Config) *gin.Engine {
	//mode: debug | release | test
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(logger.Logger(), errors.Recovery(), auth.Check(config.Token), cors.Default())
	engine.NoRoute(errors.NoFound())

	return engine
}

func NewServer(engine *gin.Engine, config *Config) *http.Server {
	//handle static file
	//engine.StaticFile("/", "public/index.html")
	//engine.StaticFile("/favicon.ico", "public/favicon.ico")
	//engine.Static("/static", "public/static")

	return &http.Server{
		Addr: config.Address,
		//Handler:        http.TimeoutHandler(engine, 60*time.Second, "request timeout"),
		Handler:      engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		//IdleTimeout:	30 * time.Second,
	}
}
