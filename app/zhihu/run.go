package zhihu

// var ZHIHUBASEURL = "https://www.zhihu.com/api/v4/topics/%s/feeds/essence?offset=%d&limit=10&include=data[?(target.type=topic_sticky_module)].target.data[?(target.type=answer)].target.content,relationship.is_authorized,is_author,voting,is_thanked,is_nothelp;data[?(target.type=topic_sticky_module)].target.data[?(target.type=answer)].target.is_normal,comment_count,voteup_count,content,relevant_info,excerpt.author.badge[?(type=best_answerer)].topics;data[?(target.type=topic_sticky_module)].target.data[?(target.type=article)].target.content,voteup_count,comment_count,voting,author.badge[?(type=best_answerer)].topics;data[?(target.type=topic_sticky_module)].target.data[?(target.type=people)].target.answer_count,articles_count,gender,follower_count,is_followed,is_following,badge[?(type=best_answerer)].topics;data[?(target.type=answer)].target.annotation_detail,content,hermes_label,is_labeled,relationship.is_authorized,is_author,voting,is_thanked,is_nothelp,answer_type;data[?(target.type=answer)].target.author.badge[?(type=best_answerer)].topics;data[?(target.type=answer)].target.paid_info;data[?(target.type=article)].target.annotation_detail,content,hermes_label,is_labeled,author.badge[?(type=best_answerer)].topics;data[?(target.type=question)].target.annotation_detail,comment_count;"

// type ZhihuSpider struct {
// 	Code    string
// 	Tag     string
// 	PageNum int
// 	URL     string
// 	Client  *http.Client
// }

// func (f *ZhihuSpider) CreateURL() {
// 	f.URL = fmt.Sprintf(ZHIHUBASEURL, f.Code, (f.PageNum-1)*10)
// }

// func (f *ZhihuSpider) CreateClien() {
// 	f.Client = app.CreateClient()
// }

// func (f *ZhihuSpider) New() {
// 	f.CreateURL()
// 	f.CreateClien()
// }

// func (f *ZhihuSpider) Run() ([]model.ZhihuModel, bool) {
// 	global.GPA_LOG.Info(fmt.Sprintf("正在获取 %s 的 第 %d 页", f.Tag, f.PageNum))
// 	for {
// 		request := app.NewRequest("GET", f.URL)
// 		resp, err := app.Fetch(request, f.Client)
// 		if err != nil {
// 			f.CreateClien()
// 		} else {
// 			webContent := app.Byte2Interface(resp)
// 			if webContent == nil {
// 				f.CreateClien()
// 				continue
// 			}
// 			return parseJson(webContent, f.Tag)
// 		}
// 	}
// }

// func map2Struct(resp map[string]interface{}, ZhihuStruct interface{}) {
// 	refType := reflect.TypeOf(ZhihuStruct).Elem()
// 	refVal := reflect.ValueOf(ZhihuStruct).Elem()

// 	answerID := strconv.Itoa(int(resp["id"].(float64)))

// 	for i := 0; i < refType.NumField(); i++ {
// 		fieldVal := refVal.Field(i)
// 		key := strings.Split(refType.Field(i).Tag.Get("json"), ",")[0]

// 		if _, ok := resp[key]; !ok {
// 			continue
// 		}
// 		if resp[key] == nil {
// 			continue
// 		}

// 		newValue := reflect.Value{}

// 		switch reflect.TypeOf(resp[key]).Name() {
// 		case "float64":
// 			intVal := int(resp[key].(float64))
// 			newValue = reflect.ValueOf(&intVal)
// 		default:
// 			stringVal := resp[key].(string)
// 			newValue = reflect.ValueOf(&stringVal)
// 		}
// 		fieldVal.Set(newValue)
// 	}

// 	if question, ok := resp["question"]; ok {
// 		title := question.(map[string]interface{})["title"].(string)
// 		ZhihuStruct.(*model.ZhihuModel).QuestionTitle = &title

// 		questionID := strconv.Itoa(int(question.(map[string]interface{})["id"].(float64)))
// 		answerLink := fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s", questionID, answerID)
// 		ZhihuStruct.(*model.ZhihuModel).AnswerLink = &answerLink
// 	} else {
// 		answerLink := fmt.Sprintf("https://zhuanlan.zhihu.com/p/%s", answerID)
// 		ZhihuStruct.(*model.ZhihuModel).AnswerLink = &answerLink
// 	}
// }

// func parseJson(resp interface{}, tag string) ([]model.ZhihuModel, bool) {
// 	var result []model.ZhihuModel
// 	if respVal, ok := resp.(map[string]interface{})["data"]; ok {
// 		if respVal == nil {
// 			return result, true
// 		}
// 		for _, data := range respVal.([]interface{}) {
// 			if a, ok := data.(map[string]interface{})["target"]; ok {
// 				zhihu := new(model.ZhihuModel)
// 				zhihu.Tag = &tag
// 				map2Struct(a.(map[string]interface{}), zhihu)
// 				if *zhihu.UpvoteCount < 100 {
// 					return result, true
// 				}
// 				result = append(result, *zhihu)
// 			}
// 		}
// 	}

// 	if paging, ok := resp.(map[string]interface{})["paging"]; ok {
// 		return result, paging.(map[string]interface{})["is_end"].(bool)
// 	}

// 	return result, true
// }

// // readTopics 从文本文件中读取信息
// func readTopics(ch chan<- *ZhihuSpider, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	filepath := path.Join(global.GPA_CONFIG.Spider.FilePath, "/", "zhihu")
// 	f, err := os.Open(filepath)
// 	if err != nil {
// 		panic("打开话题文件错误")
// 	}

// 	defer f.Close()

// 	r := bufio.NewReader(f)

// 	for {
// 		lineString, err := r.ReadString('\n')
// 		if err != nil && err != io.EOF {
// 			panic("读取文件错误")
// 		}
// 		if err == io.EOF {
// 			break
// 		}

// 		if !strings.Contains(lineString, "基金") {
// 			continue
// 		}

// 		line := strings.TrimSpace(lineString)
// 		stringLine := strings.Split(line, ",")
// 		zhihu := ZhihuSpider{Code: stringLine[0], Tag: stringLine[1], PageNum: 1}
// 		zhihu.New()
// 		ch <- &zhihu
// 	}
// 	close(ch)
// }

// func save(ch <-chan model.ZhihuModel, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for c := range ch {
// 		if (c != model.ZhihuModel{}) {
// 			c.Save()
// 		}
// 	}
// }

// // start 启动爬虫过程
// func start(ch <-chan *ZhihuSpider, sh chan<- model.ZhihuModel, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for c := range ch {
// 		for {
// 			models, isEnd := c.Run()
// 			for _, m := range models {
// 				sh <- m
// 			}
// 			if isEnd {
// 				global.GPA_LOG.Info(fmt.Sprintf("🏁 %s 抓取完成, 共 %d 页", c.Tag, c.PageNum))
// 				break
// 			}

// 			c.PageNum++
// 			c.New()
// 		}
// 	}
// }

// type Result struct {
// 	Tag string
// }

// func Run() {
// 	var wg sync.WaitGroup
// 	var wgSave sync.WaitGroup
// 	fundsChan := make(chan *ZhihuSpider, 50)
// 	saveChan := make(chan model.ZhihuModel, 200)
// 	wg.Add(1)
// 	go readTopics(fundsChan, &wg)

// 	for i := 0; i < 50; i++ {
// 		wg.Add(1)
// 		go start(fundsChan, saveChan, &wg)
// 	}

// 	wgSave.Add(1)
// 	go save(saveChan, &wgSave)

// 	wg.Wait()
// 	close(saveChan)
// 	wgSave.Wait()

// 	toExcel()
// }

// func toExcel() {
// 	var results []Result
// 	global.GPA_DB.Table("zhihu_topic").Distinct("tag").Select("tag").Scan(&results)
// 	writer := utils.Writer{}
// 	writer.New("知乎")
// 	for _, r := range results {
// 		tag := r.Tag
// 		writer.NewStreamWriter(tag)
// 		writer.WriteHeader(model.ZhihuModel{})
// 		var questions []model.ZhihuModel
// 		global.GPA_DB.Where("tag = ?", tag).Find(&questions)
// 		writer.WriteRows(questions, 2)
// 	}
// 	writer.Save()
// }
