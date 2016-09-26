package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/recipe/app"
)

// ImageController implements the image resource.
type ImageController struct {
	*goa.Controller
}

// NewImageController creates a image controller.
func NewImageController(service *goa.Service) *ImageController {
	return &ImageController{Controller: service.NewController("ImageController")}
}

// Show runs the show action.
func (c *ImageController) Show(ctx *app.ShowImageContext) error {
	// ImageController_Show: start_implement

	// Put your logic here

	// ImageController_Show: end_implement
	res := &app.ImageMedia{}
	return ctx.OK(res)
}

// Upload runs the upload action.
func (c *ImageController) Upload(ctx *app.UploadImageContext) error {
	// ImageController_Upload: start_implement

	// Put your logic here

	// ImageController_Upload: end_implement
	res := &app.ImageMedia{}
	return ctx.OK(res)
}
