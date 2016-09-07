package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Recipe ...

//this might need fixed
var _ = Resource("recipe", func() {
	BasePath("/recipe")
	DefaultMedia(RecipeMedia)

	// Action("list", func() {
	// 	Description("Retrieve a list or recipes")
	// 	Routing(GET("/"))
	// 	Response(OK, ArrayOf(Recipe))
	// })

	Action("show", func() {
		Description("Display an recipe by id")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Recipe ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("update", func() {
		Routing(PUT("/:id"))
		Payload(RecipePayload)
		Description("")
		Response(OK)
	})

	Action("create", func() {
		Routing(POST("/"))
		Payload(RecipePayload)
		Description("")
		Response(OK)
	})

	Action("delete", func() {
		Description("")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", String, "Recipe ID")
		})
		Response(OK)
	})
})
