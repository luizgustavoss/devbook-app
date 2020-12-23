package main

import (
	"devbookapp/src/config"
	"devbookapp/src/router"
	"devbookapp/src/security"
	"devbookapp/src/utils"
	"fmt"
	"log"
	"net/http"
)


func init(){
	utils.LoadTemplates()
}


func main() {

	config.Load()
	security.ConfigureSecureCookie()

	router := router.GetRouter()

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port), router))

	fmt.Sprintf("Listening on port %d", config.Port)
}
