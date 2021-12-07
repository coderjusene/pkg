package mail

import "testing"

func TestSend(t *testing.T) {
	options := &Options{
		MailHost: "smtp.163.com",
		MailPort: 465,
		MailUser: "13588@163.com",
		MailPass: "xxx",
		MailTo:   "7108@qq.com",
		Subject:  "go mail test",
		Body:     "<h1>hello world</h1>",
	}

	err := Send(options)
	if err != nil {
		t.Error(err)
	}

	t.Log("success")
}
