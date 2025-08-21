package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"scrounge-engine/api"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: 30 * time.Second,
}

var endpoint string = "http://localhost:11434/v1/chat/completions"

func Prompt(request api.GenerateRecipesRequest) (api.GenerateRecipesResponse, error) {

	var requestBuffer bytes.Buffer

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// read in init prompt
	initPromptBytes, err := os.ReadFile("init.txt")
	if err != nil {
		return api.GenerateRecipesResponse{}, err
	}
	initPrompt := string(initPromptBytes)

	// read prompt json from scrounge
	userPrompt, err := json.Marshal(request)
	if err != nil {
		return api.GenerateRecipesResponse{}, err
	}

	// form request to llm
	var llmRequest = api.LLMRequest{
		Model: "google/gemma-3-12b",
		Messages: []api.LLMMessage{
			{
				Role:    "system",
				Content: string(initPrompt),
			},
			{
				Role:    "user",
				Content: string(userPrompt),
			},
		},
		Stream: false,
	}

	err = json.NewEncoder(&requestBuffer).Encode(llmRequest)
	if err != nil {
		fmt.Println("Error when encoding request")
		return api.GenerateRecipesResponse{}, err
	}

	// POST
	httpResp, err := client.Post(endpoint, "application/json", &requestBuffer)
	if err != nil {
		fmt.Println("Error when posting LLMRequest")
		return api.GenerateRecipesResponse{}, err
	}
	responseBody := httpResp.Body
	defer httpResp.Body.Close()

	// Decode response data
	var llamaResponse api.LLMResponse
	if err := json.NewDecoder(responseBody).Decode(&llamaResponse); err != nil {
		fmt.Println("Error when reading LLMResponse", err)
		return api.GenerateRecipesResponse{}, err
	}

	recipeJsonString := TrimRecipeJSON(llamaResponse.Choices[0].Message.Content)

	var recipesResponse api.GenerateRecipesResponse
	if llamaResponse.Choices[0].Message.Content != "" {
		if err := json.Unmarshal([]byte(recipeJsonString), &recipesResponse); err != nil {
			fmt.Println("Error reading Recipes from LLMResponse:", err)
			return api.GenerateRecipesResponse{}, err
		}
	} else {
		fmt.Println("LLM Response is empty or unexpected format")
		return api.GenerateRecipesResponse{}, err
	}

	return recipesResponse, nil
}

func TrimRecipeJSON(in string) string {
    replacer := strings.NewReplacer("```json\n", "", "\n```", "")
    return replacer.Replace(in)
}
