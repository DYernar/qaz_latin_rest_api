package main

import (
	"log"
	"net/http"
	"os"
	controller "qaz_latin/controllers"

	"github.com/rs/cors"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
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
	mux.HandleFunc("/getAll", controller.GetAllUserScore)
	mux.HandleFunc("/getNews", controller.GetNews)
	mux.HandleFunc("/addNews", controller.AddNews)
	mux.HandleFunc("/getScores", controller.GetScores)
	mux.HandleFunc("/getGameScore", controller.GetScore)
	mux.HandleFunc("/saveScore", controller.SaveScore)
	mux.HandleFunc("/drop", controller.DropTables)
	handler := cors.Default().Handler(mux)

	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal("Listen and serve err: ", err)
	}
}
