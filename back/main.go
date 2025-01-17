package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ramada/api/src/cmd"
	"ramada/api/src/config"
	"ramada/api/src/router"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

func main() {
	config.Load()

	if len(os.Args) > 1 {
		cmd.Flow()

		return
	}

	apiPort := fmt.Sprintf(":%d", config.API_PORT)
	aRouter := router.GetRouter()
	log.Fatal(http.ListenAndServe(apiPort, aRouter))

	fmt.Println(`Iniciando com GO`)
}
