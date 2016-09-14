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
	// Operand value
	Value *int `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// Validate validates the recipePayload type instance.
func (ut *recipePayload) Validate() (err error) {
	if ut.Value == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "value"))
	}

	return
}

// Publicize creates RecipePayload from recipePayload
func (ut *recipePayload) Publicize() *RecipePayload {
	var pub RecipePayload
	if ut.Value != nil {
		pub.Value = *ut.Value
	}
	return &pub
}

// RecipePayload user type.
type RecipePayload struct {
	// Operand value
	Value int `form:"value" json:"value" xml:"value"`
}
