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

	router := gin.Default()
	router.GET("/recipe", generateRecipe)

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

	router.Run("localhost:8085")
}

var recipe = api.Recipe{
	Name:        "Go Test Recipe",
	Description: "This is a recipe created when starting Scrounge Engine's API.",
	Ingredients: []string{"1 tsp flour", "1 lb butter", "1 c patience"},
}

func generateRecipe(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, recipe)
}
