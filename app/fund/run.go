package fund

import (
	"fmt"
	"gospider/app/fund/model"
	"gospider/global"
	"gospider/utils"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

type FundResult struct {
	SheetName string
	Result    *model.FundModel
}

func Run() string {
	var fundChan = make(chan string, 10)
	var resultChan = make(chan FundResult, 10)

	var wg sync.WaitGroup
	wg.Add(1)
	go search(fundChan, &wg)

	wg.Add(1)
	go fundDetailFetch(fundChan, resultChan, &wg)

	waiting4Write := make(map[string][]*model.FundModel)
	go func() {
		for item := range resultChan {
			waiting4Write[item.SheetName] = append(waiting4Write[item.SheetName], item.Result)
			global.GPA_LOG.Info(fmt.Sprintf("%s - %s, %s", item.SheetName, item.Result.BaseData.Code, item.Result.BaseData.Name))
		}
	}()

	wg.Wait()

	writer := new(utils.Writer)
	fileName := fmt.Sprintf("基金汇总_%s", time.Now().Format("2006-01-02"))
	writer.New(fileName)
	for k, v := range waiting4Write {
		writer.NewStreamWriter(k)
		ncols := writer.WriteHeader(model.FundModel{})
		for i, item := range v {
			writer.WriteRow(item, i+2)
		}
		cell_pre, _ := excelize.CoordinatesToCellName(ncols, len(v)+2)
		writer.StreamWriter.AddTable("A1", cell_pre, "")
		writer.StreamWriter.Flush()
	}
	if writer.WorkBook.SheetCount >= 2 {
		writer.WorkBook.SetActiveSheet(1)
		writer.WorkBook.DeleteSheet("Sheet1")
	}

	writer.Save()

	return fileName
}
