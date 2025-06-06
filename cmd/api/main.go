package main

import (
	"fmt"
	"net/http"
	"scrounge-engine/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetReportCaller(true)

	fmt.Println("Starting gin...")
	router := gin.Default()
	router.POST("/recipe", generateRecipe)
	router.POST("/recipes", generateRecipes)
	router.Run("localhost:8085")

	printSplash()
}

func generateRecipe(context *gin.Context) {

	var request api.GenerateRecipeRequest
	if err := context.BindJSON(&request); err != nil {
		return
	}

	context.IndentedJSON(http.StatusOK, NewRecipe(request))
}

func generateRecipes(context *gin.Context) {

}

func printSplash(){
	fmt.Println(`
________  ________  ________  ________  ___  ___  ________   ________  _______           _______   ________   ________  ___  ________   _______      
|\   ____\|\   ____\|\   __  \|\   __  \|\  \|\  \|\   ___  \|\   ____\|\  ___ \         |\  ___ \ |\   ___  \|\   ____\|\  \|\   ___  \|\  ___ \     
\ \  \___|\ \  \___|\ \  \|\  \ \  \|\  \ \  \\\  \ \  \\ \  \ \  \___|\ \   __/|        \ \   __/|\ \  \\ \  \ \  \___|\ \  \ \  \\ \  \ \   __/|    
 \ \_____  \ \  \    \ \   _  _\ \  \\\  \ \  \\\  \ \  \\ \  \ \  \  __\ \  \_|/__       \ \  \_|/_\ \  \\ \  \ \  \  __\ \  \ \  \\ \  \ \  \_|/__  
  \|____|\  \ \  \____\ \  \\  \\ \  \\\  \ \  \\\  \ \  \\ \  \ \  \|\  \ \  \_|\ \       \ \  \_|\ \ \  \\ \  \ \  \|\  \ \  \ \  \\ \  \ \  \_|\ \ 
    ____\_\  \ \_______\ \__\\ _\\ \_______\ \_______\ \__\\ \__\ \_______\ \_______\       \ \_______\ \__\\ \__\ \_______\ \__\ \__\\ \__\ \_______\
   |\_________\|_______|\|__|\|__|\|_______|\|_______|\|__| \|__|\|_______|\|_______|        \|_______|\|__| \|__|\|_______|\|__|\|__| \|__|\|_______|
   \|_________|                                                                                                                                       
   
	`)
}
