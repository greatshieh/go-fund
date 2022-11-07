package model

import (
	"time"
)

type FundModel struct {
	BaseInfoData    BaseInfoData     `gorm:"embedded"`
	FeatureInfoData FeatureInfoData  `gorm:"embedded"`
	StagedGainsData stagedGainsModel `gorm:"embedded"`
	HoldInfo        HoldInfo         `json:"hold_info" gorm:"foreignKey:Fcode;comment:持仓信息"`
	// YearlyGains     YearlyGainsModel `json:"yearly_gains" gorm:"foreignKey:Fcode;comment:年度涨幅"`
}

type FundManager struct {
	FempDate    time.Time `mapstructure:"FEMPDATE,omitempty" gorm:"column:fempdate;comment:基金经理任职起始期" excel:"基金经理任职起始期"`
	ManagerName string    `mapstructure:"MGRNAME,omitempty" gorm:"column:mgrname;comment:基金经理" excel:"基金经理"`
	Penavgrowth float32   `mapstructure:"PENAVGROWTH,omitempty" gorm:"column:penavgrowth;comment:任职回报" excel:"任职回报"`
	PracDate    time.Time `mapstructure:"TOTALDAYS,omitempty" gorm:"column:pracdate;comment:基金经理从业时间" excel:"基金经理从业时间"`
}

type HoldInfo struct {
	// Fcode              string `json:"fcode" gorm:"primaryKey;index;comment:基金代码" excel:"基金代码"`
	IndividualHolder   string `json:"individual_holder" gorm:"column:individual_holder;comment:个人持有比例" excel:"个人持有比例"`
	InstitutionsHolder string `json:"institutions_holder" gorm:"column:institutions_holder;comment:机构持有比例" excel:"机构持有比例"`
	InteriorHolder     string `json:"interior_holder" gorm:"column:interior_holder;comment:内部持有比例" excel:"内部持有比例"`
}

type stagedGainsModel struct {
	RankZ  string `json:"rank_z" gorm:"comment:同类排名(近1周)" excel:"同类排名(近1周)"`
	RankY  string `json:"rank_y" gorm:"comment:同类排名(近1月)" excel:"同类排名(近1月)"`
	Rank3Y string `json:"rank_3y" gorm:"comment:同类排名(近3月)" excel:"同类排名(近3月)"`
	Rank6Y string `json:"rank_6y" gorm:"comment:同类排名(近6月)" excel:"同类排名(近6月)"`
	Rank1N string `json:"rank_1n" gorm:"comment:同类排名(近1年)" excel:"同类排名(近1年)"`
	Rank2N string `json:"rank_2n" gorm:"comment:同类排名(近2年)" excel:"同类排名(近2年)"`
	Rank3N string `json:"rank_3n" gorm:"comment:同类排名(近3年)" excel:"同类排名(近3年)"`
	Rank5N string `json:"rank_5n" gorm:"comment:同类排名(近5年)" excel:"同类排名(近5年)"`
	RankJN string `json:"rank_jn" gorm:"comment:同类排名(今年以来)" excel:"同类排名(今年以来)"`
	RankLN string `json:"rank_ln" gorm:"comment:同类排名(成立以来)" excel:"同类排名(成立以来)"`
	SylZ   string `json:"syl_z" gorm:"comment:涨幅(近1周)" excel:"涨幅(近1周)"`
	SylY   string `json:"syl_y" gorm:"comment:涨幅(近1月)" excel:"涨幅(近1月)"`
	Syl3Y  string `json:"syl_3y" gorm:"comment:涨幅(近3月)" excel:"涨幅(近3月)"`
	Syl6Y  string `json:"syl_6y" gorm:"comment:涨幅(近6月)" excel:"涨幅(近6月)"`
	Syl1N  string `json:"syl_1n" gorm:"comment:涨幅(近1年)" excel:"涨幅(近1年)"`
	Syl2N  string `json:"syl_2n" gorm:"comment:涨幅(近2年)" excel:"涨幅(近2年)"`
	Syl3N  string `json:"syl_3n" gorm:"comment:涨幅(近3年)" excel:"涨幅(近3年)"`
	Syl5N  string `json:"syl_5n" gorm:"comment:涨幅(近5年)" excel:"涨幅(近5年)"`
	SylJN  string `json:"syl_jn" gorm:"comment:涨幅(今年以来)" excel:"涨幅(今年以来)"`
	SylLN  string `json:"syl_ln" gorm:"comment:涨幅(成立以来)" excel:"涨幅(成立以来)"`
}

type YearlyGainsModel struct {
	// Fcode     string `json:"fcode" gorm:"primaryKey;index;comment:基金代码" excel:"基金代码"`
	Gains2014 string `json:"gains_2014" gorm:"comment:涨幅(2014)" excel:"涨幅(2014)"`
	Gains2015 string `json:"gains_2015" gorm:"comment:涨幅(2015)" excel:"涨幅(2015)"`
	Gains2016 string `json:"gains_2016" gorm:"comment:涨幅(2016)" excel:"涨幅(2016)"`
	Gains2017 string `json:"gains_2017" gorm:"comment:涨幅(2017)" excel:"涨幅(2017)"`
	Gains2018 string `json:"gains_2018" gorm:"comment:涨幅(2018)" excel:"涨幅(2018)"`
	Gains2019 string `json:"gains_2019" gorm:"comment:涨幅(2019)" excel:"涨幅(2019)"`
	Gains2020 string `json:"gains_2020" gorm:"comment:涨幅(2020)" excel:"涨幅(2020)"`
	Gains2021 string `json:"gains_2021" gorm:"comment:涨幅(2021)" excel:"涨幅(2021)"`
}
