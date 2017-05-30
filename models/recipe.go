package models

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
)

// RecipeDB is the implementation of the storage interface for
// Recipe.
type RecipeDB struct {
	Db *bolt.DB
}

// NewRecipeDB creates a new storage type.
func NewRecipeDB(db *bolt.DB) *RecipeDB {
	return &RecipeDB{Db: db}
}

// DB returns the underlying database.
func (m *RecipeDB) DB() interface{} {
	return m.Db
}

// RecipeStorage represents the storage interface.
type RecipeStorage interface {
	DB() interface{}
	List() ([]*app.RecipeRecipe, error)
	Get(id int) (*app.RecipeRecipe, error)
	Add(recipe *app.RecipeRecipe) error
	Update(recipe *app.RecipeRecipe) error
	Delete(id int) error

	ListRecipeRecipe() []*app.RecipeRecipe
	OneRecipeRecipe(id int) (*app.RecipeRecipe, error)

	ListRecipeRecipeIngredient() []*app.RecipeRecipeIngredient
	OneRecipeRecipeIngredient(id int) (*app.RecipeRecipeIngredient, error)

	UpdateFromRecipePayload(payload *app.RecipePayload, id int) error
}

// TableName overrides the table name settings in Bolt to force a specific table name
// in the database.
func (m *RecipeDB) TableName() string {
	return "recipes"

}

// CRUD Functions

// Get returns a single Recipe as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *RecipeDB) Get(id string) (*app.RecipeRecipe, error) {
	res := &app.RecipeRecipe{}
	m.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Recipe"))
		v := b.Get([]byte(id))
		json.Unmarshal(v, res)
		return nil
	})
	return res, nil
}

// List returns an array of Recipe
func (m *RecipeDB) List() ([]*app.RecipeRecipe, error) {
	defer goa.MeasureSince([]string{"goa", "db", "recipe", "list"}, time.Now())

	var objs []*app.RecipeRecipe
	// err := m.Db.Table(m.TableName()).Find(&objs).Error
	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return nil, err
	// }

	return objs, nil
}

// Add creates a new record.
func (m *RecipeDB) Add(model *app.RecipeRecipe) error {
	err := m.Db.Update(func(tx *bolt.Tx) error {
		// Retrieve the Recipe bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("Recipe"))

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()
		model.ID = strconv.Itoa(int(id))

		// Marshal user data into bytes.
		buf, err := json.Marshal(model)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put([]byte(model.ID), buf)
	})

	return err
}

// Update modifies a single record.
func (m *RecipeDB) Update(model *app.RecipeRecipe) error {
	// Put your logic here
	err := m.Db.Update(func(tx *bolt.Tx) error {
		// Retrieve the Recipe bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("Recipe"))

		//v := b.Get([]byte(id))
		//json.Unmarshal(v, model)

		//res.Title = ctx.Payload.Title
		// TODO: add other properties here...

		// Marshal user data into bytes.
		buf, err := json.Marshal(model)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put([]byte(model.ID), buf)
	})
	return err
}

// Delete removes a single record.
func (m *RecipeDB) Delete(id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "recipe", "delete"}, time.Now())

	// var obj Recipe

	// err := m.Db.Delete(&obj, id).Error

	// if err != nil {
	// 	goa.LogError(ctx, "error deleting Recipe", "error", err.Error())
	// 	return err
	// }

	return nil
}

// RecipeFromRecipePayload Converts source RecipePayload to target Recipe model
// only copying the non-nil fields from the source.
func RecipeFromRecipePayload(payload *app.RecipePayload) *app.RecipeRecipe {
	recipe := &app.RecipeRecipe{}
	recipe.Title = payload.Title

	return recipe
}

// UpdateFromRecipePayload applies non-nil changes from RecipePayload to the model and saves it
func (m *RecipeDB) UpdateFromRecipePayload(payload *app.RecipePayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "recipe", "updatefromrecipePayload"}, time.Now())

	// var obj Recipe
	// err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	// if err != nil {
	// 	goa.LogError(ctx, "error retrieving Recipe", "error", err.Error())
	// 	return err
	// }
	// obj.Title = payload.Title

	// err = m.Db.Save(&obj).Error
	return nil
}
