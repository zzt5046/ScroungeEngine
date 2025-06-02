package api

type GenerateRecipeParams struct {
	// List of available ingredients
	Ingredients []string `json:"ingredients"`

	// List of preferences (eg. cuisine type, complexity, anything really)
	Preferences []string `json:"preferences"`
}

// Basic structure of generated recipe
type Recipe struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	Notes        string   `json:"notes"`
}

// Response when generating a recipe
type GenerateRecipeResponse struct {
	Code                int      `json:"code"`
	Recipe              Recipe   `json:"recipe"`
	LeftoverIngredients []string `json:"leftover_ingredients"`
}

// Response when generating multiple recipes
type GenerateRecipesResponse struct {
	Code                int      `json:"code"`
	Recipes             []Recipe `json:"recipes"`
	LeftoverIngredients []string `json:"leftover_ingredients"`
}

// Error response
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
