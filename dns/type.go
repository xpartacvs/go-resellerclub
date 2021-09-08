package dns

type StdResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type ActivatingDNSServiceResponse struct {
	StdResponse
	ZoneID  string `json:"zoneid"`
	OrderID string `json:"orderid"`
}

type SearchingDNSRecords struct {
	RecsOnPage string
	Recsindb   string
	Records    []*DNSRecord
}

type DNSRecord struct {
	TimeToLive string `json:"timetolive,omitempty"`
	Status     string `json:"status,omitempty"`
	Type       string `json:"type,omitempty"`
	Host       string `json:"host,omitempty"`
	Value      string `json:"value,omitempty"`
}
