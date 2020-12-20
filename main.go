package main

import (
	"devbookapp/src/router"
	"devbookapp/src/utils"
	"fmt"
	"log"
	"net/http"
)


func init(){
	utils.LoadTemplates()
}


func main() {

	fmt.Println("DevBook Web App Init")

	router := router.GetRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
