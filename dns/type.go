package dns

import "github.com/xpartacvs/go-resellerclub/core"

type ActivatingDNSServiceResponse struct {
	Status  string `json:"status"`
	Msg     string `json:"msg"`
	ZoneID  string `json:"zoneid"`
	OrderID string `json:"orderid"`
}

type GeneralResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type SearchingDNSRecordsResponse struct {
	RecSonPage core.JSONUint16 `json:"recsonpage"`
	RecSinDB   core.JSONUint16 `json:"recsindb"`
	PageSearchingDNSRecordsResponse
}

type PageSearchingDNSRecordsResponse map[string]struct {
	TimeToLive core.JSONInt `json:"timetolive"`
	Status     string       `json:"status"`
	Type       string       `json:"type"`
	Host       string       `json:"host"`
	Value      string       `json:"value"`
}
