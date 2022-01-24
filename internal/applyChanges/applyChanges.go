package applyChanges

import (
	"os/exec"
)

func ApplyChanges() (err error) {

	cmd := exec.Command("/bin/sh", "-c", "sudo", "systemd", "restart", "docker-env")
	err = cmd.Run()
	return err
}
