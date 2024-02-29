package main

import (
	"fmt"
	"github.com/Brandon689/goReactTemplate/handlers/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	// Enable CORS
	//e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.Static("/", "frontend/reactTemplate/dist")

	// Routes
	e.GET("/example", api.GetByID)

	port := 8000
	fmt.Printf("Server is running on :%d...\n", port)
	e.Start(fmt.Sprintf(":%d", port))
}
