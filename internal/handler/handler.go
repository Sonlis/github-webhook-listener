package handler

import (
	"encoding/json"
	"github.com/Sonlis/github-webhook-listener/internal/applyChanges"
	"github.com/Sonlis/github-webhook-listener/internal/checkSignature"
	"github.com/Sonlis/github-webhook-listener/internal/config"
	"github.com/Sonlis/github-webhook-listener/internal/repoSynchronizer"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Reference struct {
	Ref string `json:"ref"`
}

func HandleRequestPrivate(c *gin.Context) {
	configuration := config.NewConfig()
	var reference Reference
	hook := new(checkSignature.Hook)
	secret := []byte(configuration.GitHookSecret)
	hook, err := checkSignature.Parse(secret, c.Request)
	if err != nil {
		log.Println("Wrong signature")
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong signature"})
		return
	}

	err = json.Unmarshal(hook.Payload, &reference)

	if reference.Ref != "refs/heads/master" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong branch"})
		log.Println("Wrong branch")
		return
	}

	c.Writer.WriteHeader(200)

	if err = repoSynchronizer.PullRepo(configuration); err != nil {
		log.Printf("Error pulling the repository: %v", err)
	}

	if err = applyChanges.ApplyChanges(); err != nil {
		log.Printf("Error applying changes: %v", err)
	}

}

func HandleRequestPublic(c *gin.Context) {
	configuration := config.NewConfig()

	if err = repoSynchronizer.PullRepo(configuration); err != nil {
		log.Printf("Error pulling the repository: %v", err)
	}
	if err = applyChanges.ApplyChanges(); err != nil {
		log.Printf("Error applying changes: %v", err)
	}
	c.Writer.WriteHeader(200)
}
