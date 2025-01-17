// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "recipe": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/jaredwarren/recipe/web/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// ShowImageContext provides the image show action context.
type ShowImageContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowImageContext parses the incoming request URL and body, performs validations and creates the
// context used by the image controller show action.
func NewShowImageContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowImageContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowImageContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowImageContext) OK(r *ImageMedia) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/recipe.image+json")
	}
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
func NewUploadImageContext(ctx context.Context, r *http.Request, service *goa.Service) (*UploadImageContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UploadImageContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UploadImageContext) OK(r *ImageMedia) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/recipe.image+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// CreateWebContext provides the web create action context.
type CreateWebContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateWebPayload
}

// NewCreateWebContext parses the incoming request URL and body, performs validations and creates the
// context used by the web controller create action.
func NewCreateWebContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateWebContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateWebContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createWebPayload is the web create action payload.
type createWebPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *createWebPayload) Validate() (err error) {
	if payload.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 0.000000, true))
		}
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 1.000000, false))
		}
	}
	for _, e := range payload.Ingredients {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 0.000000, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1.000000, false))
		}
	}
	return
}

// Publicize creates CreateWebPayload from createWebPayload
func (payload *createWebPayload) Publicize() *CreateWebPayload {
	var pub CreateWebPayload
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
	if payload.Ingredients != nil {
		pub.Ingredients = make([]*RecipePayload, len(payload.Ingredients))
		for i2, elem2 := range payload.Ingredients {
			pub.Ingredients[i2] = elem2.Publicize()
		}
	}
	if payload.PrepTime != nil {
		pub.PrepTime = payload.PrepTime
	}
	if payload.Rating != nil {
		pub.Rating = payload.Rating
	}
	if payload.Source != nil {
		pub.Source = payload.Source
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

// CreateWebPayload is the web create action payload.
type CreateWebPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *CreateWebPayload) Validate() (err error) {
	if payload.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 0.000000, true))
		}
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 1.000000, false))
		}
	}
	for _, e := range payload.Ingredients {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 0.000000, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1.000000, false))
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateWebContext) OK(r *RecipeRecipe) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKIngredient sends a HTTP response with status code 200.
func (ctx *CreateWebContext) OKIngredient(r *RecipeRecipeIngredient) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateWebContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateWebContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// DeleteWebContext provides the web delete action context.
type DeleteWebContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewDeleteWebContext parses the incoming request URL and body, performs validations and creates the
// context used by the web controller delete action.
func NewDeleteWebContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteWebContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteWebContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteWebContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteWebContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteWebContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ListWebContext provides the web list action context.
type ListWebContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListWebContext parses the incoming request URL and body, performs validations and creates the
// context used by the web controller list action.
func NewListWebContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListWebContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListWebContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListWebContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/html")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListWebContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ShowWebContext provides the web show action context.
type ShowWebContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowWebContext parses the incoming request URL and body, performs validations and creates the
// context used by the web controller show action.
func NewShowWebContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowWebContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowWebContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowWebContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/html")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// Created sends a HTTP response with status code 201.
func (ctx *ShowWebContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowWebContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowWebContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// UpdateWebContext provides the web update action context.
type UpdateWebContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *UpdateWebPayload
}

// NewUpdateWebContext parses the incoming request URL and body, performs validations and creates the
// context used by the web controller update action.
func NewUpdateWebContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateWebContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateWebContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// updateWebPayload is the web update action payload.
type updateWebPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *updateWebPayload) Validate() (err error) {
	if payload.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 0.000000, true))
		}
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 1.000000, false))
		}
	}
	for _, e := range payload.Ingredients {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 0.000000, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1.000000, false))
		}
	}
	return
}

// Publicize creates UpdateWebPayload from updateWebPayload
func (payload *updateWebPayload) Publicize() *UpdateWebPayload {
	var pub UpdateWebPayload
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
	if payload.Ingredients != nil {
		pub.Ingredients = make([]*RecipePayload, len(payload.Ingredients))
		for i2, elem2 := range payload.Ingredients {
			pub.Ingredients[i2] = elem2.Publicize()
		}
	}
	if payload.PrepTime != nil {
		pub.PrepTime = payload.PrepTime
	}
	if payload.Rating != nil {
		pub.Rating = payload.Rating
	}
	if payload.Source != nil {
		pub.Source = payload.Source
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

// UpdateWebPayload is the web update action payload.
type UpdateWebPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *UpdateWebPayload) Validate() (err error) {
	if payload.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 0.000000, true))
		}
	}
	if payload.Difficulty != nil {
		if *payload.Difficulty > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.difficulty`, *payload.Difficulty, 1.000000, false))
		}
	}
	for _, e := range payload.Ingredients {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 0.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 0.000000, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 1.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.rating`, *payload.Rating, 1.000000, false))
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateWebContext) OK(r *RecipeRecipe) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKIngredient sends a HTTP response with status code 200.
func (ctx *UpdateWebContext) OKIngredient(r *RecipeRecipeIngredient) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateWebContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateWebContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateWebContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateWebContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
