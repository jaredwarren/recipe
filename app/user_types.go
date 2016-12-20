//************************************************************************//
// API "recipe": Application User Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/jaredwarren/recipe/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

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
	// Amount of time to prepare
	PrepTime *time.Time `form:"prep_time,omitempty" json:"prep_time,omitempty" xml:"prep_time,omitempty"`
	// quantity, measure, servings, yield e.g. 4 cups.
	Quantity *recipeUnitofmeasure `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// rating between 0-1
	Rating *float64 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
	// Source of recipe
	Source *recipeSource `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// e.g. chopped, sliced, etc.. might need to be array.
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// Recipe Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Version Number e.g. 1.0.1
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Amount of time to wait for things such as mairnading
	WaitTime *time.Time `form:"wait_time,omitempty" json:"wait_time,omitempty" xml:"wait_time,omitempty"`
}

// Finalize sets the default values for recipePayload type instance.
func (ut *recipePayload) Finalize() {
	if ut.Quantity != nil {
		var defaultType = "weight"
		if ut.Quantity.Type == nil {
			ut.Quantity.Type = &defaultType
		}
	}
}

// Validate validates the recipePayload type instance.
func (ut *recipePayload) Validate() (err error) {
	if ut.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	if ut.Difficulty != nil {
		if *ut.Difficulty < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.difficulty`, *ut.Difficulty, 0, true))
		}
	}
	if ut.Difficulty != nil {
		if *ut.Difficulty > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.difficulty`, *ut.Difficulty, 1, false))
		}
	}
	if ut.Quantity != nil {
		if err2 := ut.Quantity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 0, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 1, false))
		}
	}
	if ut.Source != nil {
		if err2 := ut.Source.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
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
	if ut.PrepTime != nil {
		pub.PrepTime = ut.PrepTime
	}
	if ut.Quantity != nil {
		pub.Quantity = ut.Quantity.Publicize()
	}
	if ut.Rating != nil {
		pub.Rating = ut.Rating
	}
	if ut.Source != nil {
		pub.Source = ut.Source.Publicize()
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
	// Amount of time to prepare
	PrepTime *time.Time `form:"prep_time,omitempty" json:"prep_time,omitempty" xml:"prep_time,omitempty"`
	// quantity, measure, servings, yield e.g. 4 cups.
	Quantity *RecipeUnitofmeasure `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// rating between 0-1
	Rating *float64 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
	// Source of recipe
	Source *RecipeSource `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
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
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	if ut.Difficulty != nil {
		if *ut.Difficulty < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.difficulty`, *ut.Difficulty, 0, true))
		}
	}
	if ut.Difficulty != nil {
		if *ut.Difficulty > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.difficulty`, *ut.Difficulty, 1, false))
		}
	}
	if ut.Quantity != nil {
		if err2 := ut.Quantity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 0, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 1, false))
		}
	}
	if ut.Source != nil {
		if err2 := ut.Source.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
