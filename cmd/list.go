package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show the list of running servers",
	Run: func(cmd *cobra.Command, args []string) {
		servers := getServers()
		for _, server := range servers {
			fmt.Println("Server with PID:", server)
		}
		fmt.Println("Total servers:", len(servers))
	},
}

func getServers() []string {
	var servers []string

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to obtain user home directory: %v\n", err)
	}

	matches, err := filepath.Glob(fmt.Sprintf("%s/.uws.*", userHomeDir))

	re := regexp.MustCompile(`[\d]+$`)
	for _, match := range matches {
		pid := re.FindAllString(match, 1)
		servers = append(servers, pid[0])
	}

	return servers
}
