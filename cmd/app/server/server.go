package server

import (
	"app/cmd/app/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func StartServer() {
	readEnv()
	http.HandleFunc("/callback", handler.LineHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func readEnv() {
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "dev")
	}

	err := godotenv.Load(fmt.Sprintf("./.env.%s", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
