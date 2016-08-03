package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("recipe", func() {
	Title("Recipe API")
	Description("A simple recipe service")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("recipe", func() {
	BasePath("/recipes")
	DefaultMedia(Recipe)

	Action("show", func() {
		Description("Get recipe by id")
		Routing(GET("/:recipeID"))
		Params(func() {
			Param("recipeID", Integer, "Recipe ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

var Recipe = MediaType("application/jaredwarren.recipe+json", func() {
	Description("A recipe")
	Attributes(func() {
		Attribute("id", Integer, "Unique recipe ID")
		Attribute("title", String, "Title of recipe")
		Attribute("images", ArrayOf(String), "Title of recipe")
		Attribute("Servings", ArrayOf(String), "Title of recipe")
		Attribute("images", ArrayOf(String), "Title of recipe")
		Attribute("images", ArrayOf(String), "Title of recipe")
		Attribute("images", ArrayOf(String), "Title of recipe")
		Attribute("images", ArrayOf(String), "Title of recipe")
		Attribute("images", ArrayOf(String), "Title of recipe")
		Attribute("images", ArrayOf(String), "Title of recipe")

		Required("id", "title")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
	})
})

var UnitOfMeasure = MediaType("application/jaredwarren.unitofmeasure+json", func() {
	Description("A recipe")
	Attributes(func() {
		Attribute("name", String, "")
		Attribute("abbreviation", String, "")
		Attribute("type", String, "", func() {
			Enum("volume", "weight")
			Default("weight")
		})
		Attribute("diff", Number, "")

		Required("diff", "name")
	})
	View("default", func() {
		Attribute("name")
		Attribute("title")
	})
})
