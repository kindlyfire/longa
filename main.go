package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/kindlyfire/longa/ent"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := loadConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	config.Print()

	client_, err := ent.Open("sqlite3", config.DatabaseUri)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	client = client_
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()
	e.HideBanner = true

	e.GET("/", routeIndex)
	e.GET("/_/api/links", routeGetLinks, checkAuth)
	e.POST("/_/api/links", routeCreateLink, checkAuth)
	e.POST("/_/api/links/:id", routeUpdateLink, checkAuth)
	e.DELETE("/_/api/links/:id", routeDeleteLink, checkAuth)
	e.GET("/:link", routeAccessLink)

	// Static files
	contentHandler := echo.WrapHandler(http.FileServer(http.FS(frontendContent)))
	e.GET("/_/*", contentHandler, middleware.Rewrite(map[string]string{"/_/*": "/frontend/dist/$1"}))

	e.Logger.Fatal(e.Start(config.ListenAddr))
}
