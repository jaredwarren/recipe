//************************************************************************//
// API "recipe": Application User Types
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

import "github.com/goadesign/goa"

// recipePayload user type.
type recipePayload struct {
	// Recipe Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the recipePayload type instance.
func (ut *recipePayload) Validate() (err error) {
	if ut.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	return
}

// Publicize creates RecipePayload from recipePayload
func (ut *recipePayload) Publicize() *RecipePayload {
	var pub RecipePayload
	if ut.Title != nil {
		pub.Title = *ut.Title
	}
	return &pub
}

// RecipePayload user type.
type RecipePayload struct {
	// Recipe Title
	Title string `form:"title" json:"title" xml:"title"`
}

// Validate validates the RecipePayload type instance.
func (ut *RecipePayload) Validate() (err error) {
	if ut.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	return
}
