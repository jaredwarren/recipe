package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the cellar application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("recipe", func() {
	Title("Recipe Web interrace")
	Description("CRUD API recipe")
	Contact(func() {
		Name("Jared Warren")
		Email("jlwarren1@gmail.com")
		URL("http://jlwarren1.com")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/goadesign/goa/blob/master/LICENSE")
	})
	Docs(func() {
		Description("TODO")
		URL("http://jlwarren1.com/recipe/docs")
	})
	Host("recipe.localhost")
	Scheme("http")
	Consumes("application/x-www-form-urlencoded", func() {
		Package("github.com/goadesign/goa/encoding/form")
	})
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		Headers("Accept", "Content-Type")
		Expose("Content-Type", "Origin")
		MaxAge(600)
		Credentials()
	})
	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
