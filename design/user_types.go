package design

// mysql root: asdf1234

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// RecipePayload ...
var RecipePayload = Type("RecipePayload", func() {
	// minimum data required to create or update a recipe
	Attribute("title", String, "Recipe Title") // Attribute value of type integer
	Attribute("description", String, "Long description of recipe")
	Attribute("images", ArrayOf(String), "Images of recipe")
	Attribute("image", String, "Image of recipe")
	Attribute("quantity", UnitOfMeasure, "quantity, measure, servings, yield e.g. 4 cups.")
	Attribute("prep_time", DateTime, "Amount of time to prepare")
	Attribute("cook_time", DateTime, "Amount of time to cook")
	Attribute("wait_time", DateTime, "Amount of time to wait for things such as mairnading")
	//Attribute("cookware", ArrayOf(CookwarePayload), "List of tools needed")
	Attribute("version", String, "Version Number e.g. 1.0.1")
	//Attribute("ingredients", RecipeList, "List of ingredients")
	//Attribute("ingredients", ArrayOf("RecipePayload"), "List of ingredients")
	//Attribute("ingredients", ArrayOf(RecipeList), "List of ingredients")
	//Attribute("ingredients", "RecipePayload", "List of ingredients")
	//Attribute("directions", CollectionOf("application/recipe.step+json"), "List of steps") // ??? might need to be an array of grouped steps i.e. prep_steps cook_steps, plate_steps..
	//Attribute("categories", CollectionOf("application/recipe.category+json"), "List of categories, basically same as tag")
	Attribute("favorite", Boolean, "Is a favorite, basically a tag")
	//Attribute("tags", CollectionOf("application/recipe.tag+json"), "List of tags")
	Attribute("rating", Number, "rating between 0-1", func() {
		Minimum(0)
		Maximum(1)
	})
	Attribute("difficulty", Number, "rating between 0-1", func() {
		Minimum(0)
		Maximum(1)
	})
	Attribute("source", Source, "Source of recipe")
	//Attribute("notes", CollectionOf("application/recipe.note+json"), "List of dated notes")
	Attribute("state", String, "e.g. chopped, sliced, etc.. might need to be array.")
	Attribute("complete", Boolean, "If it's been added/included")

	Required("title")
})

/*
var User = Type("user", func() {
	Description("A user of the API")
	Attribute("name")             // Simple string attribute
	Attribute("address", func() { // Nested definition, defines a struct in Go
		Attribute("number", Integer, "Street number")
		Attribute("street", String, "Street name")
		Attribute("city", String, "City")
		Required("city") // The address must contain at least a city
	})
	Attribute("friends", ArrayOf("user"))
	Attribute("data", HashOf(String, String))
})
*/
// RecipeList ...
//var RecipeList = ArrayOf("RecipePayload")

/*
// RecipeList ...
var RecipeList = Type("RecipeList", func() {
	Attribute("ingredients", ArrayOf(RecipePayload), "List of ingredients")
})*/

// CookwarePayload ...
var CookwarePayload = Type("CookwarePayload", func() {
	Attribute("name", String, "what's it called")
	Attribute("description", String, "long description")
	//Attribute("parts", ArrayOf("CookwarePayload"), "list of parts or attachments")
	Attribute("settings", ArrayOf(String), "settings, e.g. temprature")
	//Attribute("setup", CollectionOf("application/recipe.step+json"), "Steps to setting up")
	Attribute("complete", Boolean, "If it's been checked")
})
