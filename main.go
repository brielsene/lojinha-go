package main

import (
	"gabriel/routes"
	"net/http"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
