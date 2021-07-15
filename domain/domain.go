package domain

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/xpartacvs/go-resellerclub/core"
)

type domain struct {
	core core.Core
}

type Domain interface {
	SearchOrders(criteria OrderCriteria) error
	CheckAvailability(domainsWithoutTLD, tlds []string) (DomainAvailabilities, error)
	SuggestNames(keyword, tldOnly string, exactMatch, adult bool) (SuggestNames, error)
	GetCustomerDefaultNameServers(customerID int) ([]string, error)
}

func New(c core.Core) Domain {
	return &domain{
		core: c,
	}
}

func (d *domain) CheckAvailability(domainsWithoutTLD, tlds []string) (DomainAvailabilities, error) {
	if len(domainsWithoutTLD) <= 0 || len(tlds) <= 0 {
		return DomainAvailabilities{}, errors.New("domainnames and tlds must not empty")
	}

	data := url.Values{}
	wg := sync.WaitGroup{}
	rwMutex := sync.RWMutex{}

	for _, v := range domainsWithoutTLD {
		wg.Add(1)
		go func(value string) {
			defer wg.Done()
			defer rwMutex.Unlock()
			rwMutex.Lock()
			data.Add("domain-name", value)
		}(v)
	}
	for _, v := range tlds {
		wg.Add(1)
		go func(value string) {
			defer wg.Done()
			defer rwMutex.Unlock()
			rwMutex.Lock()
			data.Add("tlds", value)
		}(v)
	}
	wg.Wait()

	resp, err := d.core.CallApi(http.MethodGet, "domains", "available", data)
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

	availabilities := DomainAvailabilities{}
	err = json.Unmarshal(bytesResp, &availabilities)
	if err != nil {
		return nil, err
	}

	return availabilities, nil
}

func (d *domain) SearchOrders(criteria OrderCriteria) error {
	urlValues, err := criteria.UrlValues()
	if err != nil {
		return err
	}
	resp, err := d.core.CallApi(http.MethodGet, "domains", "search", urlValues)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	d.core.PrintResponse(bytesResp)

	return nil
}

func (d *domain) SuggestNames(keyword, tldOnly string, exactMatch, adult bool) (SuggestNames, error) {
	data := make(url.Values)
	data.Add("keyword", keyword)
	data.Add("tld-only", tldOnly)
	data.Add("exact-match", strconv.FormatBool(exactMatch))
	data.Add("adult", strconv.FormatBool(adult))

	resp, err := d.core.CallApi(http.MethodGet, "domains/v5", "suggest-names", data)
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

	suggestNames := SuggestNames{}
	err = json.Unmarshal(bytesResp, &suggestNames)
	if err != nil {
		return nil, err
	}

	return suggestNames, nil
}

// TODO
// Missing: attr-name Map[name], attrValue Map[value]
func (d *domain) Register(domainName string, years int, ns []string, customerID, regContactID, adminContactID, techContactID, billingContactID int, invoiceOption string, purchasePrivacy bool, protectPrivacy bool, autoRenew bool, discountAmount float64, purchasePremiumDNS bool) error {
	// data := make(url.Values)
	return nil
}

// TODO
// Problems: can't get response for renew order-id
func (d *domain) Renew(orderID, years, expDate int, purchasePrivacy, autoRenew bool, invoiceOption string, discountAmount float64, purchasePremiumDNS bool) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	data.Add("exp-date", strconv.Itoa(expDate))
	data.Add("purchase-privacy", strconv.FormatBool(purchasePrivacy))
	data.Add("auto-renew", strconv.FormatBool(autoRenew))
	data.Add("invoice-option", invoiceOption)
	data.Add("discount-amount", strconv.FormatFloat(discountAmount, 'f', 2, 64))
	data.Add("purchase-premium-dns", strconv.FormatBool(purchasePremiumDNS))

	resp, err := d.core.CallApi(http.MethodPost, "domains", "renew", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return err
		}
		return errors.New(strings.ToLower(errResponse.Message))
	}

	return nil
}

func (d *domain) GetCustomerDefaultNameServers(customerID int) ([]string, error) {
	data := make(url.Values)
	data.Add("customer-id", strconv.Itoa(customerID))

	resp, err := d.core.CallApi(http.MethodGet, "domains", "customer-default-ns", data)
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

	result := make([]string, 0)
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *domain) GetOrderID(domainName string) error {
	data := make(url.Values)
	data.Add("domain-name", domainName)

	resp, err := d.core.CallApi(http.MethodGet, "domains", "order-id", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return err
		}
		return errors.New(strings.ToLower(errResponse.Message))
	}

	return nil
}

func (d *domain) GetRegistrationOrderDetails(orderID int, options []string) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	for _, option := range options {
		data.Add("options", option)
	}

	resp, err := d.core.CallApi(http.MethodGet, "domains", "details", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return err
		}
		return errors.New(strings.ToLower(errResponse.Message))
	}

	return nil
}

func (d *domain) ModifyNameServers(orderID int, ns []string) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	for _, n := range ns {
		data.Add("ns", n)
	}

	resp, err := d.core.CallApi(http.MethodGet, "domains", "modify-ns", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return err
		}
		return errors.New(strings.ToLower(errResponse.Message))
	}

	return nil
}

func (d *domain) AddChildNameServer(orderID int, cns string, ips []string) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	data.Add("cns", cns)
	for _, ip := range ips {
		data.Add("ip", ip)
	}

	resp, err := d.core.CallApi(http.MethodGet, "domains", "add-cns", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return err
		}
		return errors.New(strings.ToLower(errResponse.Message))
	}

	return nil
}
