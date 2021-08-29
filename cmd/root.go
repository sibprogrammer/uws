package cmd

import (
	"fmt"
	"github.com/pkg/browser"
	"github.com/sevlyar/go-daemon"
	"github.com/sibprogrammer/uws/internal/server"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "uws",
	Short: "Universal Web Server for development purposes",
	RunE: func(cmd *cobra.Command, args []string) error {
		daemonMode, _ := cmd.Flags().GetBool("daemon")

		done := make(chan os.Signal, 1)
		signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

		if daemonMode {
			workDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
			context := getDaemonContext("uws", workDir)

			d, err := context.Reborn()
			if err != nil {
				log.Fatal("Unable to run: ", err)
			}
			if d != nil {
				fmt.Println("The daemon has been started in the background mode.")
				return nil
			}
			defer context.Release()
		}

		port, _ := cmd.Flags().GetInt("port")
		ip, _ := cmd.Flags().GetIP("binding")

		var err error
		go func() {
			createPidFile()
			err = server.Run(strconv.Itoa(port), ip.String())
			if err != nil {
				fmt.Printf("Failed to launch the server: %v\n", err)
			}

		}()

		go func() {
			time.Sleep(time.Second)
			url := fmt.Sprintf("http://%s:%s", ip.String(), strconv.Itoa(port))
			err = browser.OpenURL(url)
			if err != nil {
				fmt.Printf("Unable to open the browser: %v\n", err)
			}
		}()

		<-done
		releasePidFile()

		return err
	},
}

func Execute() {
	rootCmd.AddCommand(
		versionCmd,
		listCmd,
	)

	rootCmd.PersistentFlags().BoolP("daemon", "d", false, "Run in the background mode")
	rootCmd.PersistentFlags().IntP("port", "p", 8080, "Run the server on the specified port")
	rootCmd.PersistentFlags().IPP("binding", "b", net.IPv4(127,0,0,1), "Bind the server to the specified IP")

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getDaemonContext(binaryName string, workDir string) *daemon.Context {
	return &daemon.Context{
		PidFileName: path.Join(workDir, binaryName+".pid"),
		PidFilePerm: 0644,
		LogFileName: path.Join(workDir, binaryName+".log"),
		LogFilePerm: 0640,
		WorkDir:     workDir,
		Umask:       027,
	}
}

func getPidFileName() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to obatin user home directory: %v\n", err)
	}
	return fmt.Sprintf("%s/.uws.%d", userHomeDir, os.Getpid())
}

func createPidFile() {
	err := ioutil.WriteFile(getPidFileName(), []byte(""), 0664)
	if err != nil {
		fmt.Printf("Unable to create PID file: %v\n", err)
	}
}

func releasePidFile() {
	err := os.Remove(getPidFileName())
	if err != nil {
		fmt.Printf("Unable to delete PID file: %v\n", err)
	}
}
