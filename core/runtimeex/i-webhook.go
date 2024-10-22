package runtimeex

// IWebhook is 钩子
type IWebhook interface {
	Send(url, text string) error
}
