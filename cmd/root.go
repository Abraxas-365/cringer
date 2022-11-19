package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Abraxas-365/cringer/pkg/msg"
	"github.com/spf13/cobra"
)

var from string
var command string
var to []string

var rootCmd = &cobra.Command{
	Use:   "cringer",
	Short: "notify when command finish execution",
	Long:  `notify when command finish execution.`,
	Run: func(cmd *cobra.Command, args []string) {

		//start timer
		start := time.Now()

		//slplit the command and run it
		commandArgs := strings.Split(command, " ")
		c := exec.Command(commandArgs[0], commandArgs[1:]...)
		stdout, err := c.Output()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		//creating twilio provider
		twilio := msg.Msg{
			From: from,
			To:   to,
		}

		//send to my phone when finish
		if err := twilio.SendMessage(string(stdout), start); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		//print result of the command
		fmt.Println(string(stdout))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&from, "from", "f", "+13023039351", "The phone from")
	rootCmd.Flags().StringVarP(&command, "command", "c", "", "Command to run")
	rootCmd.Flags().StringArrayVarP(&to, "to", "t", []string{"+51984266436"}, "Who to notify")
}
