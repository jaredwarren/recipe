// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "recipe": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/jaredwarren/recipe/web/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe/web
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// ImageHref returns the resource href.
func ImageHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/images/%v", paramid)
}

// RecipeHref returns the resource href.
func RecipeHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/recipe/%v", paramid)
}
