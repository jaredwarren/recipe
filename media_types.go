package design

// mysql root: asdf1234

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Recipe = MediaType("application/recipe.recipe+json", func() {
	Description("A recipe")
	Attributes(func() {
		Attribute("id", Integer, "Unique recipe ID")
		Attribute("title", String, "Title of recipe")
		Attribute("description", String, "Long description of recipe")
		Attribute("images", ArrayOf(String), "Title of recipe")
		Attribute("servings", ArrayOf(String), "Title of recipe")
		Attribute("prep_time", DateTime, "Amount of time to prepare")
		Attribute("cook_time", DateTime, "Amount of time to cook")
		Attribute("wait_time", DateTime, "Amount of time to wait for things such as mairnading")
		Attribute("cookware", ArrayOf(Cookware), "List of tools needed")
		Attribute("version", String, "Version Number e.g. 1.0.1")
		Attribute("ingredients", ArrayOf(Ingredient), "List of ingredients")
		Attribute("sub_recipes", ArrayOf(Recipe), "List of ingredients")
		Attribute("directions", ArrayOf(Step), "List of steps")
		Attribute("categories", ArrayOf(Category), "List of categories, basically same as tag")
		Attribute("favorite", Boolean, "Is a favorite, basically a tag")
		Attribute("tags", ArrayOf(Tag), "List of tags")
		Attribute("rating", Number, "rating between 0-1", func() {
			Minimum(0)
			Maximum(1)
		})
		Attribute("difficulty", Number, "rating between 0-1", func() {
			Minimum(0)
			Maximum(1)
		})
		Attribute("source", Source, "Source of recipe")
		Attribute("notes", ArrayOf(Note), "List of dated notes")

		Required("id", "title")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("images")
		Attribute("servings")
		Attribute("prep_time")
		Attribute("cook_time")
		Attribute("wait_time")
		Attribute("cookware")
		Attribute("version")
		Attribute("ingredients")
		Attribute("sub_recipes")
		Attribute("directions")
		Attribute("categories")
		Attribute("favorite")
		Attribute("tags")
		Attribute("rating")
		Attribute("difficulty")
		Attribute("source")
		Attribute("notes")
	})
})

var UnitOfMeasure = MediaType("application/recipe.unitofmeasure+json", func() {
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
		Attribute("type")
	})
})
