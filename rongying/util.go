package rongying

import (
	"crypto/md5"
	b64 "encoding/base64"
	"fmt"
	"github.com/levigross/grequests"
	"io"
	"strings"
	"time"
)

func UtilTimestamp() string {
	objTime := time.Now()
	return objTime.Format("20060102150405")
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

func Base64Encode(str string) string {
	return b64.StdEncoding.EncodeToString([]byte(str))
}

func SendPost(url string, data interface{}, auth string) (resp *grequests.Response) {
	session := grequests.NewSession(nil)
	resp, err := session.Post(url,
		&grequests.RequestOptions{
			JSON: data,
			Headers: map[string]string{
				"Accept":        "application/json",
				"Content-Type":  "application/json;charset=utf-8",
				"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36",
				"Authorization": auth,
			},
		})
	if err == nil && resp.StatusCode == 200 {
		return resp
	}
	return resp
}
