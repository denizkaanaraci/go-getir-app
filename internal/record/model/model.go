package model

import "time"

type Request struct {
	StartDate string `json:"startDate" `
	EndDate   string `json:"endDate" `
	MinCount  int    `json:"minCount" `
	MaxCount  int    `json:"maxCount" `
}

type Response struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Records []Record `json:"record"`
}

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}
