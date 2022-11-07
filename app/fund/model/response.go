package model

type JsonModel struct {
	BaseInfo    baseInfo    `json:"JJXQ"`
	FeatureData featureInfo `json:"TSSJ"`
	StageGains  stagedGains `json:"JDZF"`
}

type baseInfo struct {
	Data BaseInfoData `json:"Datas"`
}

type BaseInfoData struct {
	Code      string `json:"FCODE"  gorm:"primaryKey;index;column:code;comment:基金代码" excel:"基金代码"`
	Name      string `json:"SHORTNAME" gorm:"column:name;comment:基金名称" excel:"基金名称"`
	Bench     string `json:"BENCH" gorm:"column:bench;comment:业绩比较基准" excel:"业绩比较基准"`
	Ftype     string `json:"FTYPE" gorm:"column:ftype;comment:基金类型" excel:"基金类型"`
	Company   string `json:"JJGS" gorm:"column:company;comment:基金公司" excel:"基金公司"`
	EstabDate string `json:"ESTABDATE" gorm:"column:estabdate;comment:成立时间" excel:"成立时间"`
	Indexname string `json:"INDEXNAME" gorm:"column:indexname;comment:跟踪标的" excel:"跟踪标的"`
	Endnav    string `json:"ENDNAV" gorm:"column:endnav;comment:基金规模" excel:"基金规模"`
	Sgzt      string `json:"SGZT" gorm:"column:sgzt;comment:申购状态" excel:"申购状态"`
	Shzt      string `json:"SHZT" gorm:"column:shzt;comment:赎回状态" excel:"赎回状态"`
}

type featureInfo struct {
	Data FeatureInfoData `json:"Datas"`
}

type FeatureInfoData struct {
	Maxretra1 string `json:"MAXRETRA1" gorm:"column:maxretra1;comment:最大回撤(近1年)" excel:"最大回撤(近1年)"`
	Maxretra3 string `json:"MAXRETRA3" gorm:"column:maxretra3;comment:最大回撤(近3年)" excel:"最大回撤(近3年)"`
	Maxretra5 string `json:"MAXRETRA5" gorm:"column:maxretra5;comment:最大回撤(近5年)" excel:"最大回撤(近5年)"`
	Profit_z  string `json:"PROFIT_Z" gorm:"column:profit_z;comment:盈利概率(1周)" excel:"盈利概率(1周)"`
	Profit_y  string `json:"PROFIT_Y" gorm:"column:profit_y;comment:盈利概率(1月)" excel:"盈利概率(1月)"`
	Profit_3y string `json:"PROFIT_3Y" gorm:"column:profit_3y;comment:盈利概率(3月)" excel:"盈利概率(3月)"`
	Profit_6y string `json:"PROFIT_6Y" gorm:"column:profit_6y;comment:盈利概率(6月)" excel:"盈利概率(6月)"`
	Profit_1n string `json:"PROFIT_1N" gorm:"column:profit_1n;comment:盈利概率(1年)" excel:"盈利概率(1年)"`
	Sharp1    string `json:"SHARP1" gorm:"column:sharp1;comment:夏普率(近1年)" excel:"夏普率(近1年)"`
	Sharp3    string `json:"SHARP3" gorm:"column:sharp3;comment:夏普率(近3年)" excel:"夏普率(近3年)"`
	Sharp5    string `json:"SHARP5" gorm:"column:sharp5;comment:夏普率(近5年)" excel:"夏普率(近5年)"`
	Stddev1   string `json:"STDDEV1" gorm:"column:stddev1;comment:标准差(近1年)" excel:"标准差(近1年)"`
	Stddev3   string `json:"STDDEV3" gorm:"column:stddev3;comment:标准差(近3年)" excel:"标准差(近3年)"`
	Stddev5   string `json:"STDDEV5" gorm:"column:stddev5;comment:标准差(近5年)" excel:"标准差(近5年)"`
}

type stagedGains struct {
	Data []stagedGainsData `json:"Datas"`
}

type stagedGainsData struct {
	Rank  string `json:"rank" gorm:"comment:同类排名" excel:"同类排名"`
	Syl   string `json:"syl" gorm:"comment:涨幅" excel:"涨幅"`
	Title string `json:"title" gorm:"primarykey;comment:周期" excel:"周期"`
}

type YearlyGainsData struct {
	Fcode string `json:"fcode" gorm:"primarykey;column:code;comment:基金代码" excel:"基金代码"`
	Gains string `json:"gains" gorm:"comment:涨幅" excel:"涨幅"`
	Year  string `json:"year" gorm:"primarykey;comment:年度" excel:"年度"`
}
