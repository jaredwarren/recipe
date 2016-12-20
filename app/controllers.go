//************************************************************************//
// API "recipe": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ImageController is the controller interface for the Image actions.
type ImageController interface {
	goa.Muxer
	goa.FileServer
	Show(*ShowImageContext) error
	Upload(*UploadImageContext) error
}

// MountImageController "mounts" a Image resource controller on the given service.
func MountImageController(service *goa.Service, ctrl ImageController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowImageContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/recipe/images/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "action", "Show", "route", "GET /recipe/images/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUploadImageContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Upload(rctx)
	}
	service.Mux.Handle("POST", "/recipe/images", ctrl.MuxHandler("Upload", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "action", "Upload", "route", "POST /recipe/images")

	h = ctrl.FileHandler("/download/*filename", "images/")
	service.Mux.Handle("GET", "/download/*filename", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "files", "images/", "route", "GET /download/*filename")

	h = ctrl.FileHandler("/download/", "images/index.html")
	service.Mux.Handle("GET", "/download/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "files", "images/index.html", "route", "GET /download/")
}

// RecipeController is the controller interface for the Recipe actions.
type RecipeController interface {
	goa.Muxer
	Create(*CreateRecipeContext) error
	Delete(*DeleteRecipeContext) error
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
		rctx, err := NewCreateRecipeContext(ctx, service)
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
	service.Mux.Handle("POST", "/recipe/recipe", ctrl.MuxHandler("Create", h, unmarshalCreateRecipePayload))
	service.LogInfo("mount", "ctrl", "Recipe", "action", "Create", "route", "POST /recipe/recipe")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteRecipeContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/recipe/recipe/:id", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Recipe", "action", "Delete", "route", "DELETE /recipe/recipe/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowRecipeContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/recipe/recipe/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Recipe", "action", "Show", "route", "GET /recipe/recipe/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateRecipeContext(ctx, service)
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
	service.Mux.Handle("PATCH", "/recipe/recipe/:id", ctrl.MuxHandler("Update", h, unmarshalUpdateRecipePayload))
	service.LogInfo("mount", "ctrl", "Recipe", "action", "Update", "route", "PATCH /recipe/recipe/:id")
}

// unmarshalCreateRecipePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateRecipePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createRecipePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
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
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
