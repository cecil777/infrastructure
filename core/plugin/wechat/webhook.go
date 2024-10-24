package wechat

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

type SendTextBody struct {
	MsgType string          `json:"msgtype"`
	Text    sendTextContent `json:"text"`
}

type sendTextContent struct {
	Content string `json:"content"`
}

type sendTextResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type webhook struct {
}

func (w *webhook) Send(url, text string) error {
	textContent := sendTextContent{
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

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.SetBody(jsonBody)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return err
	}

	result := sendTextResp{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

func NewWebhook() *webhook {
	return &webhook{}
}
