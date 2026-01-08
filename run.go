package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func runProcesses() {
	backend := exec.Command("bun", "run", "dev")
	backend.Dir = "./server"
	backend.Stdout = os.Stdout
	backend.Stderr = os.Stderr

	if err := backend.Start(); err != nil {
		log.Fatal("Gagal jalanin backend:", err)
	}

	client := exec.Command("bun", "run", "dev")
	client.Dir = "./client"
	client.Stdout = os.Stdout
	client.Stderr = os.Stderr

	if err := client.Start(); err != nil {
		log.Fatal("Gagal jalanin client:", err)
	}

	log.Println("Backend & Client jalan. Ctrl+C buat matiin.")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Stopping processes...")
	_ = backend.Process.Kill()
	_ = client.Process.Kill()
}

func main() {
	runProcesses()
}
