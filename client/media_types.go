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
// Identifier: application/jaredwarren.recipe+json
type JaredwarrenRecipe struct {
	// Unique recipe ID
	ID int `form:"id" json:"id" xml:"id"`
	// Title of recipe
	Title string `form:"title" json:"title" xml:"title"`
}

// Validate validates the JaredwarrenRecipe media type instance.
func (mt *JaredwarrenRecipe) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	return
}

// DecodeJaredwarrenRecipe decodes the JaredwarrenRecipe instance encoded in resp body.
func (c *Client) DecodeJaredwarrenRecipe(resp *http.Response) (*JaredwarrenRecipe, error) {
	var decoded JaredwarrenRecipe
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
