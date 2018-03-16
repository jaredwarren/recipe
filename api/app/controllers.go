// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "recipe": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/jaredwarren/recipe/api/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe/api
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// RecipeController is the controller interface for the Recipe actions.
type RecipeController interface {
	goa.Muxer
	Create(*CreateRecipeContext) error
	Delete(*DeleteRecipeContext) error
	List(*ListRecipeContext) error
	Show(*ShowRecipeContext) error
	Update(*UpdateRecipeContext) error
}

// MountRecipeController "mounts" a Recipe resource controller on the given service.
func MountRecipeController(service *goa.Service, ctrl RecipeController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateRecipeContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateRecipePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/recipe/", ctrl.MuxHandler("create", h, unmarshalCreateRecipePayload))
	service.LogInfo("mount", "ctrl", "Recipe", "action", "Create", "route", "POST /recipe/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteRecipeContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/recipe/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Recipe", "action", "Delete", "route", "DELETE /recipe/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListRecipeContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/recipe/", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Recipe", "action", "List", "route", "GET /recipe/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowRecipeContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/recipe/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Recipe", "action", "Show", "route", "GET /recipe/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateRecipeContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateRecipePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PATCH", "/recipe/:id", ctrl.MuxHandler("update", h, unmarshalUpdateRecipePayload))
	service.LogInfo("mount", "ctrl", "Recipe", "action", "Update", "route", "PATCH /recipe/:id")
}

// unmarshalCreateRecipePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateRecipePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createRecipePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateRecipePayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateRecipePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateRecipePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}