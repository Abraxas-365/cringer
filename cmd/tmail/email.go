package tmail

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Abraxas-365/cringer/pkg/mail"
	"github.com/spf13/cobra"
)

var from string
var command string
var to []string
var password string

var EmailCmd = &cobra.Command{
	Use:   "email",
	Short: "notify threw mail",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()
		email := &mail.Mail{
			From:     from,
			To:       to,
			Password: password,
		}
		commandArgs := strings.Split(command, " ")
		c := exec.Command(commandArgs[0], commandArgs[1:]...)

		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)
		c.Stdout = mw
		c.Stderr = mw

		// Execute the command
		if err := c.Run(); err != nil {
			email.Msg = "Subject: Error \n" + err.Error()
		} else {
			email.Msg = "Subject: Success \n" + stdBuffer.String()
		}

		log.Println(stdBuffer.String())

		if err := email.SendMail(start); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {

	EmailCmd.Flags().StringVarP(&from, "from", "f", "jhon@gmail.com", "The phone from")
	EmailCmd.Flags().StringVarP(&command, "command", "c", "", "Command to run")
	EmailCmd.Flags().StringArrayVarP(&to, "to", "t", []string{"jhon@gmail.com"}, "Who to notify")
	EmailCmd.Flags().StringVarP(&password, "password", "p", "password", "password")
}
