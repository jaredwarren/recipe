package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
)

// Page ...
type Page struct {
	Title   string
	Recipes []*app.RecipeRecipe
	//GameID string // temporary
}

// RecipeController implements the recipe resource.
type RecipeController struct {
	*goa.Controller
	*sql.DB
}

// NewRecipeController creates a recipe controller.
func NewRecipeController(service *goa.Service, db *sql.DB) *RecipeController {
	return &RecipeController{
		Controller: service.NewController("RecipeController"),
		DB:         db,
	}
}

// Create runs the create action.
func (c *RecipeController) Create(ctx *app.CreateRecipeContext) error {
	// RecipeController_Create: start_implement

	stmt, err := c.DB.Prepare("INSERT INTO recipe (title) VALUES (?)")
	if err != nil {
		return ctx.InternalServerError(err)
	}

	res, err := stmt.Exec(ctx.Payload.Title)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	recID, err := res.LastInsertId()
	if err != nil {
		return ctx.InternalServerError(err)
	}
	rec := &app.RecipeRecipe{
		ID:    fmt.Sprintf("%d", recID),
		Title: ctx.Payload.Title,
	}

	// RecipeController_Create: end_implement
	ctx.ResponseData.Header().Set("Location", app.RecipeHref(rec.ID))
	return ctx.Created()
	//return ctx.OK(rec)
}

// Delete runs the delete action.
// curl -X DELETE http://localhost:8080/recipe/recipe/2
func (c *RecipeController) Delete(ctx *app.DeleteRecipeContext) error {
	_, err := c.DB.Exec("DELETE FROM recipe WHERE id = ?;", ctx.ID)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()
}

// List runs the list action.
func (c *RecipeController) List(ctx *app.ListRecipeContext) error {
	// RecipeController_List: start_implement

	// Put your logic here

	// RecipeController_List: end_implement
	return nil
}

// Show runs the show action.
// curl http://localhost:8080/recipe/recipe/1
func (c *RecipeController) Show(ctx *app.ShowRecipeContext) error {
	stmt, err := c.DB.Prepare("SELECT * FROM recipe WHERE id = ?;")
	if err != nil {
		return ctx.InternalServerError(err)
	}

	res, err := stmt.Exec(ctx.ID)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	// if empty
	//return ctx.NotFound()

	fmt.Println(res)

	templatePath := "game/game.html"
	// TODO: Move to outside or insice MakeMuxer func in production; user here to test, so templates are recompiled every request
	tpl := template.Must(template.New(templatePath).ParseFiles(fmt.Sprintf("templates/%s", templatePath), "templates/base.html"))

	recipes := []*app.RecipeRecipe{}
	recipes = append(recipes, &app.RecipeRecipe{
		Title: "ASDR",
	})
	page := &Page{
		Title: "Games",
		//Games:  games,
		//GameID: ctx.GameID,
	}

	var doc bytes.Buffer
	err = tpl.ExecuteTemplate(&doc, "base", page)
	if err != nil {
		ctx.InternalServerError(err)
	}

	// GameController_Start: end_implement
	return ctx.OK(doc.Bytes())
}

// Update runs the update action.
func (c *RecipeController) Update(ctx *app.UpdateRecipeContext) error {
	// RecipeController_Update: start_implement

	// Put your logic here

	// RecipeController_Update: end_implement
	res := &app.RecipeRecipe{}
	return ctx.OK(res)
}
