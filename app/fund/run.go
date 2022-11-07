package fund

import (
	"fmt"
	"gospider/app/fund/model"
	"gospider/global"
	"gospider/utils"
	"sync"
)

type FundResult struct {
	SheetName string
	Result    *model.FundModel
}

func Run() {
	// filePath := path.Join(global.GPA_CONFIG.Spider.FilePath, "/", "funds_all")
	// getAllFunds(filePath)
	// getSpecifies(filePath, "allfund")
	var fundChan = make(chan []string, 10)
	var resultChan = make(chan FundResult, 10)

	var wg sync.WaitGroup
	wg.Add(1)
	go RankRun(fundChan, &wg)

	wg.Add(1)

	go fundDetailFetch(fundChan, resultChan, &wg)

	waiting4Write := make(map[string][]*model.FundModel)
	go func() {
		for item := range resultChan {
			waiting4Write[item.SheetName] = append(waiting4Write[item.SheetName], item.Result)
			global.GPA_LOG.Info(fmt.Sprintf("%s - %s, %s", item.SheetName, item.Result.BaseInfoData.Code, item.Result.BaseInfoData.Name))
		}
	}()

	wg.Wait()

	writer := new(utils.Writer)
	writer.New("result")
	for _, name := range fundtypes {
		writer.NewStreamWriter(name)
		writer.WriteHeader(model.FundModel{})
		for i, item := range waiting4Write[name] {
			writer.WriteRow(item, i+2)
		}
		writer.StreamWriter.Flush()
	}

	writer.Save()

}
