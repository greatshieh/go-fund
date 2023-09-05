package downloader

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"gospider/global"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Loader struct {
	client *http.Client
}

type Request struct {
	Request *http.Request
	Name    string
}

type Response struct {
	// 响应内容
	Resp []byte
	// 响应类型 json / html
	RespType string
	// 请求地址
	ResqURL *url.URL
}

type Proxy struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func (l *Loader) DownLoad(req *http.Request) (*Response, error) {
	resp, err := l.client.Do(req)

	if err != nil {
		return nil, errors.New("downloader: faild to request")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("downloader: faild to read response")
	}

	if strings.Contains(string(body), "ERROR: The requested URL could not be retrieved") {
		return nil, errors.New("downloader: faild to read response")
	}

	r := new(Response)
	r.ResqURL = resp.Request.URL
	r.Resp = body

	if len(resp.Header["Content-Type"]) == 0 {
		return nil, errors.New("downloader: faild to read response")
	}

	contentType := resp.Header["Content-Type"][0]

	r.RespType = contentType

	return r, nil
}

// 生成新的 request
func NewRequest(method string, url string, body io.Reader, name string, params map[string]string) Request {
	r, _ := http.NewRequest(method, url, body)
	agent := createUserAgent()
	r.Header.Add("User-Agent", agent)

	if params != nil {
		q := r.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		r.URL.RawQuery = q.Encode()
	}

	return Request{Request: r, Name: name}
}

func Byte2Json(resp []byte) interface{} {
	var _jsonSlice interface{}

	json.Unmarshal(resp, &_jsonSlice)

	return _jsonSlice
}

func Byte2String(resp []byte) string {
	return string(resp)
}

// CreateLoader 生成新的下载器
func CreateLoader() *Loader {
	loader := new(Loader)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if global.GPA_CONFIG.Spider.ProxyEnable {
		ipAddr, _ := newProxy()
		tr.Proxy = http.ProxyURL(ipAddr)
	}

	loader.client = &http.Client{
		Transport: tr,
		Timeout:   time.Second * 30,
	}

	return loader
}

// createUserAgent 获取新的 user agent
func createUserAgent() string {
	agents := global.GPA_CONFIG.Spider.UserAgent
	index := rand.Intn(len(agents))
	return agents[index]
}

// 获取新的代理
func newProxy() (*url.URL, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/get", global.GPA_CONFIG.Spider.ProxyHost, global.GPA_CONFIG.Spider.ProxyPort))

	if err != nil {
		return new(url.URL), errors.New("proxy: faild to get proxy")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return new(url.URL), errors.New("proxy: faild to get proxy")
	}

	proxy := new(Proxy)

	err = json.Unmarshal(body, &proxy)

	if err != nil {
		return new(url.URL), errors.New("proxy: faild to get proxy")
	}

	if err != nil {
		return new(url.URL), errors.New("proxy: faild to get proxy")
	}

	ipAddr, _ := url.Parse(fmt.Sprintf("http://%s:%s", proxy.IP, proxy.Port))

	return ipAddr, nil
}
