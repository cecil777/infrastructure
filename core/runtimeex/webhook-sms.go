package runtimeex

import (
	"encoding/json"
)

type sms struct {
	Text string `json:"text"`

	url      string
	iWebhook IWebhook
}

func (s *sms) Send(req interface{}) error {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonReq, s)
	if err != nil {
		return err
	}
	return s.iWebhook.Send(s.url, s.Text)
}

func NewSms(iWebhook IWebhook, url string) *sms {
	return &sms{iWebhook: iWebhook, url: url}
}
