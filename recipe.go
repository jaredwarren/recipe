package main

import (
	"github.com/boltdb/bolt"
	"github.com/goadesign/goa"
	// RecipeController_import: start_implement
	"github.com/jaredwarren/recipe/app"
	"github.com/sergi/go-diff/diffmatchpatch"
	// RecipeController_import: end_implement
	"fmt"
)

// DbInterface ...
type DbInterface interface {
	View(func(tx *bolt.Tx) error) error
	Update(func(tx *bolt.Tx) error) error
}

// RecipeController implements the recipe resource.
type RecipeController struct {
	*goa.Controller
	// RecipeController_struct: start_implement
	DB DbInterface
	// RecipeController_struct: end_implement
}

// NewRecipeController creates a recipe controller.
func NewRecipeController(service *goa.Service, db DbInterface) *RecipeController {
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
// curl -X DELETE http://localhost:8080/recipe/recipe/2
func (c *RecipeController) Delete(ctx *app.DeleteRecipeContext) error {
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
// curl -H "Content-Type: application/json" -X PATCH -d '{"title":"new2"}' http://localhost:8080/recipe/recipe/2
func (c *RecipeController) Update(ctx *app.UpdateRecipeContext) error {
	oldRes, err := rdb.Get(ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	newRes := &app.RecipeRecipe{
		ID:    ctx.ID,
		Title: ctx.Payload.Title,
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(newRes.Title, oldRes.Title, false)
	if len(diffs) > 1 {

	}
	// description
	// cook_time
	// prep_time
	//
	//TODO: for now just store the diff(s), then figure out how to display/patch later
	fmt.Printf("%+v\n", diffs)

	err = rdb.Update(newRes)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	return nil
}
