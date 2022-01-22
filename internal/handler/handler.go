package handler 

import (
	"github.com/gin-gonic/gin"
	"github.com/Sonlis/github-webhook-listener/internal/checkSignature"
	"github.com/Sonlis/github-webhook-listener/internal/repoSynchronizer"
	"github.com/Sonlis/github-webhook-listener/internal/config"
	"github.com/Sonlis/github-webhook-listener/internal/applyChanges"
	"log"
	"os"
	"net/http"
	"encoding/json"
)

func HandleRequest(c *gin.Context) {
	configuration := config.NewConfig()
	var reference Reference
	hook := new(checkSignature.Hook)
	secret := []byte(os.Getenv("GITHUB_SIGNATURE"))
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

	if err = repoSynchronizer.PullRepo(configuration.GitPath); err != nil {
		log.Printf("Error pulling the repository: %v", err)
	}

	if err = applyChanges.ApplyChanges(); err != nil {
		log.Printf("Error applying changes: %v", err)
	}





}

type Reference struct {
	Ref string `json:"ref"`
}

