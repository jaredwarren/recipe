package design

// mysql root: asdf1234

import (
	//. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var RecipePayload = Type("RecipePayload", func() {
	// minimum data required to create or update a recipe
})
