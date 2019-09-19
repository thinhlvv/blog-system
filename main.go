package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/thinhlvv/blog-system/config"
)

func main() {
	cfg := config.New()

	e := echo.New()

	s := &http.Server{
		Addr:         cfg.Server.Port,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}
