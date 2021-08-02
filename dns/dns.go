package dns

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/xpartacvs/go-resellerclub/core"
)

type DNS interface {
	ActivatingDNSService(orderID string) (*ActivatingDNSServiceResponse, error)
	AddingIPv4AddressRecord(domainName, value, host string, ttl int) (*GeneralResponse, error)
	AddingIPv6AddressRecord(domainName, value, host string, ttl int) (*GeneralResponse, error)
	AddingCNAMERecord(domainName, value, host string, ttl int) (*GeneralResponse, error)
	AddingMXRecord(domainName, value, host string, ttl, priority int) (*GeneralResponse, error)
	AddingNSRecord(domainName, value, host string, ttl int) (*GeneralResponse, error)
	AddingTXTRecord(domainName, value, host string, ttl int) (*GeneralResponse, error)
	AddingSRVRecord(domainName, value, host string, ttl, priority, port, weight int) (*GeneralResponse, error)
	ModifyingIPv4AddressRecord(domainName, host, currentValue, newValue string, ttl int) (*GeneralResponse, error)
	ModifyingIPv6AddressRecord(domainName, host, currentValue, newValue string, ttl int) (*GeneralResponse, error)
	ModifyingCNAMERecord(domainName, host, currentValue, newValue string, ttl int) (*GeneralResponse, error)
	ModifyingMXRecord(domainName, host, currentValue, newValue string, ttl, priority int) (*GeneralResponse, error)
	ModifyingNSRecord(domainName, host, currentValue, newValue string, ttl int) (*GeneralResponse, error)
	ModifyingTXTRecord(domainName, host, currentValue, newValue string, ttl int) (*GeneralResponse, error)
	ModifyingSRVRecord(domainName, host, currentValue, newValue string, ttl, priority, port, weight int) (*GeneralResponse, error)
	ModifyingSOARecord(domainName, responsiblePerson string, refresh, retry, expire, ttl int) (*GeneralResponse, error)
	SearchingDNSRecords(domainName, typeRecord string, noOfRecords, pageNo int, host, value string) (*SearchingDNSRecordsResponse, error)
	DeletingDNSRecord(host, value string) (*GeneralResponse, error)
	DeletingIPv4AddressRecord(domainName, host, value string) (*GeneralResponse, error)
	DeletingIPv6AddressRecord(domainName, host, value string) (*GeneralResponse, error)
	DeletingCNAMERecord(domainName, host, value string) (*GeneralResponse, error)
	DeletingMXRecord(domainName, host, value string) (*GeneralResponse, error)
	DeletingNSRecord(domainName, host, value string) (*GeneralResponse, error)
	DeletingTXTRecord(domainName, host, value string) (*GeneralResponse, error)
	DeletingSRVRecord(domainName, host, value string, port, weight int) (*GeneralResponse, error)
}

func New(c core.Core) DNS {
	return &dns{c}
}

type dns struct {
	core core.Core
}

func (d *dns) ActivatingDNSService(orderID string) (*ActivatingDNSServiceResponse, error) {
	data := make(url.Values)
	data.Add("order-id", orderID)

	resp, err := d.core.CallApi(http.MethodPost, "dns", "activate", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result ActivatingDNSServiceResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) AddingIPv4AddressRecord(domainName, value, host string, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("value", value)
	data.Add("host", host)
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/add-ipv4-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) AddingIPv6AddressRecord(domainName, value, host string, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("value", value)
	data.Add("host", host)
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/add-ipv6-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) AddingCNAMERecord(domainName, value, host string, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("value", value)
	data.Add("host", host)
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/add-cname-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) AddingMXRecord(domainName, value, host string, ttl, priority int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("value", value)
	data.Add("host", host)
	data.Add("ttl", strconv.Itoa(ttl))
	data.Add("priority", strconv.Itoa(priority))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/add-mx-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) AddingNSRecord(domainName, value, host string, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("value", value)
	data.Add("host", host)
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/add-ns-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) AddingTXTRecord(domainName, value, host string, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("value", value)
	data.Add("host", host)
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/manage/add-ns-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) AddingSRVRecord(domainName, value, host string, ttl, priority, port, weight int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("value", value)
	data.Add("host", host)
	data.Add("ttl", strconv.Itoa(ttl))
	data.Add("priority", strconv.Itoa(priority))
	data.Add("port", strconv.Itoa(port))
	data.Add("weight", strconv.Itoa(weight))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/add-srv-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) ModifyingIPv4AddressRecord(domainName, host, currentValue, newValue string, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("current-value", currentValue)
	data.Add("new-value", newValue)
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/update-ipv4-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) ModifyingIPv6AddressRecord(domainName, host, currentValue, newValue string, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("current-value", currentValue)
	data.Add("new-value", newValue)
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/update-ipv6-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) ModifyingCNAMERecord(domainName, host, currentValue, newValue string, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("current-value", currentValue)
	data.Add("new-value", newValue)
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/update-cname-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) ModifyingMXRecord(domainName, host, currentValue, newValue string, ttl, priority int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("current-value", currentValue)
	data.Add("new-value", newValue)
	data.Add("ttl", strconv.Itoa(ttl))
	data.Add("priority", strconv.Itoa(priority))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/update-mx-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) ModifyingNSRecord(domainName, host, currentValue, newValue string, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("current-value", currentValue)
	data.Add("new-value", newValue)
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/update-ns-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) ModifyingTXTRecord(domainName, host, currentValue, newValue string, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("current-value", currentValue)
	data.Add("new-value", newValue)
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/update-txt-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) ModifyingSRVRecord(domainName, host, currentValue, newValue string, ttl, priority, port, weight int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("current-value", currentValue)
	data.Add("new-value", newValue)
	data.Add("ttl", strconv.Itoa(ttl))
	data.Add("priority", strconv.Itoa(priority))
	data.Add("port", strconv.Itoa(port))
	data.Add("weight", strconv.Itoa(weight))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/update-srv-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) ModifyingSOARecord(domainName, responsiblePerson string, refresh, retry, expire, ttl int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("responsible-person", responsiblePerson)
	data.Add("refresh", strconv.Itoa(refresh))
	data.Add("retry", strconv.Itoa(retry))
	data.Add("expire", strconv.Itoa(expire))
	data.Add("ttl", strconv.Itoa(ttl))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/update-soa-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) SearchingDNSRecords(domainName, typeRecord string, noOfRecords, pageNo int, host, value string) (*SearchingDNSRecordsResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("type", typeRecord)
	data.Add("no-of-records", strconv.Itoa(noOfRecords))
	data.Add("page-no", strconv.Itoa(pageNo))
	data.Add("host", host)
	data.Add("value", value)

	resp, err := d.core.CallApi(http.MethodGet, "dns", "manage/search-records", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result SearchingDNSRecordsResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) DeletingDNSRecord(host, value string) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("host", host)
	data.Add("value", value)

	resp, err := d.core.CallApi(http.MethodGet, "dns", "manage/delete-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) DeletingIPv4AddressRecord(domainName, host, value string) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("value", value)

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/delete-ipv4-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) DeletingIPv6AddressRecord(domainName, host, value string) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("value", value)

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/delete-ipv6-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) DeletingCNAMERecord(domainName, host, value string) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("value", value)

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/delete-cname-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) DeletingMXRecord(domainName, host, value string) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("value", value)

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/delete-mx-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) DeletingNSRecord(domainName, host, value string) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("value", value)

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/delete-ns-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) DeletingTXTRecord(domainName, host, value string) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("value", value)

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/delete-txt-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *dns) DeletingSRVRecord(domainName, host, value string, port, weight int) (*GeneralResponse, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)
	data.Add("host", host)
	data.Add("value", value)
	data.Add("port", strconv.Itoa(port))
	data.Add("weight", strconv.Itoa(weight))

	resp, err := d.core.CallApi(http.MethodPost, "dns", "manage/delete-srv-record", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errResponse.Message))
	}

	var result GeneralResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
