package wechat

import (
	"core/runtimeex"
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

var _ runtimeex.IWebhook = (*Wxwork)(nil)

type Wxwork struct {
}

type SendTextBody struct {
	MsgType string          `json:"msgtype"`
	Text    SendTextContent `json:"text"`
}

type SendTextContent struct {
	Content string `json:"content"`
}

type SendTextResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (w *Wxwork) Send(url string, text string) error {
	textContent := SendTextContent{
		Content: text,
	}
	body := SendTextBody{
		MsgType: "text",
		Text:    textContent,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req := &fasthttp.Request{}
	req.SetRequestURI(url)
	req.SetBody(jsonBody)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	resp := &fasthttp.Response{}
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return err
	}

	b := string(resp.Body())
	str := []byte(b)
	result := SendTextResp{}
	err = json.Unmarshal(str, &result)
	if err != nil {
		return err
	}

	if result.Errcode != 0 {
		return errors.New(result.Errmsg)
	}
	return nil
}
