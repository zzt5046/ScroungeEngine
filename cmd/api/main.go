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

	var code int
	recipesResponse, err := Prompt(request)
	if recipesResponse.Recipes != nil {
		code = http.StatusOK
	} else if err != nil {
		//(BLACK EYED PEAS) OH SHIT!
		code = http.StatusInternalServerError
		fmt.Println(err)
	}

	context.IndentedJSON(code, recipesResponse)
}
