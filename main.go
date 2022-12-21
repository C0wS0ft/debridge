package main

import (
	"debridge/bridgelib"
	"log"
	"net/http"
)

func main() {
	bridgelib.Init()

	httpsMux := http.NewServeMux()
	httpsMux.HandleFunc("/estimate", bridgelib.Estimate)
	httpsMux.HandleFunc("/transaction", bridgelib.Transaction)

	err := http.ListenAndServe(":7900", httpsMux)

	if err != nil {
		log.Printf("Error staring DeBridge server: %s", err)
	}
}
