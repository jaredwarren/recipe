//go:generate goagen bootstrap -d github.com/jaredwarren/recipe/design

package main

import (
	"github.com/boltdb/bolt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jaredwarren/recipe/app"
	"log"
)

func main() {
	// Init Db
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("recipe.db", 0600, nil)
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

	// Mount "recipe" controller
	c := NewRecipeController(service, db)
	app.MountRecipeController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
