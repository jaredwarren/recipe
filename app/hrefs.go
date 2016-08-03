//************************************************************************//
// API "recipe": Application Resource Href Factories
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/jaredwarren/recipe/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// RecipeHref returns the resource href.
func RecipeHref(recipeID interface{}) string {
	return fmt.Sprintf("/recipes/%v", recipeID)
}
