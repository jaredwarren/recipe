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

	//Parent("course")

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
		Description("")
		Routing(PATCH("/:id"))
		Params(func() {
			Param("id", String, "Recipe ID")
		})

		Payload("RecipePayload")

		Response(NoContent)
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("create", func() {
		Routing(POST("/"))
		Payload("RecipePayload")
		Description("")

		Response(OK)
		Response(InternalServerError, ErrorMedia)
		Response(Created)
	})

	Action("delete", func() {
		Description("")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", String, "Recipe ID")
		})
		Response(OK)
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
	})

	/*Action("add_image", func() {
		Description("add an image to recipe")
		Routing(POST("/:id/add_image"))
		Params(func() {
			Param("id", String, "Recipe ID")
		})

		Payload(ImagePayload) // TODO:

		Response(OK)
		Response(InternalServerError, ErrorMedia)
	})*/

})

var _ = Resource("image", func() {
	BasePath("/images")

	Action("upload", func() {
		Routing(POST("/"))
		Description("Upload an image")
		Response(OK, ImageMedia)
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Description("Show an image metadata")
		Params(func() {
			Param("id", String, "Image ID")
		})
		Response(OK, ImageMedia)
		Response(NotFound)
	})

	Files("/download/*filename", "images/") // Serve files from the "images" directory
})
