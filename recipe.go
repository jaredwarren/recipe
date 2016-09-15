package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
	"github.com/jaredwarren/recipe/models"
)

// ErrDatabaseError is the error returned when a db query fails.
var ErrDatabaseError = goa.NewErrorClass("db_error", 500)

var adb *models.RecipeDB

// RecipeController implements the recipe resource.
type RecipeController struct {
	*goa.Controller
}

// NewRecipeController creates a recipe controller.
func NewRecipeController(service *goa.Service) *RecipeController {
	return &RecipeController{Controller: service.NewController("RecipeController")}
}

// Create runs the create action.
func (c *RecipeController) Create(ctx *app.CreateRecipeContext) error {
	// RecipeController_Create: start_implement

	//a := models.Recipe{}
	a := models.Recipe{}
	a.Title = ctx.Payload.Title
	err := adb.Add(ctx.Context, &a)

	//a := models.Account{}
	//a.Name = ctx.Payload.Name
	//err := adb.Add(ctx.Context, &a)
	if err != nil {
		return ErrDatabaseError(err)
	}
	ctx.ResponseData.Header().Set("Location", app.RecipeHref(a.ID))
	return ctx.Created()

	// RecipeController_Create: end_implement
	//res := &app.RecipeRecipe{}
	//return ctx.OK(res)
}

// Delete runs the delete action.
func (c *RecipeController) Delete(ctx *app.DeleteRecipeContext) error {
	// RecipeController_Delete: start_implement

	// Put your logic here

	// RecipeController_Delete: end_implement
	res := &app.RecipeRecipe{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *RecipeController) Show(ctx *app.ShowRecipeContext) error {
	// RecipeController_Show: start_implement

	// Put your logic here

	// RecipeController_Show: end_implement
	res := &app.RecipeRecipe{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *RecipeController) Update(ctx *app.UpdateRecipeContext) error {
	// RecipeController_Update: start_implement

	// Put your logic here

	// RecipeController_Update: end_implement
	res := &app.RecipeRecipe{}
	return ctx.OK(res)
}
