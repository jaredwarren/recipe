package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/goadesign/goa"
	api "github.com/jaredwarren/recipe/api/app"
	rdb "github.com/jaredwarren/recipe/db"
	"github.com/jaredwarren/recipe/web/app"
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
	rec := &api.RecipeRecipe{
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

	templatePath := "recipe/recipeList.html"
	// TODO: Move to outside or insice MakeMuxer func in production; user here to test, so templates are recompiled every request
	tpl := template.Must(template.New(templatePath).ParseFiles(fmt.Sprintf("templates/%s", templatePath), "templates/base.html"))

	page := struct {
		Title   string
		Recipes api.RecipeRecipeCollection
	}{
		Title:   "Recipe:",
		Recipes: recipes,
	}

	var doc bytes.Buffer
	err = tpl.ExecuteTemplate(&doc, "base", page)
	if err != nil {
		fmt.Println(err)
		ctx.InternalServerError(err)
	}

	return ctx.OK(doc.Bytes())
}

// Show runs the show action.
func (c *RecipeController) Show(ctx *app.ShowRecipeContext) error {
	recipe, err := c.DB.FetchRecipe(ctx.ID)
	if err != nil {
		fmt.Println(err)
		return ctx.InternalServerError(err)
	}

	templatePath := "recipe/recipe.html"
	// TODO: Move to outside or insice MakeMuxer func in production; user here to test, so templates are recompiled every request
	tpl := template.Must(template.New(templatePath).ParseFiles(fmt.Sprintf("templates/%s", templatePath), "templates/base.html"))

	page := struct {
		Title  string
		Recipe *api.RecipeRecipe
	}{
		Title:  "Recipe:",
		Recipe: recipe,
	}

	var doc bytes.Buffer
	err = tpl.ExecuteTemplate(&doc, "base", page)
	if err != nil {
		panic(err)
		ctx.InternalServerError(err)
	}

	return ctx.OK(doc.Bytes())
}

// Update runs the update action.
func (c *RecipeController) Update(ctx *app.UpdateRecipeContext) error {
	res := &app.RecipeRecipe{}
	return ctx.OK(res)
}
