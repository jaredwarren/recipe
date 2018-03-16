//go:generate goagen bootstrap -d github.com/jaredwarren/recipe/web/design

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jaredwarren/recipe/db"
	"github.com/jaredwarren/recipe/web/app"
)

func main() {
	dbNameFile, ok := os.LookupEnv("DB_NAME_FILE")
	var dbName string
	if ok {
		bDbName, err := ioutil.ReadFile(dbNameFile)
		if err != nil {
			panic(err)
		}
		dbName = string(bDbName)
	} else {
		dbName, _ = os.LookupEnv("DB_NAME")
	}

	dbUserFile, ok := os.LookupEnv("DB_USER_FILE")
	var dbUser string
	if ok {
		bDbUser, err := ioutil.ReadFile(dbUserFile)
		if err != nil {
			panic(err)
		}
		dbUser = string(bDbUser)
	} else {
		dbUser, _ = os.LookupEnv("DB_USER")
	}

	dbPassFile, ok := os.LookupEnv("DB_PASS_FILE")
	var dbPassword string
	if ok {
		bDbPassword, err := ioutil.ReadFile(dbPassFile)
		if err != nil {
			panic(err)
		}
		dbPassword = string(bDbPassword)
	} else {
		dbPassword, _ = os.LookupEnv("DB_PASS")
	}

	dbHostFile, ok := os.LookupEnv("DB_HOST_FILE")
	var dbHost string
	if ok {
		bDbHost, err := ioutil.ReadFile(dbHostFile)
		if err != nil {
			panic(err)
		}
		dbHost = string(bDbHost)
	} else {
		dbHost, _ = os.LookupEnv("DB_HOST")
	}

	rdb, err := db.NewRecipeDb("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName))
	if err != nil {
		panic(err)
	}
	defer rdb.Close()

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
	// // Mount "recipe" controller
	// c2 := NewRecipeController(service)
	// app.MountRecipeController(service, c2)
	// Mount "web" controller
	c2 := NewWebController(service, rdb)
	app.MountWebController(service, c2)

	// Start service
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "80"
	}
	if err := service.ListenAndServe(fmt.Sprintf(":%s", port)); err != nil {
		service.LogError("startup", "err", err)
	}

}
