//go:generate goagen bootstrap -d github.com/jaredwarren/recipe/design

package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jaredwarren/recipe/app"
)

// goagen bootstrap -d github.com/jaredwarren/recipe/design

// go build -o cellar && ./cellar
//curl -H "Content-Type: application/json" -X POST -d '{"name":"xyz"}' http://localhost:8080/admin/category

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
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
