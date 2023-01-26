package fund

import (
	"encoding/json"
	"fmt"
	"gospider/downloader"
	"gospider/global"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type TableInfo struct {
	Data      []string `json:"datas"`
	AllPages  string   `json:"allPages"`
	PageIndex string   `json:"pageIndex"`
	PageNum   string   `json:"pageNum"`
	Datacount string   `json:"datacount"`
}

// 需要采集的基金类型
var fundtypes = map[string]string{"gp": "股票型", "hh": "混合型", "zq": "债券型", "zs": "指数型"}

// 需要采集的排名周期
var period = map[string]map[string]int{
	"3y": {"index": 7, "threshold": 3},
	"6y": {"index": 8, "threshold": 3},
	"1n": {"index": 9, "threshold": 4},
	"2n": {"index": 10, "threshold": 4},
	"3n": {"index": 11, "threshold": 4},
	"5n": {"index": 12, "threshold": 4},
	// "jn": {"index": 5, "threshold": 4},
}

// GET请求参数
var params = map[string]string{"dt": "4",
	"ft": "",
	"sd": "",
	"ed": "",
	"sc": "",
	"st": "desc",
	"pi": "1",
	"pn": "100",
	"zf": "diy",
	"sh": "list",
}

// 排名筛选链接
const URL = "http://fund.eastmoney.com/data/FundGuideapi.aspx"

var locker sync.Mutex

// rankParser 解析数据
func rankParser(resp string, index int, threshold int, records *map[string]int) (bool, string) {
	re := regexp.MustCompile(`var rankData =(.*)`)
	regex := re.FindStringSubmatch(resp)

	result := new(TableInfo)
	json.Unmarshal([]byte(regex[1]), result)

	data := result.Data

	limitedCount, _ := strconv.ParseInt(result.Datacount, 10, 64)
	limitedCount = int64(limitedCount / int64(threshold))
	// var limitedCount int64 = 4

	pageIndex, _ := strconv.ParseInt(result.PageIndex, 10, 64)
	pageNum, _ := strconv.ParseInt(result.PageNum, 10, 64)

	total := (pageIndex - 1) * pageNum

	for _, item := range data {
		if strings.Split(item, ",")[index] == "" {
			return true, "无有效数据"
		}

		locker.Lock()
		(*records)[item]++
		locker.Unlock()

		total++
		if total >= limitedCount {
			return true, fmt.Sprintf("超过排名 %d / %d", total, limitedCount)
		}
	}

	return result.PageIndex == result.AllPages, ""
}

// map类型拷贝
func copyMap(src map[string]string) map[string]string {
	p := make(map[string]string)

	for key, value := range src {
		p[key] = value
	}

	return p
}

func RankFetch(v map[string]int, records *map[string]int, wg *sync.WaitGroup, sc string, name string, p map[string]string) {
	defer wg.Done()
	p["sc"] = sc
	for pi := 1; ; pi++ {
		p["pi"] = strconv.Itoa(pi)
		var resp *downloader.Response
		var err error
		for {
			request := downloader.NewRequest("GET", URL, nil, "", p)
			loader := downloader.CreateLoader()
			resp, err = loader.DownLoad(request.Request)
			if err == nil && strings.Contains(string(resp.Resp), "var rankData =") {
				break
			}
		}
		isEnd, reason := rankParser(string(resp.Resp), v["index"], v["threshold"], records)
		if isEnd {
			global.GPA_LOG.Info(fmt.Sprintf("%s, %s 第 %d 页完成. 采集结束: %s", name, sc, pi, reason))
			break
		}
		global.GPA_LOG.Info(fmt.Sprintf("%s, %s 第 %d 页完成.", name, sc, pi))
	}
}

func RankRun(fundChan chan<- []string, mainWg *sync.WaitGroup) {
	defer mainWg.Done()
	for key, value := range fundtypes {
		params["ft"] = key

		// 新建一个map，用于保存基金出现的次数
		var records = make(map[string]int)

		var wg sync.WaitGroup
		for sc, v := range period {
			wg.Add(1)
			p := copyMap(params)
			p["sc"] = sc
			go RankFetch(v, &records, &wg, sc, value, p)
		}

		wg.Wait()

		nrows := 0
		for item, count := range records {
			if count == len(period) {
				content := strings.Split(item, ",")
				fundChan <- []string{content[0], content[1], value}
				nrows++
			}
		}

		global.GPA_LOG.Info(fmt.Sprintf("%s 共采集 %d 条数据", value, nrows))
	}

	close(fundChan)
}
