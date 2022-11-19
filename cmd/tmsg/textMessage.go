package tmsg

import (
	"bytes"
	"fmt"
	"io"
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

var TwilioCmd = &cobra.Command{
	Use:   "twilio",
	Short: "notify threw text message",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		//start timer
		start := time.Now()

		//creating twilio provider
		twilio := msg.Msg{
			From: from,
			To:   to,
		}

		//slplit the command and run it
		commandArgs := strings.Split(command, " ")
		c := exec.Command(commandArgs[0], commandArgs[1:]...)

		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)
		c.Stdout = mw
		c.Stderr = mw

		// Execute the command
		if err := c.Run(); err != nil {
			twilio.Msg = "Subject: Error \n" + err.Error()
		} else {
			twilio.Msg = "Subject: Success \n" + stdBuffer.String()
		}

		log.Println(stdBuffer.String())

		//send to my phone when finish
		if err := twilio.SendMessage(start); err != nil {
			log.Fatal(err)
		}

		//print result of the command
		fmt.Println(string(twilio.Msg))
	},
}

func init() {
	TwilioCmd.Flags().StringVarP(&from, "from", "f", "+13023039351", "The phone from")
	TwilioCmd.Flags().StringVarP(&command, "command", "c", "", "Command to run")
	TwilioCmd.Flags().StringArrayVarP(&to, "to", "t", []string{"+<number>"}, "Who to notify")
}
