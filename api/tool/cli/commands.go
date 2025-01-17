// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "recipe": CLI Commands
//
// Command:
// $ goagen
// --design=github.com/jaredwarren/recipe/api/design
// --out=$(GOPATH)/src/github.com/jaredwarren/recipe/api
// --version=v1.3.1

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/jaredwarren/recipe/api/client"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateRecipeCommand is the command line data structure for the create action of recipe
	CreateRecipeCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteRecipeCommand is the command line data structure for the delete action of recipe
	DeleteRecipeCommand struct {
		// Recipe ID
		ID          string
		PrettyPrint bool
	}

	// ListRecipeCommand is the command line data structure for the list action of recipe
	ListRecipeCommand struct {
		PrettyPrint bool
	}

	// ShowRecipeCommand is the command line data structure for the show action of recipe
	ShowRecipeCommand struct {
		// Recipe ID
		ID          string
		PrettyPrint bool
	}

	// UpdateRecipeCommand is the command line data structure for the update action of recipe
	UpdateRecipeCommand struct {
		Payload     string
		ContentType string
		// Recipe ID
		ID          string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: ``,
	}
	tmp1 := new(CreateRecipeCommand)
	sub = &cobra.Command{
		Use:   `recipe ["/recipe/"]`,
		Short: ``,
		Long: `

Payload example:

{
   "complete": false,
   "cook_time": "2003-08-13T12:47:16Z",
   "description": "Est facilis commodi ratione temporibus.",
   "difficulty": 0.5208787095685758,
   "favorite": false,
   "image": "Doloremque qui incidunt autem.",
   "images": [
      "Occaecati totam deleniti iusto quos."
   ],
   "ingredients": [
      {
         "complete": false,
         "cook_time": "2003-08-13T12:47:16Z",
         "description": "Est facilis commodi ratione temporibus.",
         "difficulty": 0.5208787095685758,
         "favorite": false,
         "image": "Doloremque qui incidunt autem.",
         "images": [
            "Occaecati totam deleniti iusto quos."
         ],
         "ingredients": [
            {
               "complete": false,
               "cook_time": "2003-08-13T12:47:16Z",
               "description": "Est facilis commodi ratione temporibus.",
               "difficulty": 0.5208787095685758,
               "favorite": false,
               "image": "Doloremque qui incidunt autem.",
               "images": [
                  "Occaecati totam deleniti iusto quos."
               ],
               "prep_time": "1990-08-15T00:54:39Z",
               "rating": 0.9356256018718694,
               "source": "Recusandae aut.",
               "state": "Laboriosam enim accusamus in libero.",
               "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
               "version": "Consequuntur occaecati ipsa qui.",
               "wait_time": "2004-11-23T22:16:37Z"
            }
         ],
         "prep_time": "1990-08-15T00:54:39Z",
         "rating": 0.9356256018718694,
         "source": "Recusandae aut.",
         "state": "Laboriosam enim accusamus in libero.",
         "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
         "version": "Consequuntur occaecati ipsa qui.",
         "wait_time": "2004-11-23T22:16:37Z"
      },
      {
         "complete": false,
         "cook_time": "2003-08-13T12:47:16Z",
         "description": "Est facilis commodi ratione temporibus.",
         "difficulty": 0.5208787095685758,
         "favorite": false,
         "image": "Doloremque qui incidunt autem.",
         "images": [
            "Occaecati totam deleniti iusto quos."
         ],
         "ingredients": [
            {
               "complete": false,
               "cook_time": "2003-08-13T12:47:16Z",
               "description": "Est facilis commodi ratione temporibus.",
               "difficulty": 0.5208787095685758,
               "favorite": false,
               "image": "Doloremque qui incidunt autem.",
               "images": [
                  "Occaecati totam deleniti iusto quos."
               ],
               "prep_time": "1990-08-15T00:54:39Z",
               "rating": 0.9356256018718694,
               "source": "Recusandae aut.",
               "state": "Laboriosam enim accusamus in libero.",
               "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
               "version": "Consequuntur occaecati ipsa qui.",
               "wait_time": "2004-11-23T22:16:37Z"
            }
         ],
         "prep_time": "1990-08-15T00:54:39Z",
         "rating": 0.9356256018718694,
         "source": "Recusandae aut.",
         "state": "Laboriosam enim accusamus in libero.",
         "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
         "version": "Consequuntur occaecati ipsa qui.",
         "wait_time": "2004-11-23T22:16:37Z"
      },
      {
         "complete": false,
         "cook_time": "2003-08-13T12:47:16Z",
         "description": "Est facilis commodi ratione temporibus.",
         "difficulty": 0.5208787095685758,
         "favorite": false,
         "image": "Doloremque qui incidunt autem.",
         "images": [
            "Occaecati totam deleniti iusto quos."
         ],
         "ingredients": [
            {
               "complete": false,
               "cook_time": "2003-08-13T12:47:16Z",
               "description": "Est facilis commodi ratione temporibus.",
               "difficulty": 0.5208787095685758,
               "favorite": false,
               "image": "Doloremque qui incidunt autem.",
               "images": [
                  "Occaecati totam deleniti iusto quos."
               ],
               "prep_time": "1990-08-15T00:54:39Z",
               "rating": 0.9356256018718694,
               "source": "Recusandae aut.",
               "state": "Laboriosam enim accusamus in libero.",
               "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
               "version": "Consequuntur occaecati ipsa qui.",
               "wait_time": "2004-11-23T22:16:37Z"
            }
         ],
         "prep_time": "1990-08-15T00:54:39Z",
         "rating": 0.9356256018718694,
         "source": "Recusandae aut.",
         "state": "Laboriosam enim accusamus in libero.",
         "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
         "version": "Consequuntur occaecati ipsa qui.",
         "wait_time": "2004-11-23T22:16:37Z"
      }
   ],
   "prep_time": "1990-08-15T00:54:39Z",
   "rating": 0.9356256018718694,
   "source": "Recusandae aut.",
   "state": "Laboriosam enim accusamus in libero.",
   "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
   "version": "Consequuntur occaecati ipsa qui.",
   "wait_time": "2004-11-23T22:16:37Z"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: ``,
	}
	tmp2 := new(DeleteRecipeCommand)
	sub = &cobra.Command{
		Use:   `recipe ["/recipe/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `List recipes`,
	}
	tmp3 := new(ListRecipeCommand)
	sub = &cobra.Command{
		Use:   `recipe ["/recipe/"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `Display an recipe by id`,
	}
	tmp4 := new(ShowRecipeCommand)
	sub = &cobra.Command{
		Use:   `recipe ["/recipe/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: ``,
	}
	tmp5 := new(UpdateRecipeCommand)
	sub = &cobra.Command{
		Use:   `recipe ["/recipe/ID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "complete": false,
   "cook_time": "2003-08-13T12:47:16Z",
   "description": "Est facilis commodi ratione temporibus.",
   "difficulty": 0.5208787095685758,
   "favorite": false,
   "image": "Doloremque qui incidunt autem.",
   "images": [
      "Occaecati totam deleniti iusto quos."
   ],
   "ingredients": [
      {
         "complete": false,
         "cook_time": "2003-08-13T12:47:16Z",
         "description": "Est facilis commodi ratione temporibus.",
         "difficulty": 0.5208787095685758,
         "favorite": false,
         "image": "Doloremque qui incidunt autem.",
         "images": [
            "Occaecati totam deleniti iusto quos."
         ],
         "ingredients": [
            {
               "complete": false,
               "cook_time": "2003-08-13T12:47:16Z",
               "description": "Est facilis commodi ratione temporibus.",
               "difficulty": 0.5208787095685758,
               "favorite": false,
               "image": "Doloremque qui incidunt autem.",
               "images": [
                  "Occaecati totam deleniti iusto quos."
               ],
               "prep_time": "1990-08-15T00:54:39Z",
               "rating": 0.9356256018718694,
               "source": "Recusandae aut.",
               "state": "Laboriosam enim accusamus in libero.",
               "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
               "version": "Consequuntur occaecati ipsa qui.",
               "wait_time": "2004-11-23T22:16:37Z"
            }
         ],
         "prep_time": "1990-08-15T00:54:39Z",
         "rating": 0.9356256018718694,
         "source": "Recusandae aut.",
         "state": "Laboriosam enim accusamus in libero.",
         "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
         "version": "Consequuntur occaecati ipsa qui.",
         "wait_time": "2004-11-23T22:16:37Z"
      },
      {
         "complete": false,
         "cook_time": "2003-08-13T12:47:16Z",
         "description": "Est facilis commodi ratione temporibus.",
         "difficulty": 0.5208787095685758,
         "favorite": false,
         "image": "Doloremque qui incidunt autem.",
         "images": [
            "Occaecati totam deleniti iusto quos."
         ],
         "ingredients": [
            {
               "complete": false,
               "cook_time": "2003-08-13T12:47:16Z",
               "description": "Est facilis commodi ratione temporibus.",
               "difficulty": 0.5208787095685758,
               "favorite": false,
               "image": "Doloremque qui incidunt autem.",
               "images": [
                  "Occaecati totam deleniti iusto quos."
               ],
               "prep_time": "1990-08-15T00:54:39Z",
               "rating": 0.9356256018718694,
               "source": "Recusandae aut.",
               "state": "Laboriosam enim accusamus in libero.",
               "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
               "version": "Consequuntur occaecati ipsa qui.",
               "wait_time": "2004-11-23T22:16:37Z"
            }
         ],
         "prep_time": "1990-08-15T00:54:39Z",
         "rating": 0.9356256018718694,
         "source": "Recusandae aut.",
         "state": "Laboriosam enim accusamus in libero.",
         "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
         "version": "Consequuntur occaecati ipsa qui.",
         "wait_time": "2004-11-23T22:16:37Z"
      },
      {
         "complete": false,
         "cook_time": "2003-08-13T12:47:16Z",
         "description": "Est facilis commodi ratione temporibus.",
         "difficulty": 0.5208787095685758,
         "favorite": false,
         "image": "Doloremque qui incidunt autem.",
         "images": [
            "Occaecati totam deleniti iusto quos."
         ],
         "ingredients": [
            {
               "complete": false,
               "cook_time": "2003-08-13T12:47:16Z",
               "description": "Est facilis commodi ratione temporibus.",
               "difficulty": 0.5208787095685758,
               "favorite": false,
               "image": "Doloremque qui incidunt autem.",
               "images": [
                  "Occaecati totam deleniti iusto quos."
               ],
               "prep_time": "1990-08-15T00:54:39Z",
               "rating": 0.9356256018718694,
               "source": "Recusandae aut.",
               "state": "Laboriosam enim accusamus in libero.",
               "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
               "version": "Consequuntur occaecati ipsa qui.",
               "wait_time": "2004-11-23T22:16:37Z"
            }
         ],
         "prep_time": "1990-08-15T00:54:39Z",
         "rating": 0.9356256018718694,
         "source": "Recusandae aut.",
         "state": "Laboriosam enim accusamus in libero.",
         "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
         "version": "Consequuntur occaecati ipsa qui.",
         "wait_time": "2004-11-23T22:16:37Z"
      }
   ],
   "prep_time": "1990-08-15T00:54:39Z",
   "rating": 0.9356256018718694,
   "source": "Recusandae aut.",
   "state": "Laboriosam enim accusamus in libero.",
   "title": "Itaque maiores quasi corrupti rerum aliquid reiciendis.",
   "version": "Consequuntur occaecati ipsa qui.",
   "wait_time": "2004-11-23T22:16:37Z"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the CreateRecipeCommand command.
func (cmd *CreateRecipeCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/recipe/"
	}
	var payload client.CreateRecipePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateRecipe(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateRecipeCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteRecipeCommand command.
func (cmd *DeleteRecipeCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/recipe/%v", url.QueryEscape(cmd.ID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteRecipe(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteRecipeCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `Recipe ID`)
}

// Run makes the HTTP request corresponding to the ListRecipeCommand command.
func (cmd *ListRecipeCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/recipe/"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListRecipe(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListRecipeCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowRecipeCommand command.
func (cmd *ShowRecipeCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/recipe/%v", url.QueryEscape(cmd.ID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowRecipe(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowRecipeCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `Recipe ID`)
}

// Run makes the HTTP request corresponding to the UpdateRecipeCommand command.
func (cmd *UpdateRecipeCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/recipe/%v", url.QueryEscape(cmd.ID))
	}
	var payload client.UpdateRecipePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateRecipe(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateRecipeCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `Recipe ID`)
}
