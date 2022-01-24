package main

import (
	"github.com/Sonlis/github-webhook-listener/internal/handler"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	if _, exists := os.LookupEnv("GITHUB_PATH"); !exists {
		log.Fatal("Env variable GITHUB_PATH not set")
	} else if _, exists = os.LookupEnv("GITHUB_TOKEN"); !exists {
		log.Fatal("Env variable GITHUB_TOKEN not set")
	} else if _, exists = os.LookupEnv("GITHUB_HOOK_SECRET"); !exists {
		log.Fatal("Env variable GITHUB_HOOK_SECRET not set")
	}

	router := gin.Default()
	router.POST("/github-webhook", handler.HandleRequest)
	router.Run(":9292")
}
