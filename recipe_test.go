package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
	"github.com/jaredwarren/recipe/app/test"
)

var testDB *sql.DB

func setupTestCase(t *testing.T) func(t *testing.T) {
	// setup
	var err error
	testDB, err = sql.Open("mysql", "root:bladehq@1234@tcp(192.168.100.106:3306)/")
	if err != nil {
		log.Fatal(err)
	}

	_, err = testDB.Exec("DROP DATABASE IF EXISTS test_db;")
	if err != nil {
		panic(err)
	}

	_, err = testDB.Exec("CREATE DATABASE test_db;")
	if err != nil {
		panic(err)
	}

	// reconnect to test_db
	testDB, err = sql.Open("mysql", "root:bladehq@1234@tcp(192.168.100.106:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := testDB.Query("SHOW TABLES FROM recipe;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// copy all tables from recipe
	for rows.Next() {
		var (
			tablename string
		)
		if err := rows.Scan(&tablename); err != nil {
			panic(err)
		}

		_, err = testDB.Exec(fmt.Sprintf("CREATE TABLE test_db.%s LIKE recipe.%s;", tablename, tablename))
		if err := rows.Scan(&tablename); err != nil {
			panic(err)
		}
	}

	return func(t *testing.T) {
		defer testDB.Close()
		/*/ optionally cleanup
		_, err = testDB.Exec("DROP DATABASE IF EXISTS test_db;")
		if err != nil {
			panic(err)
		}
		//*/
	}
}

// TestCreate ...
func TestCreate(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var (
		service = goa.New("recipe")
		ctrl    = NewRecipeController(service, testDB)
		ctx     = context.Background()
	)
	payload := &app.CreateRecipePayload{
		Title: randSeq(10),
	}

	r := test.CreateRecipeCreated(t, ctx, service, ctrl, payload)
	loc := r.Header().Get("Location")
	if loc == "" {
		t.Fatalf("missing Location header")
	}
	_, recID := filepath.Split(loc)

	// test
	var rowtitle string
	err := testDB.QueryRow("SELECT title FROM recipe WHERE id=?", recID).Scan(&rowtitle)
	switch {
	case err == sql.ErrNoRows:
		t.Fatal("Failed to insert")
	case err != nil:
		t.Fatal(err)
	case rowtitle != payload.Title:
		t.Fatal("Title didn't match id")
	}
}

// TestDelete ...
func TestDelete(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	// insert
	result, err := testDB.Exec("INSERT INTO recipe (title) VALUES('ASDF');")
	if err != nil {
		t.Fatal("Failed to insert")
	}
	recID, err := result.LastInsertId()
	if err != nil {
		t.Fatal("Failed to insert")
	}

	var (
		service = goa.New("recipe")
		ctrl    = NewRecipeController(service, testDB)
		ctx     = context.Background()
	)

	test.DeleteRecipeNoContent(t, ctx, service, ctrl, fmt.Sprintf("%d", recID))

	// test
	var numRows int
	err = testDB.QueryRow("SELECT COUNT(*) FROM recipe WHERE id=?", recID).Scan(&numRows)
	if err != nil {
		t.Fatal(err)
	}
	if numRows > 0 {
		t.Fatal("didn't delete")
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
