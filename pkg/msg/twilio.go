package msg

import (
	"strings"
	"time"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

// "+13023039351"
type Msg struct {
	From string
	To   []string
}

func (m *Msg) SendMessage(msg string, duration time.Time) error {

	elapsed := time.Since(duration)

	client := twilio.NewRestClient()
	params := &api.CreateMessageParams{}
	params.SetBody(msg + " " + "Duration" + shortDur(elapsed))
	params.SetFrom(m.From)
	for _, number := range m.To {
		params.SetTo(number)
		if _, err := client.Api.CreateMessage(params); err != nil {
			return err
		}
	}

	return nil
}

func shortDur(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}
