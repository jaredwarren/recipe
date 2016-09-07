package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateRecipePayload is the recipe create action payload.
type CreateRecipePayload struct {
	// Operand value
	Value int `form:"value" json:"value" xml:"value"`
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
	return fmt.Sprintf("/recipe/recipe/%v", id)
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
	return fmt.Sprintf("/recipe/recipe/%v", id)
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
	// Operand value
	Value int `form:"value" json:"value" xml:"value"`
}

// UpdateRecipePath computes a request path to the update action of recipe.
func UpdateRecipePath(id string) string {
	return fmt.Sprintf("/recipe/recipe/%v", id)
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
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
