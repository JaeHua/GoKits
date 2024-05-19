package model

// MailCode 验证码+号码结构体
type MailCode struct {
	Mail  string `json:"mail"`
	VCode string `json:"vcode"`
}
