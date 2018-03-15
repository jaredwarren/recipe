package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
	rdb "github.com/jaredwarren/recipe/db"
)

// WebController implements the web resource.
type WebController struct {
	*goa.Controller
	DB rdb.RecipeDataAccessLayer
}

// NewWebController creates a web controller.
func NewWebController(service *goa.Service, db rdb.RecipeDataAccessLayer) *WebController {
	return &WebController{
		Controller: service.NewController("WebController"),
		DB:         db,
	}
}

// Create runs the create action.
func (c *WebController) Create(ctx *app.CreateWebContext) error {
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
func (c *WebController) Delete(ctx *app.DeleteWebContext) error {
	err := c.DB.DeleteRecipe(ctx.ID)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()
}

// List runs the list action.
func (c *WebController) List(ctx *app.ListWebContext) error {
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
		Recipes app.RecipeRecipeCollection
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
func (c *WebController) Show(ctx *app.ShowWebContext) error {
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
		Recipe *app.RecipeRecipe
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
func (c *WebController) Update(ctx *app.UpdateWebContext) error {
	// WebController_Update: start_implement

	// Put your logic here

	res := &app.RecipeRecipe{}
	return ctx.OK(res)
	// WebController_Update: end_implement
}
