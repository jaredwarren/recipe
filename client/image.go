package client

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

// DownloadDownload downloads /files with the given filename and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadDownload(ctx context.Context, filename, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	p := path.Join("/download/", filename)
	u := url.URL{Host: c.Host, Scheme: scheme, Path: p}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}

// ShowImagePath computes a request path to the show action of image.
func ShowImagePath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/recipe/images/%s", param0)
}

// Show an image metadata
func (c *Client) ShowImage(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowImageRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowImageRequest create the request corresponding to the show action endpoint of the image resource.
func (c *Client) NewShowImageRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UploadImagePath computes a request path to the upload action of image.
func UploadImagePath() string {

	return fmt.Sprintf("/recipe/images")
}

// Upload an image
func (c *Client) UploadImage(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewUploadImageRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUploadImageRequest create the request corresponding to the upload action endpoint of the image resource.
func (c *Client) NewUploadImageRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
