package parser

import (
	"encoding/json"
	"fmt"
	"gospider/app/fund/model"
	"gospider/downloader"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getCode(rawQuery string) string {
	query, _ := url.ParseQuery(rawQuery)
	code := query["code"][0]
	return code
}

func map2Struct(resp map[string]interface{}, fundStruct interface{}) {
	refType := reflect.TypeOf(fundStruct).Elem()
	refVal := reflect.ValueOf(fundStruct).Elem()

	for i := 0; i < refType.NumField(); i++ {
		fieldtype := refType.Field(i).Type.Elem().Name()
		fieldVal := refVal.Field(i)
		key := strings.Split(refType.Field(i).Tag.Get("mapstructure"), ",")[0]

		if _, ok := resp[key]; !ok {
			continue
		}
		if resp[key] == nil {
			continue
		}
		mapVal := resp[key].(string)
		if mapVal == "--" || mapVal == "" {
			continue
		}
		newValue := reflect.Value{}
		switch fieldtype {
		case "float32", "float64":
			floatVal, err := strconv.ParseFloat(mapVal, 32)
			if err != nil {
				break
			}
			float32Value := float32(floatVal)
			newValue = reflect.ValueOf(&float32Value)
		case "uint":
			floatVal, err := strconv.ParseFloat(mapVal, 32)
			if err != nil {
				break
			}
			uint64Value := uint64(floatVal)
			uintValue := uint(uint64Value)
			newValue = reflect.ValueOf(&uintValue)
		case "Time":
			var timeValue time.Time
			if key == "TOTALDAYS" {
				days, err := strconv.ParseFloat(mapVal, 32)
				if err != nil {
					break
				}
				today := time.Now()
				hours := days * 24
				dur, _ := time.ParseDuration(fmt.Sprintf("-%dh", int(hours)))
				timeValue = today.Add(dur)
			} else {
				var err error
				timeValue, err = time.Parse("2006-01-02", mapVal)
				if err != nil {
					break
				}
			}
			newValue = reflect.ValueOf(&timeValue)
		default:
			newValue = reflect.ValueOf(&mapVal)
		}
		fieldVal.Set(newValue)
	}
}

func yearlyGainsParser(resp string) []model.YearlyGainsData {
	f := new([]model.YearlyGainsData)

	re := regexp.MustCompile(`content:"(.*)"`)
	regex := re.FindStringSubmatch(resp)

	if len(regex) == 0 {
		return *f
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(regex[1]))
	if err != nil {
		return *f
	}

	dom.Find("tr th").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}
		value := s.Text()
		year := strings.Replace(value, "年度", "", -1)
		gains := dom.Find("tbody tr:nth-child(1)").Find("td").Eq(i).Text()
		*f = append(*f, model.YearlyGainsData{Gains: gains, Year: year})
	})

	return *f
}

func baseInfoParser(resp []byte) model.JsonModel {
	result := new(model.JsonModel)
	json.Unmarshal(resp, result)
	return *result
}

func holdInfoParser(resp string) model.HoldInfo {
	f := new(model.HoldInfo)

	re := regexp.MustCompile(`content:"(.*)"`)
	regex := re.FindStringSubmatch(resp)
	if len(regex) == 0 {
		return *f
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(regex[1]))
	if err != nil {
		return *f
	}

	dom.Find("tbody tr:first-child").Each(func(_ int, s *goquery.Selection) {
		s.Find("td:nth-of-type(n+2)").Each(func(i int, a *goquery.Selection) {
			value := a.Text()
			switch i {
			case 0:
				f.InstitutionsHolder = value
			case 1:
				f.IndividualHolder = value
			case 2:
				f.InteriorHolder = value
			default:
				return
			}
		})
	})

	return *f
}

// ParseSelector 解释选择器,根据网页响应内容选择解释器
func ParseSelector(resp *downloader.Response) interface{} {
	if strings.Contains(resp.RespType, "json") {
		return baseInfoParser(resp.Resp)
	} else {
		if strings.Contains(string(resp.Resp), "cyrjg") {
			model := holdInfoParser(string(resp.Resp))
			// code := getCode(resp.ResqURL.RawQuery)
			// model.Fcode = code
			return model
		} else {
			models := yearlyGainsParser(string(resp.Resp))
			// code := getCode(resp.ResqURL.RawQuery)
			// for i := 0; i < len(models); i++ {
			// 	models[i].Fcode = code
			// }
			return models
		}
	}
}
