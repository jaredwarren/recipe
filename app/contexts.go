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
)

// CreateRecipeContext provides the recipe create action context.
type CreateRecipeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *RecipePayload
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
	Payload *RecipePayload
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
