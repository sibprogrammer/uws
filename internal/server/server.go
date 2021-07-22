package server

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

const defaultIp = "127.0.0.1"
const defaultPort = "8080"

func Run() error {
	ip := defaultIp
	port := defaultPort

	if fileExists("composer.json") {
		return runPHP(ip, port)
	}

	if fileExists("Gemfile") && fileExists("config.ru") {
		return runRuby(ip, port)
	}

	return runStatic(ip, port)
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func runCommand(name string, arg ...string) error {
	command := exec.Command(name, arg...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func runStatic(ip string, port string) error {
	address := fmt.Sprintf("http://%s:%s", ip, port)
	fmt.Println("Static server started:", address)
	return http.ListenAndServe(":"+port, http.FileServer(http.Dir(".")))
}

func runPHP(ip string, port string) error {
	address := fmt.Sprintf("%s:%s", ip, port)
	args := []string{"-S", address}

	if fileExists("public") {
		args = append(args, "-t", "public")
	} else if fileExists("web") {
		args = append(args, "-t", "web")
	}

	return runCommand("php", args...)
}

func runRuby(ip string, port string) error {
	return runCommand("bundle", "exec", "rackup", "--host", ip, "--port", port)
}
