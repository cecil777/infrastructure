package wechat

import (
	"encoding/json"
	"errors"

	"github.com/cecil777/infrastructure/core/runtimeex"
	"github.com/valyala/fasthttp"
)

type sendTextResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type webhook struct {
}

func (w *webhook) Send(url, text string) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.SetBody([]byte(text))
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return err
	}

	result := sendTextResp{}
	err := json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

func NewWebhook() runtimeex.IWebhook {
	return &webhook{}
}
