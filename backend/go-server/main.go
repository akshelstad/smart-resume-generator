package main

import (
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type config struct {
	fileserverHits atomic.Int32
	platform       string
	apiKey         string
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY must be set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}
	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8080"
	}
	filepathRoot := os.Getenv("FILEPATH_ROOT")
	if filepathRoot == "" {
		log.Fatal("FILEPATH_ROOT must be set")
	}

	config := config{
		fileserverHits: atomic.Int32{},
		platform:       platform,
		apiKey:         apiKey,
	}

	mux := mux.NewRouter()
	// mux.Handle("/app/", config.middleWareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))

	mux.HandleFunc("/api/healthz", handlerReadiness).Methods("GET")

	mux.HandleFunc("/generate-resume", config.handlerGenerateResume).Methods("POST")

	mux.HandleFunc("/admin/metrics", config.handlerMetrics).Methods("GET")
	mux.HandleFunc("/admin/reset", config.handlerReset).Methods("POST")

	serverWithMiddleware := config.middleWareMetricsInc(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: serverWithMiddleware,
	}

	log.Println("Go Server running on port: ", port)
	log.Fatal(srv.ListenAndServe())

}
