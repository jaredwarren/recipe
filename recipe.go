package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
)

// RecipeController implements the recipe resource.
type RecipeController struct {
	*goa.Controller
}

// NewRecipeController creates a recipe controller.
func NewRecipeController(service *goa.Service) *RecipeController {
	return &RecipeController{Controller: service.NewController("RecipeController")}
}

// Show runs the show action.
func (c *RecipeController) Show(ctx *app.ShowRecipeContext) error {
	// RecipeController_Show: start_implement

	// Put your logic here

	// RecipeController_Show: end_implement
	res := &app.JaredwarrenRecipe{}
	return ctx.OK(res)
}
