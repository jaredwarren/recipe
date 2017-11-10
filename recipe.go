package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// RecipeController implements the recipe resource.
type RecipeController struct {
	*goa.Controller
	*sql.DB
}

// Page ...
type Page struct {
	Title   string
	Recipes []*app.RecipeRecipe
	//GameID string // temporary
}

// NewRecipeController creates a recipe controller.
func NewRecipeController(service *goa.Service, db *sql.DB) *RecipeController {
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
	stmt, err := c.DB.Prepare("DELETE FROM recipe WHERE id = ?;")
	if err != nil {
		return ctx.InternalServerError(err)
	}

	res, err := stmt.Exec(ctx.ID)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	fmt.Println(res)

	return ctx.OK(nil)
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
