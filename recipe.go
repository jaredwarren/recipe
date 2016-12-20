package main

import (
	"encoding/binary"
	"encoding/json"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/goadesign/goa"
	// RecipeController_import: start_implement
	"github.com/jaredwarren/recipe/app"
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
func (c *RecipeController) Create(ctx *app.CreateRecipeContext) error {
	// RecipeController_Create: start_implement
	res := &app.RecipeRecipe{}
	fmt.Printf("---CREATE:::%+v", ctx.Payload)
	res.Title = ctx.Payload.Title
	// TODO: add other stuff here......

	// Put your logic here
	err := c.DB.Update(func(tx *bolt.Tx) error {
		// Retrieve the Recipe bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("Recipe"))

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()
		res.ID = int(id)

		// Marshal user data into bytes.
		buf, err := json.Marshal(res)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put([]byte(strconv.Itoa(res.ID)), buf)
	})
	if err != nil {
		return ctx.InternalServerError(err)
	}

	// RecipeController_Create: end_implement

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
func (c *RecipeController) Show(ctx *app.ShowRecipeContext) error {
	// RecipeController_Show: start_implement
	res := &app.RecipeRecipe{}
	c.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Recipe"))
		v := b.Get([]byte(ctx.ID))
		json.Unmarshal(v, res)

		return nil
	})

	// RecipeController_Show: end_implement

	return ctx.OK(res)
}

// Update runs the update action.
func (c *RecipeController) Update(ctx *app.UpdateRecipeContext) error {
	// RecipeController_Update: start_implement

	res := &app.RecipeRecipe{}

	// Put your logic here
	err := c.DB.Update(func(tx *bolt.Tx) error {
		// Retrieve the Recipe bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("Recipe"))

		v := b.Get([]byte(ctx.ID))
		json.Unmarshal(v, res)

		res.Title = ctx.Payload.Title
		// TODO: add other properties here...

		// Marshal user data into bytes.
		buf, err := json.Marshal(res)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put([]byte(strconv.Itoa(res.ID)), buf)
	})
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
