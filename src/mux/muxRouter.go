package trymux

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartMuxServer() {
	message := "Mux Router Listening on Port 8083"
	fmt.Println(message)

	router := mux.NewRouter()

	router.HandleFunc("/mux", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the mux router")
	})

	router.HandleFunc("/serve-me/digitsText/{digit}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileName := vars["digit"]
		fmt.Fprintf(w, "Digit file selected: %s\n", fileName)
	})

	router.HandleFunc("/serve-me/movie", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./serve-me/panda_short.mov")
	})

	http.ListenAndServe(":8083", router)
}
