//************************************************************************//
// API "recipe": Application Media Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/jaredwarren/recipe/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A recipe (default view)
//
// Identifier: application/recipe.recipe+json
type RecipeRecipe struct {
	// Unique recipe ID
	ID int `form:"id" json:"id" xml:"id"`
	// Title of recipe
	Title string `form:"title" json:"title" xml:"title"`
}

// Validate validates the RecipeRecipe media type instance.
func (mt *RecipeRecipe) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	return
}

// DecodeRecipeRecipe decodes the RecipeRecipe instance encoded in resp body.
func (c *Client) DecodeRecipeRecipe(resp *http.Response) (*RecipeRecipe, error) {
	var decoded RecipeRecipe
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A recipe (default view)
//
// Identifier: application/recipe.unitofmeasure+json
type RecipeUnitofmeasure struct {
	Name string `form:"name" json:"name" xml:"name"`
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the RecipeUnitofmeasure media type instance.
func (mt *RecipeUnitofmeasure) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if !(mt.Type == "volume" || mt.Type == "weight") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.type`, mt.Type, []interface{}{"volume", "weight"}))
	}
	return
}

// DecodeRecipeUnitofmeasure decodes the RecipeUnitofmeasure instance encoded in resp body.
func (c *Client) DecodeRecipeUnitofmeasure(resp *http.Response) (*RecipeUnitofmeasure, error) {
	var decoded RecipeUnitofmeasure
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
