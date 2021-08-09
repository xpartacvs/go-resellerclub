package dns

import "github.com/xpartacvs/go-resellerclub/core"

type StdResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type ActivatingDNSServiceResponse struct {
	StdResponse
	ZoneID  string `json:"zoneid"`
	OrderID string `json:"orderid"`
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
