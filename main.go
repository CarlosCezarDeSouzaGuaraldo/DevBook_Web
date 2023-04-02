package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
	"web/src/utils"
)

func main() {
	utils.LoadTemplate()
	r := router.Generate()

	fmt.Printf("WEB listening on PORT: %d\n", 3000)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
