package fund

import (
	"fmt"
	"gospider/app/fund/model"
	"gospider/app/fund/parser"
	"gospider/downloader"
	"reflect"
	"strings"
	"sync"
)

type FundInfo struct {
	Code string
	Name string
}

var URLS = map[string]string{"base information": "https://j5.fund.eastmoney.com/sc/tfs/qt/v2.0.1/%s.json", "yearly gains": "http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=yearzf&code=%s", "holder information": "http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=cyrjg&code=%s"}

func fundResp(fund FundInfo, name string) (*downloader.Response, error) {
	// 生成新的请求
	request := downloader.NewRequest("GET", fmt.Sprintf(URLS[name], fund.Code), nil, name, nil)
	loader := downloader.CreateLoader()
	resp, err := loader.DownLoad(request.Request)

	return resp, err
}

func respProcessor(resp *downloader.Response) interface{} {
	// m :=
	// err := model.SaveSelector(m)
	return parser.ParseSelector(resp)
}

func CopyStruct(src interface{}, dst interface{}) {
	dstValue := reflect.ValueOf(dst).Elem()
	dstType := reflect.TypeOf(dst).Elem()

	srcValue := reflect.ValueOf(src)

	for i := 0; i < dstType.NumField(); i++ {
		fieldVal := dstValue.Field(i)
		fieldName := dstType.Field(i).Name
		fieldVal.Set(srcValue.FieldByName(fieldName))
	}

}

// getInfo 获取基金的基本信息
func getInfo(modelChan chan<- interface{}, fund FundInfo, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		resp, err := fundResp(fund, name)

		if err != nil {
			continue
		}

		modelChan <- respProcessor(resp)
		break
	}
}

func downloadController(code string, name string) *model.FundModel {
	fund := FundInfo{Code: code, Name: name}

	var wig sync.WaitGroup

	infoType := []string{"base information", "holder information"} //, "yearly gains", }

	var modelChan = make(chan interface{}, 2)
	for _, t := range infoType {
		wig.Add(1)
		go getInfo(modelChan, fund, t, &wig)
	}

	wig.Wait()
	close(modelChan)

	modelResult := new(model.FundModel)

	for item := range modelChan {
		refType := reflect.TypeOf(item)
		switch refType.String() {
		case "model.HoldInfo":
			CopyStruct(item.(model.HoldInfo), &modelResult.HoldInfo)
		case "model.JsonModel":
			CopyStruct(item.(model.JsonModel).BaseInfo.Data, &modelResult.BaseInfoData)
			CopyStruct(item.(model.JsonModel).FeatureData.Data, &modelResult.FeatureInfoData)
			for _, v := range item.(model.JsonModel).StageGains.Data {
				title := v.Title
				srcVal := reflect.ValueOf(v)
				refVal := reflect.ValueOf(&modelResult.StagedGainsData).Elem()
				RankVal := refVal.FieldByName(fmt.Sprintf("Rank%s", strings.ToUpper(title)))
				RankVal.Set(srcVal.FieldByName("Rank"))
				SylVal := refVal.FieldByName(fmt.Sprintf("Syl%s", strings.ToUpper(title)))
				SylVal.Set(reflect.ValueOf(v.Syl))
			}
		default:
		}
	}

	return modelResult

}

func fundDetailFetch(fundChan <-chan []string, resultChan chan<- FundResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range fundChan {
		code := item[0]
		name := item[1]
		resultChan <- FundResult{SheetName: item[2], Result: downloadController(code, name)}
	}

	close(resultChan)
}
