//************************************************************************//
// API "recipe": Application Contexts
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/jaredwarren/recipe/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe
// --version=v0.2.dev
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
	// Operand value
	Value *int `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createRecipePayload) Validate() (err error) {
	if payload.Value == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "value"))
	}

	return
}

// Publicize creates CreateRecipePayload from createRecipePayload
func (payload *createRecipePayload) Publicize() *CreateRecipePayload {
	var pub CreateRecipePayload
	if payload.Value != nil {
		pub.Value = *payload.Value
	}
	return &pub
}

// CreateRecipePayload is the recipe create action payload.
type CreateRecipePayload struct {
	// Operand value
	Value int `form:"value" json:"value" xml:"value"`
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

// ListRecipeContext provides the recipe list action context.
type ListRecipeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListRecipeContext parses the incoming request URL and body, performs validations and creates the
// context used by the recipe controller list action.
func NewListRecipeContext(ctx context.Context, service *goa.Service) (*ListRecipeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListRecipeContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListRecipeContext) OK(r []*RecipeRecipe) error {
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
	// Operand value
	Value *int `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateRecipePayload) Validate() (err error) {
	if payload.Value == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "value"))
	}

	return
}

// Publicize creates UpdateRecipePayload from updateRecipePayload
func (payload *updateRecipePayload) Publicize() *UpdateRecipePayload {
	var pub UpdateRecipePayload
	if payload.Value != nil {
		pub.Value = *payload.Value
	}
	return &pub
}

// UpdateRecipePayload is the recipe update action payload.
type UpdateRecipePayload struct {
	// Operand value
	Value int `form:"value" json:"value" xml:"value"`
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
