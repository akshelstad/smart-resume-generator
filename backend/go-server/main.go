package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akshelstad/smart-resume-generator/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../python-ai/.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	r := gin.Default()

	r.POST("/generate-resume", handlers.GenerateResume)
	// r.POST("/generate-cover-letter", GenerateCoverLetter)

	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Go Server running on port: ", port)
	r.Run(":" + port)
}
