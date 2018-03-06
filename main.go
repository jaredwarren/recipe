//go:generate goagen bootstrap -d github.com/jaredwarren/recipe/design

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jaredwarren/recipe/app"
)

var db *sql.DB

func main() {
	dbName, err := ioutil.ReadFile("/run/secrets/mysql_name")
	if err != nil {
		panic(err)
	}

	dbUser, err := ioutil.ReadFile("/run/secrets/mysql_user")
	if err != nil {
		panic(err)
	}

	dbPassword, err := ioutil.ReadFile("/run/secrets/mysql_password")
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName))
	if err != nil {
		panic(err)
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
	if err := service.ListenAndServe(":80"); err != nil {
		service.LogError("startup", "err", err)
	}
}
