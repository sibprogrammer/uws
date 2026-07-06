package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

var killallCmd = &cobra.Command{
	Use:   "killall",
	Short: "Stop all the running servers",
	Run: func(cmd *cobra.Command, args []string) {
		servers := getServers()
		killed := 0

		for _, server := range servers {
			pid, err := strconv.Atoi(server)
			if err != nil {
				fmt.Printf("Invalid PID %q: %v\n", server, err)
				continue
			}

			process, err := os.FindProcess(pid)
			if err != nil {
				continue
			}

			if err := process.Signal(syscall.SIGTERM); err != nil {
				if errors.Is(err, os.ErrProcessDone) {
					removeStalePidFile(pid)
				}
				continue
			}

			killed++
		}

		fmt.Println("Total servers stopped:", killed)
	},
}

func removeStalePidFile(pid int) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	_ = os.Remove(fmt.Sprintf("%s/.uws.%d", userHomeDir, pid))
}
