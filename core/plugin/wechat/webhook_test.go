package wechat

import "testing"

func TestWxWorkSend(t *testing.T) {
	type args struct {
		url  string
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestWxwork_Send",
			args: args{
				url:  "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=284cc574-7e21-4333-bb9f-6f87e0308529",
				text: "企业微信机器人 hello world",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wxwork{}
			if err := w.Send(tt.args.url, tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}