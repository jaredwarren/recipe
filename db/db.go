package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jaredwarren/recipe/api/app"
)

// RecipeDataAccessLayer abstract recipe access
type RecipeDataAccessLayer interface {
	FetchRecipe(string) (*app.RecipeRecipe, error)
	FetchAllRecipes() (app.RecipeRecipeCollection, error)
	DeleteRecipe(string) error
	InsertRecipe(*app.RecipeRecipe) (string, error)
	UpdateRecipe(*app.RecipeRecipe) error
	Close()
}

// NewRecipeDb ...
func NewRecipeDb(driverName string, dataSourceName string) (*RecipeDAL, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &RecipeDAL{
		DB: db,
	}, nil
}

// RecipeDAL mysql implementation of RecipeDataAccessLayer
type RecipeDAL struct {
	DB *sql.DB
}

// InsertRecipe returns a recipe by id
func (r *RecipeDAL) InsertRecipe(recipe *app.RecipeRecipe) (string, error) {
	stmt, err := r.DB.Prepare("INSERT INTO recipe (title, created_at, updated_at) VALUES (?, NOW(), NOW())")
	if err != nil {
		return "", err
	}

	res, err := stmt.Exec(recipe.Title)
	if err != nil {
		return "", err
	}

	recID, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", recID), nil
}

// UpdateRecipe returns a recipe by id
func (r *RecipeDAL) UpdateRecipe(recipe *app.RecipeRecipe) error {
	_, err := r.DB.Exec("UPDATE recipe SET title = ?, updated_at = NOW() WHERE id = ? LIMIT 1;", recipe.Title, recipe.ID)
	return err
}

// FetchRecipe returns a recipe by id
func (r *RecipeDAL) FetchRecipe(id string) (*app.RecipeRecipe, error) {
	recipe := &app.RecipeRecipe{}
	row := r.DB.QueryRow("SELECT * FROM recipe WHERE id = ?;", id)
	if row == nil {
		return nil, nil
	}
	createdDate := ""
	updateDate := ""
	// TODO: updateDate can be null, either fix in db, or scan for possible null here... somehow
	err := row.Scan(&recipe.ID, &recipe.Title, &createdDate, &updateDate)
	if err != nil {
		return nil, err
	}
	ct, _ := time.Parse("2006-01-02", createdDate)
	recipe.Created = &ct
	ut, _ := time.Parse("2006-01-02", createdDate)
	recipe.Updated = &ut

	return recipe, nil
}

// FetchAllRecipes returns all recipes
func (r *RecipeDAL) FetchAllRecipes() (app.RecipeRecipeCollection, error) {
	rows, err := r.DB.Query("SELECT * FROM recipe;")
	if err != nil {
		return nil, err
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
			// TODO: maybe keep going
			return nil, err
		}
		ct, _ := time.Parse("2006-01-02", createdDate)
		r.Created = &ct
		ut, _ := time.Parse("2006-01-02", createdDate)
		r.Updated = &ut

		recipes = append(recipes, &r)
	}

	return recipes, nil
}

// DeleteRecipe delete a recipe by id
func (r *RecipeDAL) DeleteRecipe(id string) error {
	_, err := r.DB.Exec("DELETE FROM recipe WHERE id = ? LIMIT 1;", id)
	return err
}

// func (t *TestDAL) FindPostsForAuthor(Author) []Post {
//   return t.Posts
// }

// func (t *TestDAL) FindCommentsForPost(Post) []Comment {
//   return t.Comments
// }

// Close ...
func (r *RecipeDAL) Close() {
	r.DB.Close()
}
