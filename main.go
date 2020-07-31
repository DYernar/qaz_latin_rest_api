package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	controller "qaz_latin/controllers"

	"github.com/rs/cors"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("port not found")
		fmt.Println("port not found")
		fmt.Println("port not found")
		port = "4444"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})
	mux.HandleFunc("/signin", controller.Signin)

	mux.HandleFunc("/all", controller.GetData)
	mux.HandleFunc("/score", controller.UpdateScore)

	handler := cors.Default().Handler(mux)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
