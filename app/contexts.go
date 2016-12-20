//************************************************************************//
// API "recipe": Application Contexts
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
	"golang.org/x/net/context"
	"strconv"
	"time"
)

// ShowImageContext provides the image show action context.
type ShowImageContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowImageContext parses the incoming request URL and body, performs validations and creates the
// context used by the image controller show action.
func NewShowImageContext(ctx context.Context, service *goa.Service) (*ShowImageContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowImageContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowImageContext) OK(r *ImageMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.image+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowImageContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UploadImageContext provides the image upload action context.
type UploadImageContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewUploadImageContext parses the incoming request URL and body, performs validations and creates the
// context used by the image controller upload action.
func NewUploadImageContext(ctx context.Context, service *goa.Service) (*UploadImageContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UploadImageContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UploadImageContext) OK(r *ImageMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.image+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// CreateRecipeContext provides the recipe create action context.
type CreateRecipeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateRecipePayload
}

// NewCreateRecipeContext parses the incoming request URL and body, performs validations and creates the
// context used by the recipe controller create action.
func NewCreateRecipeContext(ctx context.Context, service *goa.Service) (*CreateRecipeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateRecipeContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createRecipePayload is the recipe create action payload.
type createRecipePayload struct {
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

// Finalize sets the default values defined in the design.
func (payload *createRecipePayload) Finalize() {
	if payload.Quantity != nil {
		var defaultType = "weight"
		if payload.Quantity.Type == nil {
			payload.Quantity.Type = &defaultType
		}
	}
}

// Validate runs the validation rules defined in the design.
func (payload *createRecipePayload) Validate() (err error) {
	if payload.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}

	if payload.Difficulty != nil {
		if *payload.Difficulty < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 0, true))
		}
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 1, false))
		}
	}
	if payload.Quantity != nil {
		if err2 := payload.Quantity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 0, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1, false))
		}
	}
	if payload.Source != nil {
		if err2 := payload.Source.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates CreateRecipePayload from createRecipePayload
func (payload *createRecipePayload) Publicize() *CreateRecipePayload {
	var pub CreateRecipePayload
	if payload.Complete != nil {
		pub.Complete = payload.Complete
	}
	if payload.CookTime != nil {
		pub.CookTime = payload.CookTime
	}
	if payload.Description != nil {
		pub.Description = payload.Description
	}
	if payload.Difficulty != nil {
		pub.Difficulty = payload.Difficulty
	}
	if payload.Favorite != nil {
		pub.Favorite = payload.Favorite
	}
	if payload.Image != nil {
		pub.Image = payload.Image
	}
	if payload.Images != nil {
		pub.Images = payload.Images
	}
	if payload.PrepTime != nil {
		pub.PrepTime = payload.PrepTime
	}
	if payload.Quantity != nil {
		pub.Quantity = payload.Quantity.Publicize()
	}
	if payload.Rating != nil {
		pub.Rating = payload.Rating
	}
	if payload.Source != nil {
		pub.Source = payload.Source.Publicize()
	}
	if payload.State != nil {
		pub.State = payload.State
	}
	if payload.Title != nil {
		pub.Title = *payload.Title
	}
	if payload.Version != nil {
		pub.Version = payload.Version
	}
	if payload.WaitTime != nil {
		pub.WaitTime = payload.WaitTime
	}
	return &pub
}

// CreateRecipePayload is the recipe create action payload.
type CreateRecipePayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *CreateRecipePayload) Validate() (err error) {
	if payload.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}

	if payload.Difficulty != nil {
		if *payload.Difficulty < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 0, true))
		}
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 1, false))
		}
	}
	if payload.Quantity != nil {
		if err2 := payload.Quantity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 0, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1, false))
		}
	}
	if payload.Source != nil {
		if err2 := payload.Source.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateRecipeContext) OK(r *RecipeRecipe) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.recipe+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKIngredient sends a HTTP response with status code 200.
func (ctx *CreateRecipeContext) OKIngredient(r *RecipeRecipeIngredient) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.recipe+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateRecipeContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateRecipeContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// DeleteRecipeContext provides the recipe delete action context.
type DeleteRecipeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewDeleteRecipeContext parses the incoming request URL and body, performs validations and creates the
// context used by the recipe controller delete action.
func NewDeleteRecipeContext(ctx context.Context, service *goa.Service) (*DeleteRecipeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteRecipeContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteRecipeContext) OK(r *RecipeRecipe) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.recipe+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKIngredient sends a HTTP response with status code 200.
func (ctx *DeleteRecipeContext) OKIngredient(r *RecipeRecipeIngredient) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.recipe+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteRecipeContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteRecipeContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ShowRecipeContext provides the recipe show action context.
type ShowRecipeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowRecipeContext parses the incoming request URL and body, performs validations and creates the
// context used by the recipe controller show action.
func NewShowRecipeContext(ctx context.Context, service *goa.Service) (*ShowRecipeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowRecipeContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowRecipeContext) OK(r *RecipeRecipe) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.recipe+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKIngredient sends a HTTP response with status code 200.
func (ctx *ShowRecipeContext) OKIngredient(r *RecipeRecipeIngredient) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.recipe+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowRecipeContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateRecipeContext provides the recipe update action context.
type UpdateRecipeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *UpdateRecipePayload
}

// NewUpdateRecipeContext parses the incoming request URL and body, performs validations and creates the
// context used by the recipe controller update action.
func NewUpdateRecipeContext(ctx context.Context, service *goa.Service) (*UpdateRecipeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateRecipeContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// updateRecipePayload is the recipe update action payload.
type updateRecipePayload struct {
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

// Finalize sets the default values defined in the design.
func (payload *updateRecipePayload) Finalize() {
	if payload.Quantity != nil {
		var defaultType = "weight"
		if payload.Quantity.Type == nil {
			payload.Quantity.Type = &defaultType
		}
	}
}

// Validate runs the validation rules defined in the design.
func (payload *updateRecipePayload) Validate() (err error) {
	if payload.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}

	if payload.Difficulty != nil {
		if *payload.Difficulty < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 0, true))
		}
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 1, false))
		}
	}
	if payload.Quantity != nil {
		if err2 := payload.Quantity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 0, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1, false))
		}
	}
	if payload.Source != nil {
		if err2 := payload.Source.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates UpdateRecipePayload from updateRecipePayload
func (payload *updateRecipePayload) Publicize() *UpdateRecipePayload {
	var pub UpdateRecipePayload
	if payload.Complete != nil {
		pub.Complete = payload.Complete
	}
	if payload.CookTime != nil {
		pub.CookTime = payload.CookTime
	}
	if payload.Description != nil {
		pub.Description = payload.Description
	}
	if payload.Difficulty != nil {
		pub.Difficulty = payload.Difficulty
	}
	if payload.Favorite != nil {
		pub.Favorite = payload.Favorite
	}
	if payload.Image != nil {
		pub.Image = payload.Image
	}
	if payload.Images != nil {
		pub.Images = payload.Images
	}
	if payload.PrepTime != nil {
		pub.PrepTime = payload.PrepTime
	}
	if payload.Quantity != nil {
		pub.Quantity = payload.Quantity.Publicize()
	}
	if payload.Rating != nil {
		pub.Rating = payload.Rating
	}
	if payload.Source != nil {
		pub.Source = payload.Source.Publicize()
	}
	if payload.State != nil {
		pub.State = payload.State
	}
	if payload.Title != nil {
		pub.Title = *payload.Title
	}
	if payload.Version != nil {
		pub.Version = payload.Version
	}
	if payload.WaitTime != nil {
		pub.WaitTime = payload.WaitTime
	}
	return &pub
}

// UpdateRecipePayload is the recipe update action payload.
type UpdateRecipePayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *UpdateRecipePayload) Validate() (err error) {
	if payload.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}

	if payload.Difficulty != nil {
		if *payload.Difficulty < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 0, true))
		}
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 1, false))
		}
	}
	if payload.Quantity != nil {
		if err2 := payload.Quantity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 0, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1, false))
		}
	}
	if payload.Source != nil {
		if err2 := payload.Source.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateRecipeContext) OK(r *RecipeRecipe) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.recipe+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKIngredient sends a HTTP response with status code 200.
func (ctx *UpdateRecipeContext) OKIngredient(r *RecipeRecipeIngredient) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.recipe+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateRecipeContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateRecipeContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateRecipeContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateRecipeContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
