// 从且慢上获取基金的相关系数
package fund

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

// type

// 生成时间戳
func CreateTimeStamp() (string, string) {
	// 生成时间戳
	timeStamp := time.Now().UnixMilli()
	// 数字时间戳转化为字符串
	timeStampString := strconv.FormatInt(timeStamp, 10)

	// 时间戳计算, 并转化为字符串
	e := strconv.FormatInt(int64(math.Floor(float64(timeStamp)*1.01)), 10)

	return timeStampString, e
}

// 时间戳每一位转化为ascii码
func String2Bytes(timeStamp string) []int {
	timeBytes := []int{}
	for _, c := range timeStamp {
		timeBytes = append(timeBytes, int(rune(c)))
	}

	return timeBytes
}

// UnsignedRightShitf 无符号右移
func UnsignedRightShitf(n int, i int) int32 {
	// 如果 n < 0, 先取绝对值, 求补码再移位
	if n < 0 {
		return int32(^(uint32(math.Abs(float64(n)) - 1)) >> i)
	}

	return int32(n >> i)
}

var globalI [64]int

func CreateGloablI() {
	e := func(e int) bool {
		for t, n := math.Sqrt(float64(e)), 2; float64(n) <= t; n++ {
			if e%n == 0 {
				return false
			}
		}
		return true
	}

	t := func(e float64) int {
		return int(int32(int(4294967296 * (e - float64(int32(e))))))
	}

	for n, r := 2, 0; r < 64; {
		if e(n) {
			globalI[r] = t(math.Pow(float64(n), float64(1)/3))
			r++
		}
		n++
	}
}

func CreateSign(e [8]int, t []int, n int) [8]int {
	var o [64]int
	r := e[0]
	a := e[1]
	s := e[2]
	u := e[3]
	c := e[4]
	l := e[5]
	f := e[6]
	p := e[7]

	for d := 0; d < 64; d++ {
		if d < 16 {
			o[d] = int(int32(t[n+d]))
		} else {
			h := int(int32(o[d-15]))
			v := (int32(h<<25) | UnsignedRightShitf(h, 7)) ^ (int32(h<<14) | UnsignedRightShitf(h, 18)) ^ UnsignedRightShitf(h, 3)
			y := int(int32(o[d-2]))
			m := (int32(y<<15) | UnsignedRightShitf(y, 17)) ^ (int32(y<<13) | UnsignedRightShitf(y, 19)) ^ UnsignedRightShitf(y, 10)
			o[d] = int(v) + o[d-7] + int(m) + o[d-16]
		}

		g := int(int32(r&a) ^ int32(r&s) ^ int32(a&s))
		b := int((int32(r<<30) | UnsignedRightShitf(r, 2)) ^ (int32(r<<19) | UnsignedRightShitf(r, 13)) ^ (int32(r<<10) | UnsignedRightShitf(r, 22)))
		z := p + int((int32(c<<26)|UnsignedRightShitf(c, 6))^(int32(c<<21)|UnsignedRightShitf(c, 11))^(int32(c<<7)|UnsignedRightShitf(c, 25))) + int(int32(c&l)^(int32(^c)&int32(f))) + globalI[d] + o[d]

		p = f
		f = l
		l = c
		c = int(int32(u + z))
		u = s
		s = a
		a = r
		r = int(int32(z + b + g))
	}
	e[0] = int(int32(e[0] + r))
	e[1] = int(int32(e[1] + a))
	e[2] = int(int32(e[2] + s))
	e[3] = int(int32(e[3] + u))
	e[4] = int(int32(e[4] + c))
	e[5] = int(int32(e[5] + l))
	e[6] = int(int32(e[6] + f))
	e[7] = int(int32(e[7] + p))

	return e
}

func TransSign(signStamp string) string {
	// 获取时间戳
	// timeStamp := CreateTimeStamp()

	// 时间戳转化为ascii码
	timeBytes := String2Bytes(signStamp)

	n := [8]int{1779033703, 3144134277, 1013904242, 2773480762, 1359893119, 2600822924, 528734635, 1541459225}

	o := 8 * len(timeBytes)
	total := 15 + ((o + 64) >> 9 << 4) + 1
	i := func(e []int) []int {
		t := make([]int, total)
		for n, r := 0, 0; n < len(e); n, r = n+1, r+8 {
			t[r>>5] = int(int32(t[r>>5]) | int32(e[n]<<(24-r%32)))
		}
		return t
	}(timeBytes)

	i[o>>5] = int(int32(i[o>>5]) | int32(128<<(24-(o%32))))
	i[15+((o+64)>>9<<4)] = o

	for s := 0; s < len(i); s += 16 {
		n = CreateSign(n, i, s)
	}

	u := func(e [8]int) []int {
		var t []int
		for n := 0; n < 32*len(e); n += 8 {
			index := UnsignedRightShitf(n, 5)
			t = append(t, int(int32(UnsignedRightShitf(e[index], (24-(n%32)))&255)))
		}
		return t
	}(n)

	var signBuilder strings.Builder
	for _, e := range u {
		r := strconv.FormatInt(int64(e), 16)
		if len(r) >= 2 {
			signBuilder.WriteString(r)
		} else {
			for i := 0; i < 2-len(r); i++ {
				signBuilder.WriteString("0")
			}
			signBuilder.WriteString(r)
		}
	}

	return strings.ToUpper(signBuilder.String())
}

// CreateXSign 生成x-sign签名
func CreateXSign() string {
	timeStampString, timeStamp := CreateTimeStamp()
	// timeStampString := "1669381765617"
	// timeStamp := "1686075583273"
	sign := TransSign(timeStamp)

	// 截取sign的前32位
	var signBuilder strings.Builder
	for i, e := range sign {
		if i >= 32 {
			break
		}
		signBuilder.WriteString(string(e))
	}

	return timeStampString + strings.ToUpper(signBuilder.String())
}

func CreateRequestID(url string) string {
	timeStamp, _ := CreateTimeStamp()
	random := strconv.FormatFloat(rand.Float64(), 'f', 16, 64)
	firstID := "184928216bd6b1-06aa62bf74a8474-7d5d5475-2359296-184928216be907"
	requestID := TransSign(random + timeStamp + url + firstID)

	// 截短requestID, 取后20位
	var requestIDBuilder strings.Builder
	for i := len(requestID) - 20; i < len(requestID); i++ {
		requestIDBuilder.WriteString(fmt.Sprintf("%c", requestID[i]))
	}

	return "albus." + strings.ToUpper(requestIDBuilder.String())
}

type FundInvestType struct {
	FundInvestType string `json:"fundInvestType"`
}

// GetFundType 获取基金的类型代码
func GetFundType(code string) string {
	requests, err := http.NewRequest("GET", fmt.Sprintf("https://qieman.com/pmdj/v1/search/funds?q=%s&limit=6&tradableFunds=1&includeInvestPoolTag=1", code), nil)

	if err != nil {
		panic(err)
	}

	// 生成请求头的request-id
	url := "/pmdj/v1/search/funds"
	requestID := CreateRequestID(url)
	requests.Header.Add("x-request-id", requestID)

	// 生成请求头的x-sign
	xsign := CreateXSign()
	requests.Header.Add("x-sign", xsign)

	requests.Header.Add("Content-Type", "application/json;charset=UTF-8")

	req, err := http.DefaultClient.Do(requests)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	reqBytes, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fundType := []FundInvestType{}
	json.Unmarshal(reqBytes, &fundType)

	if len(fundType) == 0 {
		return ""
	}

	return fundType[0].FundInvestType
}

// GetReturnHistory 组合历史收益回测
func GetReturnHistory() {
	// 投资组合 代码: 比例
	load := `{"100032":0.0968,"160513":0.1498,"510080":0.1016,"519702":0.1038,"005177":0.0987,"000991":0.1003,"003986":0.0946,"003304":0.1022,"003547":0.1522}`

	playLoad, err := json.Marshal(load)
	if err != nil {
		panic(err)
	}
	requests, err := http.NewRequest("POST", "https://qieman.com/pmdj/v1/pomodels/return-history", bytes.NewBuffer(playLoad))

	if err != nil {
		panic(err)
	}

	// 生成请求头的request-id
	url := "/pmdj/v1/pomodels/return-history"
	requestID := CreateRequestID(url)
	requests.Header.Add("x-request-id", requestID)

	// 生成请求头的x-sign
	xsign := CreateXSign()
	requests.Header.Add("x-sign", xsign)

	requests.Header.Add("Content-Type", "application/json;charset=UTF-8")
	requests.Header.Add("Referer", "https://qieman.com/portfolios/ZH129476/analyze")
	requests.Header.Add("Origin", "https://qieman.com")
	requests.Header.Add("Host", "qieman.com")
	requests.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56")
	req, err := http.DefaultClient.Do(requests)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()
}

type Relativity struct {
	Codes []string   `json:"codes"`
	Data  [][]string `json:"data"`
}

// GetFundsRelativity 获取基金的相关性
func GetFundsRelativity() Relativity {
	funds := []string{"090007", "260112", "003547"}

	load := make([]string, len(funds))

	for i, code := range funds {
		fundType := GetFundType(code)

		if fundType == "" {
			break
		}

		load[i] = fmt.Sprintf("%s:%s", code, fundType)
	}

	jsonStr, err := json.Marshal(load)
	if err != nil {
		panic(err)
	}
	requests, err := http.NewRequest("POST", "https://qieman.com/pmdj/v1/funds/relativity", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic("err")
	}

	// 生成请求头的request-id
	url := "/pmdj/v1/funds/relativity"
	requestID := CreateRequestID(url)
	requests.Header.Add("x-request-id", requestID)

	xsign := CreateXSign()
	requests.Header.Add("x-sign", xsign)

	requests.Header.Add("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	req, err := client.Do(requests)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	bodyres, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	relativity := Relativity{}
	json.Unmarshal(bodyres, &relativity)

	return relativity
}

func RrelativityRun() {
	CreateGloablI()

	relativity := GetFundsRelativity()
	// fmt.Println(relativity)

	// 创建一个excel文档
	f := excelize.NewFile()
	name := f.GetSheetName(0)
	writer, err := f.NewStreamWriter(name)
	if err != nil {
		panic(err)
	}

	// 写入第一行
	row := make([]interface{}, len(relativity.Codes)+1)
	row[0] = ""
	for i, v := range relativity.Codes {
		row[i+1] = v
	}
	writer.SetRow("A1", row)

	n := 2
	for i, v := range relativity.Data {
		row := make([]interface{}, len(v)+1)
		row[0] = relativity.Codes[i]
		for j, e := range v {
			row[j+1] = e
		}
		writer.SetRow(fmt.Sprintf("A%d", n), row)
		n++
	}

	writer.Flush()

	f.SaveAs("relativity.xlsx")

	// 写入基金代码

}
