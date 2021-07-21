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
	GetOrderID(domainName string) (string, error)
	GetRegistrationOrderDetails(orderID string, options []string) (*OrderDetail, error)
	ModifyNameServers(orderID string, ns []string) (*ModifyNameServersResponse, error)
	AddChildNameServer(orderID string, cns string, ips []string) (*AddChildNameServerResponse, error)
	ModifyPrivacyProtectionStatus(orderID string, protectPrivacy bool, reason string) (*ModifyPrivacyProtectionStatusResponse, error)
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
func (d *domain) Register(domainName string, years int, ns []string, customerID, regContactID, adminContactID, techContactID, billingContactID int, invoiceOption string, purchasePrivacy bool, protectPrivacy bool, autoRenew bool, attrName, attrValue string, discountAmount float64, purchasePremiumDNS bool) error {
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

func (d *domain) GetOrderID(domainName string) (string, error) {
	data := make(url.Values)
	data.Add("domain-name", domainName)

	resp, err := d.core.CallApi(http.MethodGet, "domains", "orderid", data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return "", err
		}
		return "", errors.New(strings.ToLower(errResponse.Message))
	}

	return string(bytesResp), nil
}

func (d *domain) GetRegistrationOrderDetails(orderID string, options []string) (*OrderDetail, error) {
	data := make(url.Values)
	data.Add("order-id", orderID)
	for _, option := range options {
		data.Add("options", option)
	}

	resp, err := d.core.CallApi(http.MethodGet, "domains", "details", data)
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

	var orderDetail OrderDetail
	err = json.Unmarshal(bytesResp, &orderDetail)
	if err != nil {
		return nil, err
	}

	return &orderDetail, nil
}

func (d *domain) ModifyNameServers(orderID string, ns []string) (*ModifyNameServersResponse, error) {
	data := make(url.Values)
	data.Add("order-id", orderID)
	for _, n := range ns {
		data.Add("ns", n)
	}

	resp, err := d.core.CallApi(http.MethodPost, "domains", "modify-ns", data)
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

	var result ModifyNameServersResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *domain) AddChildNameServer(orderID string, cns string, ips []string) (*AddChildNameServerResponse, error) {
	data := make(url.Values)
	data.Add("order-id", orderID)
	data.Add("cns", cns)
	for _, ip := range ips {
		data.Add("ip", ip)
	}

	resp, err := d.core.CallApi(http.MethodPost, "domains", "add-cns", data)
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

	var result AddChildNameServerResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *domain) ModifyChildNameServerHostName(orderID string, oldCNS, newCNS string) error {
	data := make(url.Values)
	data.Add("order-id", orderID)
	data.Add("old-cns", oldCNS)
	data.Add("new-cns", newCNS)

	resp, err := d.core.CallApi(http.MethodPost, "domains", "modify-cns-name", data)
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

func (d *domain) ModifyChildNameServerIPAddress(orderID string, cns, oldIP, newIP string) error {
	data := make(url.Values)
	data.Add("order-id", orderID)
	data.Add("cns", cns)
	data.Add("old-ip", oldIP)
	data.Add("new-ip", newIP)

	resp, err := d.core.CallApi(http.MethodPost, "domains", "modify-cns-ip", data)
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

func (d *domain) DeletingChildNameServerIPAddress(orderID string, cns string, ips []string) error {
	data := make(url.Values)
	data.Add("order-id", orderID)
	data.Add("cns", cns)
	for _, ip := range ips {
		data.Add("ip", ip)
	}

	resp, err := d.core.CallApi(http.MethodPost, "domains", "delete-cns-ip", data)
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

func (d *domain) ModifyContacts(orderID, regContactID, adminContactID, techContactID, billingContactID int, sixtyDayLockOptout, designatedAgent bool, attrName, attrValue string) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	data.Add("reg-contact-id", strconv.Itoa(regContactID))
	data.Add("admin-contact-id", strconv.Itoa(adminContactID))
	data.Add("tech-contact-id", strconv.Itoa(techContactID))
	data.Add("billing-contact-id", strconv.Itoa(billingContactID))
	data.Add("sixty-day-lock-optout", strconv.FormatBool(sixtyDayLockOptout))
	data.Add("designated-agent", strconv.FormatBool(designatedAgent))
	data.Add("attr-name", attrName)
	data.Add("attr-value", attrValue)

	resp, err := d.core.CallApi(http.MethodPost, "domains", "modify-contact", data)
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

func (d *domain) ModifyPrivacyProtectionStatus(orderID string, protectPrivacy bool, reason string) (*ModifyPrivacyProtectionStatusResponse, error) {
	data := make(url.Values)
	data.Add("order-id", orderID)
	data.Add("protect-privacy", strconv.FormatBool(protectPrivacy))
	data.Add("reason", reason)

	resp, err := d.core.CallApi(http.MethodPost, "domains", "modify-privacy-protection", data)
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

	var result ModifyPrivacyProtectionStatusResponse
	err = json.Unmarshal(bytesResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *domain) ModifyAuthCode(orderID int, authCode string) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	data.Add("auth-code", authCode)

	resp, err := d.core.CallApi(http.MethodPost, "domains", "modify-auth-code", data)
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

func (d *domain) ApplyTheftProtectionLock(orderID int) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))

	resp, err := d.core.CallApi(http.MethodPost, "domains", "enable-theft-protection", data)
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

func (d *domain) RemoveTheftProtectionLock(orderID int) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))

	resp, err := d.core.CallApi(http.MethodPost, "domains", "disable-theft-protection", data)
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

func (d *domain) GettheListofLocksAppliedOnDomainName(orderID int) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))

	resp, err := d.core.CallApi(http.MethodGet, "domains", "locks", data)
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

// Issues:
// - Missing documentation
// func (d *domain) GetTELCTHLoginDetails()

func (d *domain) ModifyTELWhoisPreference(orderID int, whoisType, publish string) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	data.Add("whois-type", whoisType)
	data.Add("publish", publish)

	resp, err := d.core.CallApi(http.MethodPost, "domains", "tel/modify-whois-pref", data)
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

func (d *domain) ResendTransferApprovalMail(orderID int) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))

	resp, err := d.core.CallApi(http.MethodPost, "domains", "resend-rfa", data)
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

func (d *domain) ReleaseUKDomainName(orderID int, newTag string) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	data.Add("new-tag", newTag)

	resp, err := d.core.CallApi(http.MethodPost, "domains", "uk/release", data)
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

func (d *domain) CancelTransfer(orderID int) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))

	resp, err := d.core.CallApi(http.MethodPost, "domains", "cancel-transfer", data)
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

func (d *domain) Suspend(orderID int, reason string) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	data.Add("reason", reason)

	resp, err := d.core.CallApi(http.MethodPost, "orders", "suspend", data)
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

func (d *domain) Unsuspend(orderID int) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))

	resp, err := d.core.CallApi(http.MethodPost, "orders", "unsuspend", data)
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

func (d *domain) Delete(orderID int) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))

	resp, err := d.core.CallApi(http.MethodPost, "domains", "delete", data)
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

func (d *domain) Restore(orderID int, invoiceOption string) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	data.Add("invoice-option", invoiceOption)

	resp, err := d.core.CallApi(http.MethodPost, "domains", "restore", data)
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

func (d *domain) RecheckingNSWithDERegistry(orderID int) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))

	resp, err := d.core.CallApi(http.MethodPost, "domains", "de/recheck-ns", data)
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

func (d *domain) AssociatingOrDissociatingXXXMembershipTokenID(orderID int, associationID string) error {
	data := make(url.Values)
	data.Add("order-id", strconv.Itoa(orderID))
	data.Add("association-id", associationID)

	resp, err := d.core.CallApi(http.MethodPost, "domains", "dotxxx/association-details", data)
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
