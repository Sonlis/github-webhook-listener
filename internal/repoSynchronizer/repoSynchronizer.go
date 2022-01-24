package repoSynchronizer

import (
	"github.com/Sonlis/github-webhook-listener/internal/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func PullRepo(config config.Config) (err error) {

	auth := &http.BasicAuth{
		Username: config.GitUsername,
		Password: config.GitToken,
	}
	r, err := git.PlainOpen(config.GitPath)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	if err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth: auth,
		}); err != nil {
		return err
	}

	return nil
}
