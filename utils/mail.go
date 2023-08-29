package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"gospider/global"
	"io"
	"net/smtp"
	"path"
	"runtime"
	"strings"

	"github.com/jordan-wright/email"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func utf8togbk(s string) string {
	sbyte := []byte(s)
	reader := transform.NewReader(bytes.NewReader(sbyte), simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		panic(e)
	}
	return string(d)
}

// getPath 获取当前工作目录
func getPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(path.Dir(filename))
	}
	return abPath
}

// 发送邮件
func SendEmail(fileName string) error {
	to := strings.Split(global.GPA_CONFIG.Email.To, ",")
	from := global.GPA_CONFIG.Email.From
	nickname := global.GPA_CONFIG.Email.Nickname
	secret := global.GPA_CONFIG.Email.Secret
	host := global.GPA_CONFIG.Email.Host
	port := global.GPA_CONFIG.Email.Port
	isSSL := global.GPA_CONFIG.Email.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = "基金数据"
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)

	attachment, _ := e.AttachFile(path.Join(getPath(), "/", fmt.Sprintf("%s.xlsx", fileName)))

	fileName = utf8togbk(fileName)
	attachment.Filename = fmt.Sprintf("%s.xlsx", fileName)

	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	if err != nil {
		panic(err)
	}
	return err
}
