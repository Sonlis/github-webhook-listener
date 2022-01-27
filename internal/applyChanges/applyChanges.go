package applyChanges

import (
	"os/exec"
	"net/http"
	"log"
)

func ApplyChanges() (err error) {
	resp, err := http.Get("http://192.168.0.150:9292")
	if err != nil {
		log.Println(err)
	}
	cmd := exec.Command("sudo", "systemctl", "restart", "docker-env")
	err = cmd.Run()
	return err
}
