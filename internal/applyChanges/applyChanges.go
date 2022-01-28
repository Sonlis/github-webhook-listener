package applyChanges

import (
	"log"
	"net/http"
	"os/exec"
)

func ApplyChanges() (err error) {
	_, err = http.Get("http://192.168.0.150:9292")
	if err != nil {
		log.Println(err)
	}
	cmd := exec.Command("sudo", "systemctl", "restart", "docker-env")
	err = cmd.Run()
	return err
}
