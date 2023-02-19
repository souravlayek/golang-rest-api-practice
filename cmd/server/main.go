package main

import (
	"log"
	"net/http"

	"github.com/souravlayek/rest-api-tutorial/internal/api/router"
	"github.com/souravlayek/rest-api-tutorial/internal/database"
	"github.com/souravlayek/rest-api-tutorial/utils"
)

func main() {
	utils.LoadENV()
	database.ConnectDB()
	router := router.CreateRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
