package wechat

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

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

type webhook struct {
}

func NewWebhook() *webhook {
	return &webhook{}
}

func (w *webhook) Send(url, text string) error {
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

	result := SendTextResp{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return err
	}

	if result.Errcode != 0 {
		return errors.New(result.Errmsg)
	}
	return nil
}
