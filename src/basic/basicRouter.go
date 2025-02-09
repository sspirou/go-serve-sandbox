package trybasicrouter

import (
	"fmt"
	"net/http"
)

func StartBasicServer() {
	message := "Go Server Sandbox Listening on Port 8083"
	fmt.Println(message)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my Sandbox Go Server")
	})

	fs := http.FileServer(http.Dir("./serve-me/"))
	http.Handle("/serve-me/", http.StripPrefix("/serve-me/", fs))

	http.ListenAndServe(":8083", nil)
}
