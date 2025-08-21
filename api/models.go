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

// REQUESTS TO LLM----------------------------------------------------

type LLMMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LLMRequest struct {
	Model    string       `json:"model"`
	Messages []LLMMessage `json:"messages"`
	Stream   bool         `json:"stream"`
}

// RESPONSES FROM LLM-------------------------------------------------

type LLMChoice struct {
	Message LLMMessage `json:"message"`
}

// General response object from llm api
type LLMResponse struct {
	Choices []LLMChoice `json:"choices"`
}

// Structure of generated recipe ("Response" from LLMResponse)
type Recipe struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Directions  []string `json:"directions"`
	Notes       string   `json:"notes"`
}
