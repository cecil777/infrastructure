package sms

// ISMS is 短信接口
type ISMS interface {
	Send(req interface{}) error
}
