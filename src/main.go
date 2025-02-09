package main

import (
	"fmt"
	"net/http"
)

func main() {
	message := "Diagrammer Backend Service Started on Port 8083"
	fmt.Println(message)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Diagrammer Backend Service Listening on Port 8083")
	})

	fs := http.FileServer(http.Dir("serve-me/"))
	http.Handle("/serve-me/", http.StripPrefix("/serve-me/", fs))

	http.ListenAndServe(":8083", nil)
}
