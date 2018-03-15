package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Recipe ...
var _ = Resource("recipe", func() {
	BasePath("/api")
	DefaultMedia(RecipeMedia)

	Action("list", func() {
		Description("List recipes")
		Scheme("http")
		Routing(GET("/"))
		Response(OK, CollectionOf(RecipeMedia))
		Response(InternalServerError, ErrorMedia)
		// Headers(func() { // Headers describe relevant action headers
		// 	// Header("Authorization", String)
		// 	// Header("X-Account", Integer)
		// 	// Required("Authorization", "X-Account")
		// })
	})

	Action("show", func() {
		Description("Display an recipe by id")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Recipe ID")
		})
		Response(OK)
		Response(Created)
		Response(InternalServerError, ErrorMedia)
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
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
	})
})

// web ...
var _ = Resource("web", func() {
	BasePath("/recipe")
	DefaultMedia(RecipeMedia)

	Action("list", func() {
		Description("List recipes")
		Scheme("http")
		Routing(GET("/"))
		Response(OK, "text/html")
		Response(InternalServerError, ErrorMedia)
	})

	Action("show", func() {
		Description("Display an recipe by id")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Recipe ID")
		})
		Response(OK, "text/html")
		Response(Created, "text/html")
		Response(InternalServerError, ErrorMedia)
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
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
	})
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
