//************************************************************************//
// API "recipe": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/jaredwarren/recipe/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// A Category (default view)
//
// Identifier: application/recipe.category+json; view=default
type RecipeCategory struct {
	Categories RecipeCategoryCollection `form:"categories,omitempty" json:"categories,omitempty" xml:"categories,omitempty"`
	Name       string                   `form:"name" json:"name" xml:"name"`
	Type       *string                  `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Validate validates the RecipeCategory media type instance.
func (mt *RecipeCategory) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if err2 := mt.Categories.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeRecipeCategory decodes the RecipeCategory instance encoded in resp body.
func (c *Client) DecodeRecipeCategory(resp *http.Response) (*RecipeCategory, error) {
	var decoded RecipeCategory
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RecipeCategoryCollection is the media type for an array of RecipeCategory (default view)
//
// Identifier: application/recipe.category+json; type=collection; view=default
type RecipeCategoryCollection []*RecipeCategory

// Validate validates the RecipeCategoryCollection media type instance.
func (mt RecipeCategoryCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

		if err2 := e.Categories.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeRecipeCategoryCollection decodes the RecipeCategoryCollection instance encoded in resp body.
func (c *Client) DecodeRecipeCategoryCollection(resp *http.Response) (RecipeCategoryCollection, error) {
	var decoded RecipeCategoryCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A Cookware (default view)
//
// Identifier: application/recipe.cookware+json; view=default
type RecipeCookware struct {
	// long description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// what's it called
	Name string `form:"name" json:"name" xml:"name"`
	// list of parts or attachments
	Parts RecipeCookwareCollection `form:"parts,omitempty" json:"parts,omitempty" xml:"parts,omitempty"`
	// settings, e.g. temprature
	Settings []string `form:"settings,omitempty" json:"settings,omitempty" xml:"settings,omitempty"`
	// Steps to setting up
	Setup RecipeStepCollection `form:"setup,omitempty" json:"setup,omitempty" xml:"setup,omitempty"`
}

// Validate validates the RecipeCookware media type instance.
func (mt *RecipeCookware) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if err2 := mt.Parts.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if err2 := mt.Setup.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeRecipeCookware decodes the RecipeCookware instance encoded in resp body.
func (c *Client) DecodeRecipeCookware(resp *http.Response) (*RecipeCookware, error) {
	var decoded RecipeCookware
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RecipeCookwareCollection is the media type for an array of RecipeCookware (default view)
//
// Identifier: application/recipe.cookware+json; type=collection; view=default
type RecipeCookwareCollection []*RecipeCookware

// Validate validates the RecipeCookwareCollection media type instance.
func (mt RecipeCookwareCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

		if err2 := e.Parts.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if err2 := e.Setup.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeRecipeCookwareCollection decodes the RecipeCookwareCollection instance encoded in resp body.
func (c *Client) DecodeRecipeCookwareCollection(resp *http.Response) (RecipeCookwareCollection, error) {
	var decoded RecipeCookwareCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// An Meal (default view)
//
// Identifier: application/recipe.course+json; view=default
type RecipeCourse struct {
	// all recipes
	Recipes RecipeRecipeCollection `form:"recipes" json:"recipes" xml:"recipes"`
	// How much to scale eacy recipe by, overrides meal
	Servings *int `form:"servings,omitempty" json:"servings,omitempty" xml:"servings,omitempty"`
	// title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the RecipeCourse media type instance.
func (mt *RecipeCourse) Validate() (err error) {
	if mt.Recipes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "recipes"))
	}

	if err2 := mt.Recipes.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeRecipeCourse decodes the RecipeCourse instance encoded in resp body.
func (c *Client) DecodeRecipeCourse(resp *http.Response) (*RecipeCourse, error) {
	var decoded RecipeCourse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RecipeCourseCollection is the media type for an array of RecipeCourse (default view)
//
// Identifier: application/recipe.course+json; type=collection; view=default
type RecipeCourseCollection []*RecipeCourse

// Validate validates the RecipeCourseCollection media type instance.
func (mt RecipeCourseCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Recipes == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "recipes"))
		}

		if err2 := e.Recipes.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeRecipeCourseCollection decodes the RecipeCourseCollection instance encoded in resp body.
func (c *Client) DecodeRecipeCourseCollection(resp *http.Response) (RecipeCourseCollection, error) {
	var decoded RecipeCourseCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// An Meal (default view)
//
// Identifier: application/recipe.meal+json; view=default
type RecipeMeal struct {
	// all courses
	Courses RecipeCourseCollection `form:"courses" json:"courses" xml:"courses"`
	// How much to scale eacy recipe by
	Servings *int `form:"servings,omitempty" json:"servings,omitempty" xml:"servings,omitempty"`
	// title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the RecipeMeal media type instance.
func (mt *RecipeMeal) Validate() (err error) {
	if mt.Courses == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "courses"))
	}

	if err2 := mt.Courses.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeRecipeMeal decodes the RecipeMeal instance encoded in resp body.
func (c *Client) DecodeRecipeMeal(resp *http.Response) (*RecipeMeal, error) {
	var decoded RecipeMeal
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A Note (default view)
//
// Identifier: application/recipe.note+json; view=default
type RecipeNote struct {
	// Unique ID
	ID int `form:"id" json:"id" xml:"id"`
	// html text
	Text string `form:"text" json:"text" xml:"text"`
}

// Validate validates the RecipeNote media type instance.
func (mt *RecipeNote) Validate() (err error) {
	if mt.Text == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "text"))
	}

	return
}

// DecodeRecipeNote decodes the RecipeNote instance encoded in resp body.
func (c *Client) DecodeRecipeNote(resp *http.Response) (*RecipeNote, error) {
	var decoded RecipeNote
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RecipeNoteCollection is the media type for an array of RecipeNote (default view)
//
// Identifier: application/recipe.note+json; type=collection; view=default
type RecipeNoteCollection []*RecipeNote

// Validate validates the RecipeNoteCollection media type instance.
func (mt RecipeNoteCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Text == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "text"))
		}

	}
	return
}

// DecodeRecipeNoteCollection decodes the RecipeNoteCollection instance encoded in resp body.
func (c *Client) DecodeRecipeNoteCollection(resp *http.Response) (RecipeNoteCollection, error) {
	var decoded RecipeNoteCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A recipe (default view)
//
// Identifier: application/recipe.recipe+json; view=default
type RecipeRecipe struct {
	// List of categories, basically same as tag
	Categories RecipeCategoryCollection `form:"categories,omitempty" json:"categories,omitempty" xml:"categories,omitempty"`
	// Amount of time to cook
	CookTime *time.Time `form:"cook_time,omitempty" json:"cook_time,omitempty" xml:"cook_time,omitempty"`
	// List of tools needed
	Cookware RecipeCookwareCollection `form:"cookware,omitempty" json:"cookware,omitempty" xml:"cookware,omitempty"`
	// First created
	Created *time.Time `form:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	// Long description of recipe
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// rating between 0-1
	Difficulty *float64 `form:"difficulty,omitempty" json:"difficulty,omitempty" xml:"difficulty,omitempty"`
	// List of steps
	Directions RecipeStepCollection `form:"directions,omitempty" json:"directions,omitempty" xml:"directions,omitempty"`
	// Is a favorite, basically a tag
	Favorite *bool `form:"favorite,omitempty" json:"favorite,omitempty" xml:"favorite,omitempty"`
	// Unique recipe ID
	ID int `form:"id" json:"id" xml:"id"`
	// Title of recipe
	Images []string `form:"images,omitempty" json:"images,omitempty" xml:"images,omitempty"`
	// List of ingredients
	Ingredients RecipeRecipeCollection `form:"ingredients,omitempty" json:"ingredients,omitempty" xml:"ingredients,omitempty"`
	// List of dated notes
	Notes RecipeNoteCollection `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	// Amount of time to prepare
	PrepTime *time.Time `form:"prep_time,omitempty" json:"prep_time,omitempty" xml:"prep_time,omitempty"`
	// quantity, measure, servings, yield e.g. 4 cups.
	Quantity *RecipeUnitofmeasure `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// rating between 0-1
	Rating *float64 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
	// Source of recipe
	Source *RecipeSource `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// List of tags
	Tags RecipeTagCollection `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Title of recipe
	Title string `form:"title" json:"title" xml:"title"`
	// Last Updated
	Updated *time.Time `form:"updated,omitempty" json:"updated,omitempty" xml:"updated,omitempty"`
	// Version Number e.g. 1.0.1
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Amount of time to wait for things such as mairnading
	WaitTime *time.Time `form:"wait_time,omitempty" json:"wait_time,omitempty" xml:"wait_time,omitempty"`
}

// Validate validates the RecipeRecipe media type instance.
func (mt *RecipeRecipe) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	if err2 := mt.Categories.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if err2 := mt.Cookware.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if mt.Difficulty != nil {
		if *mt.Difficulty < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.difficulty`, *mt.Difficulty, 0, true))
		}
	}
	if mt.Difficulty != nil {
		if *mt.Difficulty > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.difficulty`, *mt.Difficulty, 1, false))
		}
	}
	if err2 := mt.Directions.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if err2 := mt.Ingredients.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if err2 := mt.Notes.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if mt.Quantity != nil {
		if err2 := mt.Quantity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *mt.Rating, 0, true))
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, false))
		}
	}
	if mt.Source != nil {
		if err2 := mt.Source.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if err2 := mt.Tags.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// A recipe (ingredient view)
//
// Identifier: application/recipe.recipe+json; view=ingredient
type RecipeRecipeIngredient struct {
	// If it's been added/included
	Complete *bool `form:"complete,omitempty" json:"complete,omitempty" xml:"complete,omitempty"`
	// quantity, measure, servings, yield e.g. 4 cups.
	Quantity *RecipeUnitofmeasure `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// e.g. chopped, sliced, etc.. might need to be array.
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// Title of recipe
	Title string `form:"title" json:"title" xml:"title"`
}

// Validate validates the RecipeRecipeIngredient media type instance.
func (mt *RecipeRecipeIngredient) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	if mt.Quantity != nil {
		if err2 := mt.Quantity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeRecipeRecipe decodes the RecipeRecipe instance encoded in resp body.
func (c *Client) DecodeRecipeRecipe(resp *http.Response) (*RecipeRecipe, error) {
	var decoded RecipeRecipe
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeRecipeRecipeIngredient decodes the RecipeRecipeIngredient instance encoded in resp body.
func (c *Client) DecodeRecipeRecipeIngredient(resp *http.Response) (*RecipeRecipeIngredient, error) {
	var decoded RecipeRecipeIngredient
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RecipeRecipeCollection is the media type for an array of RecipeRecipe (default view)
//
// Identifier: application/recipe.recipe+json; type=collection; view=default
type RecipeRecipeCollection []*RecipeRecipe

// Validate validates the RecipeRecipeCollection media type instance.
func (mt RecipeRecipeCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Title == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "title"))
		}

		if err2 := e.Categories.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if err2 := e.Cookware.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if e.Difficulty != nil {
			if *e.Difficulty < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].difficulty`, *e.Difficulty, 0, true))
			}
		}
		if e.Difficulty != nil {
			if *e.Difficulty > 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].difficulty`, *e.Difficulty, 1, false))
			}
		}
		if err2 := e.Directions.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if err2 := e.Ingredients.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if err2 := e.Notes.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if e.Quantity != nil {
			if err2 := e.Quantity.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].rating`, *e.Rating, 0, true))
			}
		}
		if e.Rating != nil {
			if *e.Rating > 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, false))
			}
		}
		if e.Source != nil {
			if err2 := e.Source.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err2 := e.Tags.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// RecipeRecipeCollection is the media type for an array of RecipeRecipe (ingredient view)
//
// Identifier: application/recipe.recipe+json; type=collection; view=ingredient
type RecipeRecipeIngredientCollection []*RecipeRecipeIngredient

// Validate validates the RecipeRecipeIngredientCollection media type instance.
func (mt RecipeRecipeIngredientCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Title == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "title"))
		}

		if e.Quantity != nil {
			if err2 := e.Quantity.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeRecipeRecipeCollection decodes the RecipeRecipeCollection instance encoded in resp body.
func (c *Client) DecodeRecipeRecipeCollection(resp *http.Response) (RecipeRecipeCollection, error) {
	var decoded RecipeRecipeCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeRecipeRecipeIngredientCollection decodes the RecipeRecipeIngredientCollection instance encoded in resp body.
func (c *Client) DecodeRecipeRecipeIngredientCollection(resp *http.Response) (RecipeRecipeIngredientCollection, error) {
	var decoded RecipeRecipeIngredientCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Converted from ingredient or cookware (default view)
//
// Identifier: application/recipe.shoppingitem+json; view=default
type RecipeShoppingitem struct {
	// Unique ID
	ID int `form:"id" json:"id" xml:"id"`
	// a name for the list
	Name string `form:"name" json:"name" xml:"name"`
	// Store where to get items
	Store *string `form:"store,omitempty" json:"store,omitempty" xml:"store,omitempty"`
}

// Validate validates the RecipeShoppingitem media type instance.
func (mt *RecipeShoppingitem) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// DecodeRecipeShoppingitem decodes the RecipeShoppingitem instance encoded in resp body.
func (c *Client) DecodeRecipeShoppingitem(resp *http.Response) (*RecipeShoppingitem, error) {
	var decoded RecipeShoppingitem
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RecipeShoppingitemCollection is the media type for an array of RecipeShoppingitem (default view)
//
// Identifier: application/recipe.shoppingitem+json; type=collection; view=default
type RecipeShoppingitemCollection []*RecipeShoppingitem

// Validate validates the RecipeShoppingitemCollection media type instance.
func (mt RecipeShoppingitemCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// DecodeRecipeShoppingitemCollection decodes the RecipeShoppingitemCollection instance encoded in resp body.
func (c *Client) DecodeRecipeShoppingitemCollection(resp *http.Response) (RecipeShoppingitemCollection, error) {
	var decoded RecipeShoppingitemCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A list of ingredients, and/or cookware (default view)
//
// Identifier: application/recipe.shoppinglist+json; view=default
type RecipeShoppinglist struct {
	// The list of items to buy
	Items RecipeShoppingitemCollection `form:"items" json:"items" xml:"items"`
	// a name for the list
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Store where to tet items
	Store *string `form:"store,omitempty" json:"store,omitempty" xml:"store,omitempty"`
}

// Validate validates the RecipeShoppinglist media type instance.
func (mt *RecipeShoppinglist) Validate() (err error) {
	if mt.Items == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "items"))
	}

	if err2 := mt.Items.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeRecipeShoppinglist decodes the RecipeShoppinglist instance encoded in resp body.
func (c *Client) DecodeRecipeShoppinglist(resp *http.Response) (*RecipeShoppinglist, error) {
	var decoded RecipeShoppinglist
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RecipeShoppinglistCollection is the media type for an array of RecipeShoppinglist (default view)
//
// Identifier: application/recipe.shoppinglist+json; type=collection; view=default
type RecipeShoppinglistCollection []*RecipeShoppinglist

// Validate validates the RecipeShoppinglistCollection media type instance.
func (mt RecipeShoppinglistCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Items == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "items"))
		}

		if err2 := e.Items.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeRecipeShoppinglistCollection decodes the RecipeShoppinglistCollection instance encoded in resp body.
func (c *Client) DecodeRecipeShoppinglistCollection(resp *http.Response) (RecipeShoppinglistCollection, error) {
	var decoded RecipeShoppinglistCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A list of shopping lists (default view)
//
// Identifier: application/recipe.shoppinglists+json; view=default
type RecipeShoppinglists struct {
	// The list of lists
	ShoppingLists RecipeShoppinglistCollection `form:"shopping_lists" json:"shopping_lists" xml:"shopping_lists"`
}

// Validate validates the RecipeShoppinglists media type instance.
func (mt *RecipeShoppinglists) Validate() (err error) {
	if mt.ShoppingLists == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "shopping_lists"))
	}

	if err2 := mt.ShoppingLists.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeRecipeShoppinglists decodes the RecipeShoppinglists instance encoded in resp body.
func (c *Client) DecodeRecipeShoppinglists(resp *http.Response) (*RecipeShoppinglists, error) {
	var decoded RecipeShoppinglists
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A Source (default view)
//
// Identifier: application/recipe.source+json; view=default
type RecipeSource struct {
	// Unique ID
	ID   int     `form:"id" json:"id" xml:"id"`
	Name string  `form:"name" json:"name" xml:"name"`
	URL  *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// Validate validates the RecipeSource media type instance.
func (mt *RecipeSource) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// DecodeRecipeSource decodes the RecipeSource instance encoded in resp body.
func (c *Client) DecodeRecipeSource(resp *http.Response) (*RecipeSource, error) {
	var decoded RecipeSource
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A Step (default view)
//
// Identifier: application/recipe.step+json; view=default
type RecipeStep struct {
	// HTML
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Title       string  `form:"title" json:"title" xml:"title"`
}

// Validate validates the RecipeStep media type instance.
func (mt *RecipeStep) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	return
}

// DecodeRecipeStep decodes the RecipeStep instance encoded in resp body.
func (c *Client) DecodeRecipeStep(resp *http.Response) (*RecipeStep, error) {
	var decoded RecipeStep
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RecipeStepCollection is the media type for an array of RecipeStep (default view)
//
// Identifier: application/recipe.step+json; type=collection; view=default
type RecipeStepCollection []*RecipeStep

// Validate validates the RecipeStepCollection media type instance.
func (mt RecipeStepCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Title == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "title"))
		}

	}
	return
}

// DecodeRecipeStepCollection decodes the RecipeStepCollection instance encoded in resp body.
func (c *Client) DecodeRecipeStepCollection(resp *http.Response) (RecipeStepCollection, error) {
	var decoded RecipeStepCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A Tag (default view)
//
// Identifier: application/recipe.tag+json; view=default
type RecipeTag struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the RecipeTag media type instance.
func (mt *RecipeTag) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// DecodeRecipeTag decodes the RecipeTag instance encoded in resp body.
func (c *Client) DecodeRecipeTag(resp *http.Response) (*RecipeTag, error) {
	var decoded RecipeTag
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RecipeTagCollection is the media type for an array of RecipeTag (default view)
//
// Identifier: application/recipe.tag+json; type=collection; view=default
type RecipeTagCollection []*RecipeTag

// Validate validates the RecipeTagCollection media type instance.
func (mt RecipeTagCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// DecodeRecipeTagCollection decodes the RecipeTagCollection instance encoded in resp body.
func (c *Client) DecodeRecipeTagCollection(resp *http.Response) (RecipeTagCollection, error) {
	var decoded RecipeTagCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A recipe (default view)
//
// Identifier: application/recipe.unitofmeasure+json; view=default
type RecipeUnitofmeasure struct {
	// name e.g. Cups, Table Spoon
	Name string `form:"name" json:"name" xml:"name"`
	// volume, weight, time, length...
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the RecipeUnitofmeasure media type instance.
func (mt *RecipeUnitofmeasure) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// DecodeRecipeUnitofmeasure decodes the RecipeUnitofmeasure instance encoded in resp body.
func (c *Client) DecodeRecipeUnitofmeasure(resp *http.Response) (*RecipeUnitofmeasure, error) {
	var decoded RecipeUnitofmeasure
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
