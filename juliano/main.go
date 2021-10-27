package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"soder_loja/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
