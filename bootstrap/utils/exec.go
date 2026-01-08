package utils

import (
	"log"
	"os"
	"os/exec"
)

func Run(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunCmd(cmdStr string, args ...string) {
	cmd := exec.Command(cmdStr, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Gagal jalanin %s %v: %v", cmdStr, args, err)
	}
}
