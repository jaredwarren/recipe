package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/api/app"
	rdb "github.com/jaredwarren/recipe/db"
)

// RecipeController implements the recipe resource.
type RecipeController struct {
	*goa.Controller
	DB rdb.RecipeDataAccessLayer
}

// NewRecipeController creates a recipe controller.
func NewRecipeController(service *goa.Service, db rdb.RecipeDataAccessLayer) *RecipeController {
	return &RecipeController{
		Controller: service.NewController("RecipeController"),
		DB:         db,
	}
}

// Create runs the create action.
func (c *RecipeController) Create(ctx *app.CreateRecipeContext) error {
	rec := &app.RecipeRecipe{
		Title: ctx.Payload.Title,
	}

	id, err := c.DB.InsertRecipe(rec)
	if err != nil {
		return ctx.InternalServerError(err)
	}
	rec.ID = id

	ctx.ResponseData.Header().Set("Location", app.RecipeHref(rec.ID))
	return ctx.Created()
}

// Delete runs the delete action.
// curl -X DELETE http://localhost:8080/recipe/recipe/2
func (c *RecipeController) Delete(ctx *app.DeleteRecipeContext) error {
	err := c.DB.DeleteRecipe(ctx.ID)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()
}

// List runs the list action.
func (c *RecipeController) List(ctx *app.ListRecipeContext) error {
	recipes, err := c.DB.FetchAllRecipes()
	if err != nil {
		fmt.Println(err)
		ctx.InternalServerError(err)
	}
	return ctx.OK(recipes)
}

// Show runs the show action.
// curl http://localhost:8080/recipe/recipe/1
func (c *RecipeController) Show(ctx *app.ShowRecipeContext) error {
	recipe, err := c.DB.FetchRecipe(ctx.ID)
	if err != nil {
		fmt.Println(err)
		return ctx.InternalServerError(err)
	}

	return ctx.OK(recipe)
}

// Update runs the update action.
func (c *RecipeController) Update(ctx *app.UpdateRecipeContext) error {
	rec := &app.RecipeRecipe{
		ID:    ctx.ID,
		Title: ctx.Payload.Title,
	}
	err := c.DB.UpdateRecipe(rec)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()
}
