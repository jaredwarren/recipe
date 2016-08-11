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
	"strconv"
)

// ShowRecipeContext provides the recipe show action context.
type ShowRecipeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	RecipeID int
}

// NewShowRecipeContext parses the incoming request URL and body, performs validations and creates the
// context used by the recipe controller show action.
func NewShowRecipeContext(ctx context.Context, service *goa.Service) (*ShowRecipeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowRecipeContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramRecipeID := req.Params["recipeID"]
	if len(paramRecipeID) > 0 {
		rawRecipeID := paramRecipeID[0]
		if recipeID, err2 := strconv.Atoi(rawRecipeID); err2 == nil {
			rctx.RecipeID = recipeID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("recipeID", rawRecipeID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowRecipeContext) OK(r *RecipeRecipe) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/recipe.recipe+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowRecipeContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
