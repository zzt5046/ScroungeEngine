package main

import (
	"fmt"
	"net/http"
	"scrounge-engine/api"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting ScroungeEngine...")
	initRouter()
}

func initRouter() {
	fmt.Println("Starting gin...")
	router := gin.Default()
	router.POST("/generate-recipe", generateRecipe)
	router.Run("localhost:8085")
}

// API Methods ------------------------------------------------

func generateRecipe(context *gin.Context) {

	var request api.GenerateRecipesRequest
	if err := context.BindJSON(&request); err != nil {
		fmt.Println(err)
		fmt.Println("error serializing request")
	}

	var response api.GenerateRecipesResponse
	var code int
	recipes := Prompt(request).Recipes
	if recipes != nil {
		response = api.GenerateRecipesResponse{
			Recipes: recipes,
		}
		code = http.StatusOK
	} else {
		response = api.GenerateRecipesResponse{}
		code = http.StatusInternalServerError
	}

	context.IndentedJSON(code, response)
}
