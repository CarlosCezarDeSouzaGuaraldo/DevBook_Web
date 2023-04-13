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

	fmt.Printf(fmt.Sprintf("WEB application running on %s:%d\n", config.HOST, config.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.HOST, config.PORT), r))
}
