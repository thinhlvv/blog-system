package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/thinhlvv/blog-system/config"
	"github.com/thinhlvv/blog-system/handler/article"
)

func main() {
	cfg := config.New()

	db := config.MustInitDB(cfg)
	fmt.Println(db)

	e := echo.New()

	// Setup router.
	{
		ctrl := article.New()
		e.POST("/articles", ctrl.Create)
		e.GET("/articles/:id", ctrl.GetByID)
		e.GET("/articles", ctrl.GetAll)
	}

	s := &http.Server{
		Addr:         cfg.Server.Port,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}
