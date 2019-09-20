package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/thinhlvv/blog-system/config"
	"github.com/thinhlvv/blog-system/handler/article"
	"github.com/thinhlvv/blog-system/model"
	"github.com/thinhlvv/blog-system/pkg"
)

func main() {
	cfg := config.New()

	db := config.MustInitDB(cfg)
	defer db.Close()
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	app := model.App{
		DB:               db,
		RequestValidator: pkg.NewRequestValidator(),
	}

	e := echo.New()

	// Setup router.
	{
		ctrl := article.New(app)
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
