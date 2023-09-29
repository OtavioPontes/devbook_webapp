package main

import (
	"devbook_webapp/src/config"
	"devbook_webapp/src/cookies"
	"devbook_webapp/src/router"
	"devbook_webapp/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	cookies.Config()
	r := router.Generate()
	utils.LoadTemplates()

	fmt.Println("Running WebApp in 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
