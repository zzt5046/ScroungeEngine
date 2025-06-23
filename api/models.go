package api

// REQUESTS TO ENGINE

type GenerateRecipesRequest struct {
	// List of available ingredients
	Ingredients []string `json:"ingredients"`

	// List of preferences (eg. cuisine type, complexity, anything really)
	Preferences []string `json:"preferences"`

	// How many recipes to generate
	Count int `json:"count"`
}

// RESPONSES FROM ENGINE -----------------------------------------------------------

// Response when generating multiple recipes
type GenerateRecipesResponse struct {
	Recipes []Recipe `json:"recipes"`
	Error   string   `json:"error"`
}

// REQUESTS TO LLAMA----------------------------------------------------

type LlamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	// Format []byte `json:"format"`
	Format string `json:"format"`
	Stream bool   `json:"stream"`
}

// RESPONSES FROM LLAMA-------------------------------------------------

// General response object from ollama api
type LlamaResponse struct {
	Model            string `json:"model"`
	CreatedTimestamp string `json:"created_at"`
	Response         string `json:"response"`
	Done             bool   `json:"done"`
	DoneReason       string `json:"done_reason"`
	Context          []int  `json:"context"`
}

// Structure of generated recipe ("Response" from LlamaResponse)
type Recipe struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Directions  []string `json:"directions"`
	Notes       string   `json:"notes"`
}
