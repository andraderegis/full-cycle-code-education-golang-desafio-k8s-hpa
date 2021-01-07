package main

import (
	"fmt"
	"golang-fullcycle-desafio-k8s-hpa/arithmetics"
	"net/http"
)

func main() {
	fmt.Println("Server running at port 8080")
	http.HandleFunc("/", handlerRoot)
	http.ListenAndServe(":8080", nil)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	arithmetics.SumSqrtLoop(10000000, 4)

	fmt.Fprint(w, "Code Education Rocks!")
}
