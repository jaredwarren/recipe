package main

import (
	"database/sql"
	"fmt"
	"time"
)

// RecipeStorage :
type RecipeStorage interface {
	Type() string        // Return a string describing the storage type.
	Load(string) []byte  // Load a recipe by name.
	Save(string, []byte) // Save a recipe by name.
}

// IngredientStorage :
type IngredientStorage interface {
	Type() string        // Return a string describing the storage type.
	Load(string) []byte  // Load a ingredient by name.
	Save(string, []byte) // Save a ingredient by name.
}

// Ingredient stuff
type Ingredient struct {
	Name     string
	Quantity int
	Units    string
}

// TODO: convert to different units
// TODO: if i have density, I can convert from volume to weight

// UnitConverter ...
type UnitConverter interface {
	Convert(f, t UnitOfMeasure) UnitOfMeasure
}

// VolumeUnitConverter ...
type VolumeUnitConverter struct {
	test string
}

// Convert ...
func (c *VolumeUnitConverter) Convert(f, t UnitOfMeasure) UnitOfMeasure {
	return UnitOfMeasure{}
}

// VolumeToWeightUnitConverter ...
type VolumeToWeightUnitConverter struct {
	test string
}

// Convert ...
func (c *VolumeToWeightUnitConverter) Convert(f, t UnitOfMeasure) UnitOfMeasure {
	return UnitOfMeasure{}
}

// UnitOfMeasure ...
type UnitOfMeasure struct {
	Name         string  // name e.g. Cups, Table Spoon
	Abbreviation string  // abbreviation e.g C, Tb
	Type         string  // volume, weight
	Diff         float32 // difference from SI
}

// Convert ingredient, probably shouldn't  be part of intredient, but UnitOfMeasure
func (n *Ingredient) Convert(f, t UnitOfMeasure) bool {
	_, month, day := time.Now().Date()
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
	return true
}

// URL ...
type URL string

// HTML ...
type HTML string

// Recipe is a "PoemStorage" for ingredients
type Recipe struct {
	Title                string
	Images               []URL  // image path
	Servings             int    // should be quantity and measure e.g. 4 cups. same as ingredients
	NutritionInformation string // TODO: make struct
	PrepTime             time.Time
	CookTime             time.Time
	Cookware             []string
	Description          HTML
	Version              string // git sha TODO: need a good way to serialize data, dif-merge-patch, marshall/unmarshall
	Ingredients          []Ingredient
	Recipes              []*Recipe // sub recipes
	Directions           []Step    // html or []steps
	Categories           []string  // array of Categories; Categories is interface of tag
	Favorite             bool
	Tags                 []Tag // Categories
	Rating               int   // 0-1
	Difficulity          int   // 0-1
	Source               string
	SourceURL            URL
	Notes                []Note
}

// Share recipe
func (n *Recipe) Share() bool {
	_, month, day := time.Now().Date()
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
	return true
}

// Step ...
type Step struct {
	Instructions HTML
}

// Note ...
type Note struct {
	Text string
	Date time.Time
}

// Tag ...
type Tag struct {
	Name string
	Type string
}

// Calendar

// Calendar ...
type Calendar struct {
	Name string
	Type string
}

// Meal ...
type Meal struct {
	Recipes []Recipe
	// TODO: what eles goes with a meal. # servings needed???
}

// Shopping List

// ShoppingList ...
type ShoppingList struct {
	Items []ShoppingItem
}

// AddIngredient ingredient
func (n *ShoppingList) AddIngredient(ingredient *Ingredient) bool {
	// TODO: search for same ingredient and combile total. might have to copy *ingredient into n.Ingredients
	return true
}

// ShoppingItem ...
type ShoppingItem struct {
	Done          bool
	Ingredients   []Ingredient
	Recipe        Recipe // What recipe does this item belong to
	Aile          string
	Department    string
	LastPruchased time.Time
	DateAdded     time.Time
}

//
var alphaDB *sql.DB

func mainish() {
	fmt.Printf("hello, world\n")
}
