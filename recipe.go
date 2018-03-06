package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
)

// Page ...
type Page struct {
	Title   string
	Recipes []*app.RecipeRecipe
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
	fmt.Println(" - Create Recipe:", ctx.Payload.Title)
	stmt, err := c.DB.Prepare("INSERT INTO recipe (title, created_at, updated_at) VALUES (?, NOW(), NOW())")
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

	ctx.ResponseData.Header().Set("Location", app.RecipeHref(rec.ID))
	return ctx.Created()
}

// Delete runs the delete action.
// curl -X DELETE http://localhost:8080/recipe/recipe/2
func (c *RecipeController) Delete(ctx *app.DeleteRecipeContext) error {
	_, err := c.DB.Exec("DELETE FROM recipe WHERE id = ? LIMIT 1;", ctx.ID)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()
}

// List runs the list action.
func (c *RecipeController) List(ctx *app.ListRecipeContext) error {
	rows, err := c.DB.Query("SELECT * FROM recipe;")
	if err != nil {
		fmt.Println(err)
		return ctx.InternalServerError(err)
	}

	recipes := make(app.RecipeRecipeCollection, 0)

	defer rows.Close()
	for rows.Next() {
		r := app.RecipeRecipe{}
		createdDate := ""
		updateDate := ""
		// TODO: updateDate can be null, either fix in db, or scan for possible null here... somehow
		err := rows.Scan(&r.ID, &r.Title, &createdDate, &updateDate)
		if err != nil {
			fmt.Println(err)
			return ctx.InternalServerError(err)
		}
		ct, _ := time.Parse("2006-01-02", createdDate)
		r.Created = &ct
		ut, _ := time.Parse("2006-01-02", createdDate)
		r.Updated = &ut

		recipes = append(recipes, &r)
	}

	fmt.Printf("~~~%+v~~~", recipes)

	templatePath := "recipe/recipeList.html"
	// TODO: Move to outside or insice MakeMuxer func in production; user here to test, so templates are recompiled every request
	tpl := template.Must(template.New(templatePath).ParseFiles(fmt.Sprintf("templates/%s", templatePath), "templates/base.html"))

	// recipes := []*app.RecipeRecipe{}
	// recipes = append(recipes, &app.RecipeRecipe{
	// 	Title: "ASDR",
	// })
	page := &Page{
		Title:   "Games",
		Recipes: recipes,
	}

	var doc bytes.Buffer
	err = tpl.ExecuteTemplate(&doc, "base", page)
	if err != nil {
		fmt.Println(err)
		ctx.InternalServerError(err)
	}

	// GameController_Start: end_implement
	return ctx.OK(doc.Bytes())
}

// Show runs the show action.
// curl http://localhost:8080/recipe/recipe/1
func (c *RecipeController) Show(ctx *app.ShowRecipeContext) error {
	row := c.DB.QueryRow("SELECT * FROM recipe WHERE id = ?;", ctx.ID)
	if row == nil {
		return ctx.NotFound()
	}

	recipe := app.RecipeRecipe{}
	err := row.Scan(&recipe.ID, &recipe.Title, &recipe.Created, &recipe.Updated)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	fmt.Printf("~~~%+v~~~", recipe)

	templatePath := "recipe/recipe.html"
	// TODO: Move to outside or insice MakeMuxer func in production; user here to test, so templates are recompiled every request
	tpl := template.Must(template.New(templatePath).ParseFiles(fmt.Sprintf("templates/%s", templatePath), "templates/base.html"))

	recipes := []*app.RecipeRecipe{}
	recipes = append(recipes, &app.RecipeRecipe{
		Title: "ASDR",
	})
	page := &Page{
		Title: "Games",
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

	_, err := c.DB.Exec("UPDATE recipe SET title = ?, updated_at = NOW() WHERE id = ? LIMIT 1;", ctx.Payload.Title, ctx.ID)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	// RecipeController_Update: end_implement
	return ctx.NoContent()
}
