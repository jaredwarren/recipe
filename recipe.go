package main

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
	"github.com/goadesign/goa"
	// RecipeController_import: start_implement
	"github.com/jaredwarren/recipe/app"
	"github.com/sergi/go-diff/diffmatchpatch"
	// RecipeController_import: end_implement
	"fmt"
)

// RecipeController implements the recipe resource.
type RecipeController struct {
	*goa.Controller
	// RecipeController_struct: start_implement
	*bolt.DB
	// RecipeController_struct: end_implement
}

// NewRecipeController creates a recipe controller.
func NewRecipeController(service *goa.Service, db *bolt.DB) *RecipeController {
	// NewRecipeController_struct: start_implement
	return &RecipeController{
		Controller: service.NewController("RecipeController"),
		DB:         db,
	}
	// NewRecipeController_struct: end_implement
}

// Create runs the create action.
// curl -H "Content-Type: application/json" -X POST -d '{"title":"xyz"}' http://localhost:8080/recipe/recipe
func (c *RecipeController) Create(ctx *app.CreateRecipeContext) error {
	res := &app.RecipeRecipe{}
	res.Title = ctx.Payload.Title
	// TODO: add other stuff here......

	err := rdb.Add(res)
	if err != nil {
		return ctx.InternalServerError(err)
	}
	return ctx.OK(res)
}

// Delete runs the delete action.
func (c *RecipeController) Delete(ctx *app.DeleteRecipeContext) error {
	// RecipeController_Delete: start_implement

	err := c.DB.Update(func(tx *bolt.Tx) error {
		// Retrieve the Recipe bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("Recipe"))

		// Persist bytes to users bucket.
		return b.Delete([]byte(ctx.ID))
	})
	if err != nil {
		return ctx.InternalServerError(err)
	}

	// RecipeController_Delete: end_implement
	return ctx.OK(nil)
}

// Show runs the show action.
// curl http://localhost:8080/recipe/recipe/1
func (c *RecipeController) Show(ctx *app.ShowRecipeContext) error {
	res, err := rdb.Get(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}
	return ctx.OK(res)
}

// Update runs the update action.
/*
curl --user admin:admin \
 --header "Content-Type:application/json" \
 --header "Accept: application/json" \
 --request PATCH \
 --data '{"title":"new"}' \
 http://localhost:8080/recipe/recipe/2
*/
// curl -H "Content-Type: application/json" -X PATCH -d '{"title":"new2"}' http://localhost:8080/recipe/recipe/2
func (c *RecipeController) Update(ctx *app.UpdateRecipeContext) error {
	// RecipeController_Update: start_implement

	oldRes, err := rdb.Get(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	newRes := &app.RecipeRecipe{
		ID: ctx.ID,
	}
	dmp := diffmatchpatch.New()
	//TODO: change id to string
	//TODO: add version control here

	// TODO: load by id, if not found either create or throw not found error
	newRes.Title = ctx.Payload.Title
	// TODO: add other stuff here......

	diffs := dmp.DiffMain(newRes.Title, oldRes.Title, false)
	TODO: for now just store the diff(s), then figure out how to display/patch later
	fmt.Printf("%+v\n", diffs)


	err = rdb.Update(newRes)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	// RecipeController_Update: end_implement
	return nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
