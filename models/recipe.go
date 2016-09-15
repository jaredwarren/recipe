//************************************************************************//
// API "recipe": Models
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

// This is the recipe model
type Recipe struct {
	ID        int `gorm:"primary_key"` // This is the ID PK field
	Title     string
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Recipe) TableName() string {
	return "recipes"

}

// RecipeDB is the implementation of the storage interface for
// Recipe.
type RecipeDB struct {
	Db *gorm.DB
}

// NewRecipeDB creates a new storage type.
func NewRecipeDB(db *gorm.DB) *RecipeDB {
	return &RecipeDB{Db: db}
}

// DB returns the underlying database.
func (m *RecipeDB) DB() interface{} {
	return m.Db
}

// RecipeStorage represents the storage interface.
type RecipeStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Recipe, error)
	Get(ctx context.Context, id int) (*Recipe, error)
	Add(ctx context.Context, recipe *Recipe) error
	Update(ctx context.Context, recipe *Recipe) error
	Delete(ctx context.Context, id int) error

	ListRecipeRecipe(ctx context.Context) []*app.RecipeRecipe
	OneRecipeRecipe(ctx context.Context, id int) (*app.RecipeRecipe, error)

	ListRecipeRecipeIngredient(ctx context.Context) []*app.RecipeRecipeIngredient
	OneRecipeRecipeIngredient(ctx context.Context, id int) (*app.RecipeRecipeIngredient, error)

	UpdateFromRecipePayload(ctx context.Context, payload *app.RecipePayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *RecipeDB) TableName() string {
	return "recipes"

}

// CRUD Functions

// Get returns a single Recipe as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *RecipeDB) Get(ctx context.Context, id int) (*Recipe, error) {
	defer goa.MeasureSince([]string{"goa", "db", "recipe", "get"}, time.Now())

	var native Recipe
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Recipe
func (m *RecipeDB) List(ctx context.Context) ([]*Recipe, error) {
	defer goa.MeasureSince([]string{"goa", "db", "recipe", "list"}, time.Now())

	var objs []*Recipe
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *RecipeDB) Add(ctx context.Context, model *Recipe) error {
	defer goa.MeasureSince([]string{"goa", "db", "recipe", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Recipe", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *RecipeDB) Update(ctx context.Context, model *Recipe) error {
	defer goa.MeasureSince([]string{"goa", "db", "recipe", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Recipe", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *RecipeDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "recipe", "delete"}, time.Now())

	var obj Recipe

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Recipe", "error", err.Error())
		return err
	}

	return nil
}

// RecipeFromRecipePayload Converts source RecipePayload to target Recipe model
// only copying the non-nil fields from the source.
func RecipeFromRecipePayload(payload *app.RecipePayload) *Recipe {
	recipe := &Recipe{}
	recipe.Title = payload.Title

	return recipe
}

// UpdateFromRecipePayload applies non-nil changes from RecipePayload to the model and saves it
func (m *RecipeDB) UpdateFromRecipePayload(ctx context.Context, payload *app.RecipePayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "recipe", "updatefromrecipePayload"}, time.Now())

	var obj Recipe
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Recipe", "error", err.Error())
		return err
	}
	obj.Title = payload.Title

	err = m.Db.Save(&obj).Error
	return err
}
