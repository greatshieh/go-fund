package model

type SearchResponse struct {
	Data []SearchResponseData `json:"Data"`
}

type SearchResponseData struct {
	Code string `json:"fundCode"`
	// Name string `json:"ftype"`
}

type BaseDataResponse struct {
	Data       []BaseData `json:"data"`
	Success    bool       `json:"success"`
	TotalCount int        `json:"totalCount"`
}

type GainDataResponse struct {
	Data       []GainData `json:"data"`
	Success    bool       `json:"success"`
	TotalCount int        `json:"totalCount"`
}

type StageGainsResponseData struct {
	Rank  string `json:"rank"`
	Syl   string `json:"syl"`
	Title string `json:"title"`
}

type StageGainResponse struct {
	Data []StageGainsResponseData `json:"data"`
}
