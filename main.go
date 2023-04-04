package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/config"
	"web/src/cookies"
	"web/src/router"
	"web/src/utils"
)

func main() {
	config.Load()
	cookies.Configure()
	utils.LoadTemplate()
	r := router.Generate()

	fmt.Printf("WEB listening on PORT: %d\n", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.HOST, config.PORT), r))
}
