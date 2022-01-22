package repoSynchronizer


import (
	"github.com/go-git/go-git/v5"
	"os"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)


func PullRepo(path string) (err error) {

	r, err := git.PlainOpen(path)
	if err != nil {
		return err 
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}
	
	if err = w.Pull(&git.PullOptions{RemoteName: "origin",
	Auth: &http.BasicAuth{
		Username: "Sonlis", // yes, this can be anything except an empty string
		Password: os.Getenv("GITHUB_TOKEN"),
	}}); err != nil {
		return err 
	}

	return nil 
}