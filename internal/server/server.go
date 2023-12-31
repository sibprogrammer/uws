package server

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func Run(port string, ip string) error {
	if fileExists("composer.json") {
		return runPHP(ip, port)
	}

	if fileExists("Gemfile") && fileExists("config.ru") {
		return runRuby(ip, port)
	}

	if fileExists("package.json") {
		return runJS(ip, port)
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

func runCommand(command *exec.Cmd) error {
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func runStatic(ip string, port string) error {
	address := fmt.Sprintf("http://%s:%s", ip, port)
	workDir, _ := os.Getwd()
	fmt.Printf("Static server for %s started: %s\n", workDir, address)
	return http.ListenAndServe(":"+port, http.FileServer(http.Dir(workDir)))
}

func runPHP(ip string, port string) error {
	address := fmt.Sprintf("%s:%s", ip, port)
	args := []string{"-S", address}

	if fileExists("public") {
		args = append(args, "-t", "public")
	} else if fileExists("web") {
		args = append(args, "-t", "web")
	}

	return runCommand(exec.Command("php", args...))
}

func runRuby(ip string, port string) error {
	return runCommand(exec.Command("bundle", "exec", "rackup", "--host", ip, "--port", port))
}

func runJS(ip string, port string) error {
	var command *exec.Cmd

	if fileExists("yarn.lock") {
		command = exec.Command("yarn", "start")
	} else {
		command = exec.Command("npm", "start")
	}

	command.Env = append(os.Environ(), "PORT="+port)

	return runCommand(command)
}
