package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/jaredwarren/recipe/client"
)

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

type testDb struct{}

func (d *testDb) View(func(tx *bolt.Tx) error) error {
	return nil
}
func (d *testDb) Update(func(tx *bolt.Tx) error) error {
	return nil
}
func TestAverage(t *testing.T) {
	path := "/recipe/recipe"
	httpClient := http.DefaultClient
	c := client.New(goaclient.HTTPClientDoer(httpClient))

	var payload client.CreateRecipePayload
	err := json.Unmarshal([]byte(`{"title":"xyz"}`), &payload)
	if err != nil {
		t.Errorf("failed to deserialize payload: %s", err)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateRecipe(ctx, path, &payload, "application/json")
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		t.Errorf("failed to deserialize payload: %s", err)
	}

	goaclient.HandleResponse(c.Client, resp, true)

	//
	//
	//

	// service := goa.New("recipe")
	// db := &testDb{}
	// rdb = models.NewRecipeDB(db)
	// c2 := NewRecipeController(service, db)

	// type favContextKey string
	// //k := favContextKey("2")
	// //c := context.WithValue(context.Background(), k, "Go")
	// c := context.Background()

	// var jsonStr = []byte(`{"title":"xyz"}`)
	// req, _ := http.NewRequest("POST", "http://localhost:8080/recipe/recipe", bytes.NewBuffer(jsonStr))
	// req.Header.Set("Content-Type", "application/json")
	// ctx, _ := app.NewCreateRecipeContext(c, req, service)

	// c2.Create(ctx)

}
