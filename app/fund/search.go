package fund

import (
	"encoding/json"
	"gospider/app/fund/model"
	"gospider/downloader"
	"sync"
)

func rankParser(resp []byte) []model.FundBaseInfo {
	var result model.SearchModel

	json.Unmarshal(resp, &result)

	return result.Data

}

func search(fundChan chan<- string, mainWg *sync.WaitGroup) {
	defer mainWg.Done()

	var resp *downloader.Response
	var err error

	params := map[string]string{
		"bType":        "01,04,21,23,24,22,31,32,33,34",
		"stageRanking": "6_0-25,7_0-25,8_0-25,9_0-25,10_0-25,4_0-33,5_0-33,",
		"isSale":       "1",
		"isBuy":        "1",
		"orderField":   "5_6_-1",
		// "stageMaxReturn": "8_0-10",
		// "cmRanking":      "6_0-10,7_0-10,8_0-10",
	}

	for {
		request := downloader.NewRequest("POST", "https://uni-fundts.1234567.com.cn/dataapi/conditionFund/fundSelect", nil, "", params)
		request.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		loader := downloader.CreateLoader()
		resp, err = loader.DownLoad(request.Request)
		if err == nil {
			break
		}
	}

	for _, v := range rankParser(resp.Resp) {
		fundChan <- v.Code
	}

	close(fundChan)
}
