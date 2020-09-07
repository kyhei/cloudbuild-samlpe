package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mymod/utils"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		j := json.NewDecoder(r.Body)
		var req struct {
			A int `json:"a"`
			B int `json:"b"`
		}

		j.Decode(&req)
		fmt.Fprintf(w, "%d\n", utils.Add(req.A, req.B))
	})

	log.Println("start server...")
	http.ListenAndServe(":8080", mux)
}
