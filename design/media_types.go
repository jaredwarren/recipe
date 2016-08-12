package design

// mysql root: asdf1234

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Meal = MediaType("application/recipe.meal+json", func() {
	Description("An Meal")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("title", String, "title")
		Attribute("courses", ArrayOf(Course), "all courses")
		Attribute("servings", int, "How much to scale eacy recipe by")

		Required("id", "courses")
	})
	View("default", func() {
		Attribute("title")
		Attribute("courses")
		Attribute("servings")
	})
})

var Course = MediaType("application/recipe.course+json", func() {
	Description("An Meal")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("title", String, "title")
		Attribute("recipes", ArrayOf(Recipe), "all recipes")
		Attribute("servings", int, "How much to scale eacy recipe by, overrides meal")

		Required("id", "recipes")
	})
	View("default", func() {
		Attribute("title")
		Attribute("recipes")
		Attribute("servings")
	})
})

var Recipe = MediaType("application/recipe.recipe+json", func() {
	Description("A recipe")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("id", Integer, "Unique recipe ID")
		Attribute("title", String, "Title of recipe")
		Attribute("description", String, "Long description of recipe")
		Attribute("images", ArrayOf(String), "Title of recipe")
		Attribute("servings", Ingredient, "should be quantity and measure e.g. 4 cups. same as ingredients")
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
		Attribute("created", DateTime, "First created")
		Attribute("updated", DateTime, "Last Updated")

		Required("id", "id", "title")
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
		Attribute("created")
		Attribute("updated")
	})
})

var Ingredient = MediaType("application/recipe.ingredient+json", func() {
	Description("An Ingredient")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("name", String, "what's it called")
		Attribute("quantity", UnitOfMeasure, "How much")                                                                             // ???
		Attribute("state", String, "e.g. chopped, sliced, etc.. might need to be array. this might be a mini recipe or just a word") // ???
		Attribute("complete", Boolean, "If it's been added/included")                                                                // ???

		Required("id", "name")
	})
	View("default", func() {
		Attribute("name")
		Attribute("quantity")
		Attribute("state")
		Attribute("complete")
	})
})

var Cookware = MediaType("application/recipe.cookware+json", func() {
	Description("A Cookware")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("name", String, "what's it called")
		Attribute("description", String, "long description")
		Attribute("parts", ArrayOf(Cookware), "list of parts or attachments")
		Attribute("settings", ArrayOf(String), "settings, e.g. temprature")
		Attribute("setup", ArrayOf(Step), "Steps to setting up")
		Attribute("complete", Boolean, "If it's been checked")

		Required("id", "name")
	})
	View("default", func() {
		Attribute("name")
		Attribute("description")
		Attribute("parts")
		Attribute("settings")
		Attribute("setup")
	})
})

var Step = MediaType("application/recipe.step+json", func() {
	Description("A Step")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("title", String, "")
		Attribute("description", String, "HTML")
		Attribute("in_progress", Boolean, "current state") // ??
		Attribute("complete", Boolean, "is completed")
		Attribute("time", DateTime, "time to complete. e.g. boil for 20 min")

		Required("id", "title")
	})
	View("default", func() {
		Attribute("title")
		Attribute("description")
	})
})

var Category = MediaType("application/recipe.category+json", func() {
	Description("A Category")
	Reference(Tag)
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("name", String, "")
		Attribute("type", String, "")
		Attribute("categories", ArrayOf(Category), "")

		Required("id", "name")
	})
	View("default", func() {
		Attribute("name")
		Attribute("type")
		Attribute("categories")
	})
})

var Tag = MediaType("application/recipe.tag+json", func() {
	Description("A Tag")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("name", String, "")

		Required("id", "name")
	})
	View("default", func() {
		Attribute("name")
	})
})

var Source = MediaType("application/recipe.source+json", func() {
	Description("A Source")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("name", String, "")
		Attribute("url", String, "")

		Required("id", "name")
	})
	View("default", func() {
		Attribute("name")
	})
})

var Note = MediaType("application/recipe.note+json", func() {
	Description("A Note")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("name", String, "")

		Required("id", "name")
	})
	View("default", func() {
		Attribute("name")
	})
})

var UnitOfMeasure = MediaType("application/recipe.unitofmeasure+json", func() {
	Description("A recipe")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("name", String, "name e.g. Cups, Table Spoon")
		Attribute("abbreviation", String, "Abbreviation e.g C, Tb")
		Attribute("type", String, "volume, weight, time, length...", func() {
			Default("weight")
		})
		Attribute("diff", Number, "difference from SI")

		Required("id", "diff", "name")
	})
	View("default", func() {
		Attribute("name")
		Attribute("type")
	})
})

/**
* Shippint List
**/

var ShoppingLists = MediaType("application/recipe.shoppinglists+json", func() {
	Description("A list of shopping lists")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("shopping_lists", ArrayOf(ShoppingList), "The list of lists")
		Required("id", "shopping_lists")
	})
	View("default", func() {
		Attribute("shopping_lists")
	})
})

var ShoppingList = MediaType("application/recipe.shoppinglist+json", func() {
	Description("A list of ingredients, and/or cookware")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("name", String, "a name for the list")
		Attribute("items", ArrayOf(ShoppingItem), "The list of items to buy")
		Attribute("store", String, "Store where to tet items")
		Required("id", "items")
	})
	View("default", func() {
		Attribute("name")
		Attribute("ingredients")
		Attribute("cookware")
		Attribute("store")
	})
})

var ShoppingItem = MediaType("application/recipe.shoppingitem+json", func() {
	Description("Converted from ingredient or cookware")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")

		Attribute("name", String, "a name for the list")
		Attribute("items", ArrayOf(Ingredient), "The list of ingredients")
		Attribute("store", String, "Store where to tet items")
		Required("id", "items")
	})
	View("default", func() {
		Attribute("name")
		Attribute("ingredients")
		Attribute("cookware")
		Attribute("store")
	})
})
