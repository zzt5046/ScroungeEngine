package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"scrounge-engine/api"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}
var endpoint string = "http://localhost:11434/api/generate"
var model string = "llama3.1"
var format string = "json"
var useStream bool = false

var requestBuffer bytes.Buffer

func InitLlama() {
	fmt.Println("sending init prompt to Llama...")
	Prompt(api.GenerateRecipeRequest{}, true)
}

func Prompt(request api.GenerateRecipeRequest, init bool) api.LlamaResponse {

	var llamaReq api.LlamaRequest
	if init {
		// read in init prompt
		bytes, err := os.ReadFile("init.txt")
		if err != nil {
			fmt.Print(err)
		}

		initPrompt := string(bytes) // convert content to a 'string'

		llamaReq = api.LlamaRequest{
			Model:  model,
			Prompt: initPrompt,
			Format: format,
			Stream: useStream,
		}
	} else {
		llamaReq = convertToLlama(request)
	}

	err := json.NewEncoder(&requestBuffer).Encode(llamaReq)
	if err != nil {
		fmt.Println("Error when encoding request")
	}

	httpResp, err := client.Post(endpoint, "json", &requestBuffer)
	if err != nil {
		fmt.Println("error doing llama request: ", err)
	}
	requestBuffer.Reset()
	defer httpResp.Body.Close()

	var response api.LlamaResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&response); err != nil {
		fmt.Println("Error reading json response: ", err)
	}
	return response

}

func convertToLlama(request api.GenerateRecipeRequest) api.LlamaRequest {

	promptJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}

	return api.LlamaRequest{
		Model:  model,
		Prompt: string(promptJSON),
		Format: format,
		Stream: useStream,
	}
}
