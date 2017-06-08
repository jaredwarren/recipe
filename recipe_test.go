package main

import (
	"testing"

	"github.com/boltdb/bolt"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
	"github.com/jaredwarren/recipe/models"
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
	service := goa.New("recipe")
	db := &testDb{}
	rdb = models.NewRecipeDB(db)
	c2 := NewRecipeController(service, db)

	ctx := &app.CreateRecipeContext{
		Payload: &app.CreateRecipePayload{
			Title: "TEST",
		},
		Context: &app.CreateRecipeContext{},
	}
	c2.Create(ctx)

}
