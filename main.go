
package main

import (
	"net/http"

	"github.com/0xlax/go-rest-api/database"
	"github.com/0xlax/go-rest-api/routes"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	database.SetupMongo()
	routes.SetUp(r)
	http.ListenAndServe("localhost:3000", r)
}