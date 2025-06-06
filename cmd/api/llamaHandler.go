package main

import (
	"fmt"
	"net/http"
	"scrounge-engine/api"
	"bytes"
	"encoding/json"
	"time"
)

var endpoint string = "http://localhost:11434/api/generate"

func NewRecipe(request api.GenerateRecipeRequest) api.LlamaResponse {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	var requestBuffer bytes.Buffer
    err := json.NewEncoder(&requestBuffer).Encode(request)
    if err != nil {
        fmt.Println("Error when encoding request")
    }

	httpResp, err := client.Post(endpoint, "json", &requestBuffer)
	if err != nil {
		fmt.Println("error doing llama request: ", err)
	}
	defer httpResp.Body.Close()

	var response api.LlamaResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&response); err != nil {
		fmt.Println("Error reading json response: ", err)
	}
	return response
	
}