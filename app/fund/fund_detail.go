package fund

import (
	"encoding/json"
	"fmt"
	"gospider/app/fund/model"
	"gospider/downloader"
	"strings"
	"sync"
)

type FundInfo struct {
	Code string
	Name string
}

func fetch(url string, name string, params map[string]string) (*downloader.Response, error) {
	// 生成新的请求
	request := downloader.NewRequest("GET", url, nil, name, params)
	loader := downloader.CreateLoader()
	resp, err := loader.DownLoad(request.Request)

	return resp, err
}

func parseGainData(resp *downloader.Response) model.GainData {
	r := new(model.GainDataResponse)

	if err := json.Unmarshal(resp.Resp, r); err != nil {
		fmt.Println(err)
	}
	return r.Data[0]
}

func NewGainRequestParam(code string) map[string]string {
	return map[string]string{
		"FIELDS": "PV_Y,DTCOUNT_Y,STDDEV1,MAXRETRA1,SHARP1,STDDEV3,MAXRETRA3,SHARP3,STDDEV5,MAXRETRA5,SHARP5,NVRENO_Y,NVRENO_TRY,NVRENO_FY,CALMAR_Y,CALMAR_TRY,CALMAR_FY",
		"FCODES": code,
	}
}

// 获取收益数据
func getGainData(code string, name string, m *model.FundModel, wg *sync.WaitGroup) {
	defer wg.Done()

	url := "https://fundcomapi.tiantianfunds.com/mm/FundMNewApi/FundUniqueInfo"

	params := NewGainRequestParam(code)

	for {
		resp, err := fetch(url, name, params)

		if err != nil {
			continue
		}

		// 解析收益数据
		data := parseGainData(resp)
		m.GainData = data
		return
	}
}

func NewBaseRequestParam(code string) map[string]string {
	return map[string]string{
		"FIELDS": "FCODE,FTYPE,ESTABDATE,ENDNAV,SHORTNAME,JJJL,JJGS,BENCH,RISKLEVEL,INVTGT,INVSTRA",
		"FCODES": code,
	}
}

func parseBaseData(resp *downloader.Response) model.BaseData {
	r := new(model.BaseDataResponse)

	if err := json.Unmarshal(resp.Resp, r); err != nil {
		fmt.Println(err)
	}
	return r.Data[0]
}

// 获取基本信息
func getBaseData(code string, name string, m *model.FundModel, wg *sync.WaitGroup) {
	defer wg.Done()

	url := "https://fundcomapi.tiantianfunds.com/mm/FundMNewApi/FundBaseInfos"

	params := NewBaseRequestParam(code)

	for {
		resp, err := fetch(url, name, params)

		if err != nil {
			continue
		}

		// 解析基本数据
		data := parseBaseData(resp)
		m.BaseData = data
		return
	}
}

func NewStageGainRequestParam(code string) map[string]string {
	return map[string]string{
		"FCODE": code,
	}
}

func parseStageGainData(resp *downloader.Response) model.StageGainData {
	r := new(model.StageGainResponse)

	if err := json.Unmarshal(resp.Resp, r); err != nil {
		fmt.Println(err)
	}

	m := model.StageGainData{}

	for _, v := range r.Data {
		switch v.Title {
		case "Y":
			m.SylY = v.Syl
			m.RankY = v.Rank
		case "3Y":
			m.Syl3Y = v.Syl
			m.Rank3Y = v.Rank
		case "6Y":
			m.Syl6Y = v.Syl
			m.Rank6Y = v.Rank
		case "1N":
			m.Syl1N = v.Syl
			m.Rank1N = v.Rank
		case "2N":
			m.Syl2N = v.Syl
			m.Rank2N = v.Rank
		case "3N":
			m.Syl3N = v.Syl
			m.Rank3N = v.Rank
		case "5N":
			m.Syl5N = v.Syl
			m.Rank5N = v.Rank
		case "JN":
			m.SylJN = v.Syl
			m.RankJN = v.Rank
		}
	}

	return m
}

// 获取阶段信息
func getStageGainData(code string, name string, m *model.FundModel, wg *sync.WaitGroup) {
	defer wg.Done()

	url := "https://fundcomapi.tiantianfunds.com/mm/FundMNewApi/FundPeriodIncrease"

	params := NewStageGainRequestParam(code)

	for {
		resp, err := fetch(url, name, params)

		if err != nil {
			continue
		}

		// 解析基本数据
		data := parseStageGainData(resp)
		m.StageGainData = data
		return
	}
}

func fundDetailFetch(fundChan <-chan string, resultChan chan<- FundResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range fundChan {
		code := item

		m := new(model.FundModel)

		var fundwg sync.WaitGroup
		// 获取基本信息
		fundwg.Add(1)
		go getBaseData(code, "基本信息", m, &fundwg)
		// 获取基金的收益数据
		fundwg.Add(1)
		go getGainData(code, "收益数据", m, &fundwg)
		// 获取阶段收益
		fundwg.Add(1)
		go getStageGainData(code, "阶段收益", m, &fundwg)

		fundwg.Wait()

		sheetName := strings.Split(m.BaseData.Ftype, "-")[0]
		resultChan <- FundResult{SheetName: sheetName, Result: m}
	}

	close(resultChan)
}
