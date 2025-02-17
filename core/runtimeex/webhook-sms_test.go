package runtimeex

import (
	"github.com/cecil777/infrastructure/core/plugin/wechat"
	"testing"
)

func Test_sms_Send(t *testing.T) {
	type fields struct {
		Text string
	}
	tests := []struct {
		name    string
		args    fields
		wantErr bool
	}{
		{
			name:    "Test_sms_Send",
			args:    fields{"msg from webhook sms"},
			wantErr: false,
		},
	}
	iw := wechat.NewWebhook()
	s := NewSms(iw, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=284cc574-7e21-4333-bb9f-6f87e0308529")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Send(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
