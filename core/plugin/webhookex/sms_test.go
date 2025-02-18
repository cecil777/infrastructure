package webhookex

import (
	"encoding/json"
	"testing"

	"github.com/cecil777/infrastructure/core/plugin/wechat"
)

func Test_sms_Send(t *testing.T) {
	type fileds struct {
		MsgType string `json:"msgtype"`
		News    string `json:"news"`
	}
	var f fileds
	jsonStr := `{
		"msgtype": "news",
		"news": {
		   "articles" : [
			   {
				   "title" : "中秋节礼品领取",
				   "description" : "今年中秋节公司有豪礼相送",
				   "url" : "www.qq.com",
				   "picurl" : "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"
			   }
			]
		}
	}`
	_ = json.Unmarshal([]byte(jsonStr), &f)
	tests := []struct {
		name    string
		args    fileds
		wantErr bool
	}{
		{
			name:    "Test_sms_Send",
			args:    f,
			wantErr: false,
		},
	}
	iw := wechat.NewWebhook()
	s := NewSms(iw, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=52204af2-4248-46ef-bd00-46ec0262772f")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Send(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
