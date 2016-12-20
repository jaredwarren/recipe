// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/jaredwarren/recipe/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe
// --version=v1.0.0
//
// API "recipe": Application Resource Href Factories
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"fmt"
	"strings"
)

// ImageHref returns the resource href.
func ImageHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/recipe/images/%v", paramid)
}

// RecipeHref returns the resource href.
func RecipeHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/recipe/recipe/%v", paramid)
}
