package api

// REQUESTS TO ENGINE

type GenerateRecipeRequest struct {
	// List of available ingredients
	Ingredients []string `json:"ingredients"`

	// List of preferences (eg. cuisine type, complexity, anything really)
	Preferences []string `json:"preferences"`

	// How many recipes to generate
	Count int `json:"count"`
}


// RESPONSES FROM ENGINE -----------------------------------------------------------

// Response when generating a recipe
type GenerateRecipeResponse struct {
	Code                int      `json:"code"`
	Recipe              Recipe   `json:"recipe"`
	// LeftoverIngredients []string `json:"leftover_ingredients"`
}

// Response when generating multiple recipes
type GenerateRecipesResponse struct {
	Code                int      `json:"code"`
	Recipes             []Recipe `json:"recipes"`
	// LeftoverIngredients []string `json:"leftover_ingredients"`
}

// Error response
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}


// REQUESTS TO LLAMA----------------------------------------------------

type LlamaRequest struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
	Format string `json:"format"`
	Stream bool `json:"stream"`
}


// RESPONSES FROM LLAMA-------------------------------------------------

// General response object from ollama api
type LlamaResponse struct {
	Model string `json:"model"`
	CreatedTimestamp string `json:"created_at"`
	Response string `json:"response"`
	Done bool `json:"done"`
	DoneReason string `json:"done_reason"`
	Context []int `json:"context"`
}

// Structure of generated recipe ("Response" from LlamaResponse)
type Recipe struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	Notes        string   `json:"notes"`
}