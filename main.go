//go:generate goagen bootstrap -d github.com/jaredwarren/recipe/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jaredwarren/recipe/app"
	// main_import start_implement
	"log"

	"github.com/boltdb/bolt"
	// main_import end_implement
)

func main() {
	// main_db start_implement
	// Init Db
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("recipe.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// make sure buckets exsists
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Recipe"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	// main_db end_implement

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
