//go:generate goagen bootstrap -d github.com/jaredwarren/recipe/design

package main

import (
	"database/sql"
	"log"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jaredwarren/recipe/app"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:bladehq@1234@tcp(192.168.100.106:3306)/recipe")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Create service
	service := goa.New("recipe")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "image" controller
	c := NewImageController(service)
	app.MountImageController(service, c)
	// Mount "recipe" controller
	c2 := NewRecipeController(service, db)
	app.MountRecipeController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
