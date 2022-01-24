package applyChanges

import (
	"os/exec"
)

func ApplyChanges() (err error) {

	cmd := exec.Command("sudo", "systemctl", "restart", "docker-env")
	err = cmd.Run()
	return err
}
