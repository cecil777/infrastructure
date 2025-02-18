package wechat

import "testing"

func TestWebhook_Send(t *testing.T) {
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
			name: "TestWebhook_Send",
			args: args{
				url: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=52204af2-4248-46ef-bd00-46ec0262772f",
				text: `{
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
				}`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &webhook{}
			if err := w.Send(tt.args.url, tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
