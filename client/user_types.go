// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "recipe": Application User Types
//
// Command:
// $ goagen
// --design=github.com/jaredwarren/recipe/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"time"
)

// cookwarePayload user type.
type cookwarePayload struct {
	// If it's been checked
	Complete *bool `form:"complete,omitempty" json:"complete,omitempty" xml:"complete,omitempty"`
	// long description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// what's it called
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// settings, e.g. temprature
	Settings []string `form:"settings,omitempty" json:"settings,omitempty" xml:"settings,omitempty"`
}

// Publicize creates CookwarePayload from cookwarePayload
func (ut *cookwarePayload) Publicize() *CookwarePayload {
	var pub CookwarePayload
	if ut.Complete != nil {
		pub.Complete = ut.Complete
	}
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Settings != nil {
		pub.Settings = ut.Settings
	}
	return &pub
}

// CookwarePayload user type.
type CookwarePayload struct {
	// If it's been checked
	Complete *bool `form:"complete,omitempty" json:"complete,omitempty" xml:"complete,omitempty"`
	// long description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// what's it called
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// settings, e.g. temprature
	Settings []string `form:"settings,omitempty" json:"settings,omitempty" xml:"settings,omitempty"`
}

// recipePayload user type.
type recipePayload struct {
	// If it's been added/included
	Complete *bool `form:"complete,omitempty" json:"complete,omitempty" xml:"complete,omitempty"`
	// Amount of time to cook
	CookTime *time.Time `form:"cook_time,omitempty" json:"cook_time,omitempty" xml:"cook_time,omitempty"`
	// Long description of recipe
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// rating between 0-1
	Difficulty *float64 `form:"difficulty,omitempty" json:"difficulty,omitempty" xml:"difficulty,omitempty"`
	// Is a favorite, basically a tag
	Favorite *bool `form:"favorite,omitempty" json:"favorite,omitempty" xml:"favorite,omitempty"`
	// Image of recipe
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// Images of recipe
	Images []string `form:"images,omitempty" json:"images,omitempty" xml:"images,omitempty"`
	// List of ingredients
	Ingredients []*recipePayload `form:"ingredients,omitempty" json:"ingredients,omitempty" xml:"ingredients,omitempty"`
	// Amount of time to prepare
	PrepTime *time.Time `form:"prep_time,omitempty" json:"prep_time,omitempty" xml:"prep_time,omitempty"`
	// rating between 0-1
	Rating *float64 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
	// Source of recipe
	Source *string `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// e.g. chopped, sliced, etc.. might need to be array.
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// Recipe Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Version Number e.g. 1.0.1
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Amount of time to wait for things such as mairnading
	WaitTime *time.Time `form:"wait_time,omitempty" json:"wait_time,omitempty" xml:"wait_time,omitempty"`
}

// Validate validates the recipePayload type instance.
func (ut *recipePayload) Validate() (err error) {
	if ut.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "title"))
	}
	if ut.Difficulty != nil {
		if *ut.Difficulty < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.difficulty`, *ut.Difficulty, 0.000000, true))
		}
	}
	if ut.Difficulty != nil {
		if *ut.Difficulty > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.difficulty`, *ut.Difficulty, 1.000000, false))
		}
	}
	for _, e := range ut.Ingredients {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.rating`, *ut.Rating, 0.000000, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.rating`, *ut.Rating, 1.000000, false))
		}
	}
	return
}

// Publicize creates RecipePayload from recipePayload
func (ut *recipePayload) Publicize() *RecipePayload {
	var pub RecipePayload
	if ut.Complete != nil {
		pub.Complete = ut.Complete
	}
	if ut.CookTime != nil {
		pub.CookTime = ut.CookTime
	}
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.Difficulty != nil {
		pub.Difficulty = ut.Difficulty
	}
	if ut.Favorite != nil {
		pub.Favorite = ut.Favorite
	}
	if ut.Image != nil {
		pub.Image = ut.Image
	}
	if ut.Images != nil {
		pub.Images = ut.Images
	}
	if ut.Ingredients != nil {
		pub.Ingredients = make([]*RecipePayload, len(ut.Ingredients))
		for i2, elem2 := range ut.Ingredients {
			pub.Ingredients[i2] = elem2.Publicize()
		}
	}
	if ut.PrepTime != nil {
		pub.PrepTime = ut.PrepTime
	}
	if ut.Rating != nil {
		pub.Rating = ut.Rating
	}
	if ut.Source != nil {
		pub.Source = ut.Source
	}
	if ut.State != nil {
		pub.State = ut.State
	}
	if ut.Title != nil {
		pub.Title = *ut.Title
	}
	if ut.Version != nil {
		pub.Version = ut.Version
	}
	if ut.WaitTime != nil {
		pub.WaitTime = ut.WaitTime
	}
	return &pub
}

// RecipePayload user type.
type RecipePayload struct {
	// If it's been added/included
	Complete *bool `form:"complete,omitempty" json:"complete,omitempty" xml:"complete,omitempty"`
	// Amount of time to cook
	CookTime *time.Time `form:"cook_time,omitempty" json:"cook_time,omitempty" xml:"cook_time,omitempty"`
	// Long description of recipe
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// rating between 0-1
	Difficulty *float64 `form:"difficulty,omitempty" json:"difficulty,omitempty" xml:"difficulty,omitempty"`
	// Is a favorite, basically a tag
	Favorite *bool `form:"favorite,omitempty" json:"favorite,omitempty" xml:"favorite,omitempty"`
	// Image of recipe
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// Images of recipe
	Images []string `form:"images,omitempty" json:"images,omitempty" xml:"images,omitempty"`
	// List of ingredients
	Ingredients []*RecipePayload `form:"ingredients,omitempty" json:"ingredients,omitempty" xml:"ingredients,omitempty"`
	// Amount of time to prepare
	PrepTime *time.Time `form:"prep_time,omitempty" json:"prep_time,omitempty" xml:"prep_time,omitempty"`
	// rating between 0-1
	Rating *float64 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
	// Source of recipe
	Source *string `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// e.g. chopped, sliced, etc.. might need to be array.
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// Recipe Title
	Title string `form:"title" json:"title" xml:"title"`
	// Version Number e.g. 1.0.1
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Amount of time to wait for things such as mairnading
	WaitTime *time.Time `form:"wait_time,omitempty" json:"wait_time,omitempty" xml:"wait_time,omitempty"`
}

// Validate validates the RecipePayload type instance.
func (ut *RecipePayload) Validate() (err error) {
	if ut.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "title"))
	}
	if ut.Difficulty != nil {
		if *ut.Difficulty < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.difficulty`, *ut.Difficulty, 0.000000, true))
		}
	}
	if ut.Difficulty != nil {
		if *ut.Difficulty > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.difficulty`, *ut.Difficulty, 1.000000, false))
		}
	}
	for _, e := range ut.Ingredients {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.rating`, *ut.Rating, 0.000000, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.rating`, *ut.Rating, 1.000000, false))
		}
	}
	return
}
