//************************************************************************//
// API "recipe": Model Helpers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/jaredwarren/recipe/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListRecipeRecipe returns an array of view: default.
func (m *RecipeDB) ListRecipeRecipe(ctx context.Context) []*app.RecipeRecipe {
	defer goa.MeasureSince([]string{"goa", "db", "recipeRecipe", "listrecipeRecipe"}, time.Now())

	var native []*Recipe
	var objs []*app.RecipeRecipe
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Recipe", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.RecipeToRecipeRecipe())
	}

	return objs
}

// RecipeToRecipeRecipe loads a Recipe and builds the default view of media type RecipeRecipe.
func (m *Recipe) RecipeToRecipeRecipe() *app.RecipeRecipe {
	recipe := &app.RecipeRecipe{}
	recipe.ID = m.ID
	recipe.Title = m.Title

	return recipe
}

// OneRecipeRecipe loads a Recipe and builds the default view of media type RecipeRecipe.
func (m *RecipeDB) OneRecipeRecipe(ctx context.Context, id int) (*app.RecipeRecipe, error) {
	defer goa.MeasureSince([]string{"goa", "db", "recipeRecipe", "onerecipeRecipe"}, time.Now())

	var native Recipe
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Recipe", "error", err.Error())
		return nil, err
	}

	view := *native.RecipeToRecipeRecipe()
	return &view, err
}

// MediaType Retrieval Functions

// ListRecipeRecipeIngredient returns an array of view: ingredient.
func (m *RecipeDB) ListRecipeRecipeIngredient(ctx context.Context) []*app.RecipeRecipeIngredient {
	defer goa.MeasureSince([]string{"goa", "db", "recipeRecipe", "listrecipeRecipeingredient"}, time.Now())

	var native []*Recipe
	var objs []*app.RecipeRecipeIngredient
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Recipe", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.RecipeToRecipeRecipeIngredient())
	}

	return objs
}

// RecipeToRecipeRecipeIngredient loads a Recipe and builds the ingredient view of media type RecipeRecipe.
func (m *Recipe) RecipeToRecipeRecipeIngredient() *app.RecipeRecipeIngredient {
	recipe := &app.RecipeRecipeIngredient{}
	recipe.Title = m.Title

	return recipe
}

// OneRecipeRecipeIngredient loads a Recipe and builds the ingredient view of media type RecipeRecipe.
func (m *RecipeDB) OneRecipeRecipeIngredient(ctx context.Context, id int) (*app.RecipeRecipeIngredient, error) {
	defer goa.MeasureSince([]string{"goa", "db", "recipeRecipe", "onerecipeRecipeingredient"}, time.Now())

	var native Recipe
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Recipe", "error", err.Error())
		return nil, err
	}

	view := *native.RecipeToRecipeRecipeIngredient()
	return &view, err
}
