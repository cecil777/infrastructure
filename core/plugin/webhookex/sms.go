package webhookex

import (
	"encoding/json"

	isms "github.com/cecil777/infrastructure/core/plugin/sms"
	"github.com/cecil777/infrastructure/core/runtimeex"
)

type sms struct {
	url     string
	webhook runtimeex.IWebhook
}

func (s *sms) Send(req interface{}) error {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return s.webhook.Send(s.url, string(jsonReq))
}

func NewSms(webhook runtimeex.IWebhook, url string) isms.ISMS {
	return &sms{
		webhook: webhook,
		url:     url,
	}
}
