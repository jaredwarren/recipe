package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"time"
)

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

// CreateRecipePath computes a request path to the create action of recipe.
func CreateRecipePath() string {

	return fmt.Sprintf("/recipe/recipe")
}

// CreateRecipe makes a request to the create action endpoint of the recipe resource
func (c *Client) CreateRecipe(ctx context.Context, path string, payload *CreateRecipePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateRecipeRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateRecipeRequest create the request corresponding to the create action endpoint of the recipe resource.
func (c *Client) NewCreateRecipeRequest(ctx context.Context, path string, payload *CreateRecipePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteRecipePath computes a request path to the delete action of recipe.
func DeleteRecipePath(id string) string {
	param0 := id

	return fmt.Sprintf("/recipe/recipe/%s", param0)
}

// DeleteRecipe makes a request to the delete action endpoint of the recipe resource
func (c *Client) DeleteRecipe(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteRecipeRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteRecipeRequest create the request corresponding to the delete action endpoint of the recipe resource.
func (c *Client) NewDeleteRecipeRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowRecipePath computes a request path to the show action of recipe.
func ShowRecipePath(id string) string {
	param0 := id

	return fmt.Sprintf("/recipe/recipe/%s", param0)
}

// Display an recipe by id
func (c *Client) ShowRecipe(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowRecipeRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowRecipeRequest create the request corresponding to the show action endpoint of the recipe resource.
func (c *Client) NewShowRecipeRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
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

// UpdateRecipePath computes a request path to the update action of recipe.
func UpdateRecipePath(id string) string {
	param0 := id

	return fmt.Sprintf("/recipe/recipe/%s", param0)
}

// UpdateRecipe makes a request to the update action endpoint of the recipe resource
func (c *Client) UpdateRecipe(ctx context.Context, path string, payload *UpdateRecipePayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateRecipeRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateRecipeRequest create the request corresponding to the update action endpoint of the recipe resource.
func (c *Client) NewUpdateRecipeRequest(ctx context.Context, path string, payload *UpdateRecipePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
