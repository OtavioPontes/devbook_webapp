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

/* func init() {
	hashKey := securecookie.GenerateRandomKey(16)

	blockKey := securecookie.GenerateRandomKey(16)

	fmt.Println(hex.EncodeToString(hashKey))
	fmt.Println(hex.EncodeToString(blockKey))
} */

func main() {
	config.Load()
	cookies.Config()
	r := router.Generate()
	utils.LoadTemplates()

	fmt.Printf("Running WebApp in %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
