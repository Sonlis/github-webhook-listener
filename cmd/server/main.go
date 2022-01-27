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
	} else if _, exists = os.LookupEnv("GITHUB_USERNAME"); !exists {
		log.Fatal("Env varaible GITHUB_USERNAME not set")
	}

	router := gin.Default()
	if os.Args[0] == "public" {
		router.GET("/github-webhook", handler.HandleRequestPublic)
	} else if os.Args[0] == "private" {
		router.POST("/github-webhook", handler.HandleRequestPrivate)
	}
	router.Run(":9292")
}
