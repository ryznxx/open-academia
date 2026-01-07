package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func check(name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s: NOT FOUND\n", strings.Title(name))
		return false
	}
	fmt.Printf("%s: %s", strings.Title(name), out)
	return true
}

func RequiredRuntime() bool {
	fmt.Print("\n")
	node := check("node", "-v")
	bun := check("bun", "--version")
	deno := check("deno", "--version")

	if node || bun || deno {
		return true
	}
	return false
}

func RequiredTools() bool {
	docker := check("docker", "-v")

	if docker {
		return true
	}

	return false
}
