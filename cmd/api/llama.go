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
	Timeout: 20 * time.Second,
}
var endpoint string = "http://localhost:11434/api/generate"

var requestBuffer bytes.Buffer

func Prompt(request api.GenerateRecipesRequest) api.GenerateRecipesResponse {

	var error bool = false

	// read in init prompt
	initPromptBytes, err := os.ReadFile("init.txt")
	if err != nil {
		fmt.Print(err)
		error = true
	}
	initPrompt := string(initPromptBytes)

	// read prompt json from scrounge
	userPrompt, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		error = true
	}

	// form request to ollama
	var llamaReq = api.LlamaRequest{
		Model:  "llama3.1",
		Prompt: initPrompt + "\n" + string(userPrompt),
		Format: "json",
		Stream: false,
	}

	err = json.NewEncoder(&requestBuffer).Encode(llamaReq)
	if err != nil {
		fmt.Println("Error when encoding request")
		error = true
	}

	// POST
	httpResp, err := client.Post(endpoint, "json", &requestBuffer)
	if err != nil {
		fmt.Println("Error when posting LlamaRequest")
		error = true
	}
	requestBuffer.Reset()
	defer httpResp.Body.Close()

	// Decode response data
	var llamaResponse api.LlamaResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&llamaResponse); err != nil {
		fmt.Println("Error when reading LlamaResponse", err)
		error = true
	}

	var recipesResponse api.GenerateRecipesResponse
	if err := json.NewDecoder(strings.NewReader(llamaResponse.Response)).Decode(&recipesResponse); err != nil {
		fmt.Println("Error reading Recipes from LlamaResponse", err)
		error = true
	}

	if error {
		return api.GenerateRecipesResponse{}
	} else {
		return recipesResponse
	}
}
