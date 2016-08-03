package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ShowRecipePath computes a request path to the show action of recipe.
func ShowRecipePath(recipeID int) string {
	return fmt.Sprintf("/recipes/%v", recipeID)
}

// Get recipe by id
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
