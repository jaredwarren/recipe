package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var sg = StorageGroup("RecipeStorageGroup", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("This is the mysql relational store")
		Model("Recipe", func() {
			BuildsFrom(func() {
				Payload("recipe", "create") // e.g. "bottle", "create" resource definition
			})

			RendersTo(RecipeMedia) // a Media Type definition
			Description("This is the recipe model")

			Field("id", gorma.Integer, func() { // Required for CRUD getters to take a PK argument!
				PrimaryKey()
				Description("This is the ID PK field")
			})

			Field("title", gorma.String)

			Field("CreatedAt", gorma.Timestamp)
			Field("UpdatedAt", gorma.Timestamp)         // Shown for demonstration
			Field("DeletedAt", gorma.NullableTimestamp) // These are added by default
		})
	})
})
