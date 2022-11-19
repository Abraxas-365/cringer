package msg

import (
	"time"

	"github.com/Abraxas-365/cringer/internal"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type Msg struct {
	From string
	To   []string
	Msg  string
}

func (m *Msg) SendMessage(duration time.Time) error {

	elapsed := time.Since(duration)

	client := twilio.NewRestClient()
	params := &api.CreateMessageParams{}
	params.SetBody(m.Msg + " " + "Duration: " + internal.ShortDur(elapsed))
	params.SetFrom(m.From)
	for _, number := range m.To {
		params.SetTo(number)
		if _, err := client.Api.CreateMessage(params); err != nil {
			return err
		}
	}

	return nil
}
