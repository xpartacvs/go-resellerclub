package customer

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

type customer struct {
	core core.Core
}

type Customer interface {
	SignUp(regForm *RegistrationForm) error
	ChangePassword(customerId, newPassword string) error
	Details(customerIdOrEmail string) (CustomerDetail, error)
	Delete(customerId string) error
}

func (c *customer) Delete(customerId string) error {
	if !core.RgxNumber.MatchString(customerId) {
		return core.ErrRcInvalidCredential
	}

	resp, err := c.core.CallApi(http.MethodPost, "customers", "delete", url.Values{"customer-id": {customerId}})
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

	boolResult, err := strconv.ParseBool(string(bytesResp))
	if err != nil {
		return err
	}
	if !boolResult {
		return core.ErrRcOperationFailed
	}

	return nil
}

func (c *customer) Details(customerIdOrEmail string) (CustomerDetail, error) {
	ret := CustomerDetail{}
	data := url.Values{}

	var funcName, query string
	switch {
	case core.RgxEmail.MatchString(customerIdOrEmail):
		funcName = "details"
		query = "username"
	case core.RgxNumber.MatchString(customerIdOrEmail):
		funcName = "details-by-id"
		query = "customer-id"
	default:
		return ret, core.ErrRcInvalidCredential
	}
	data.Add(query, customerIdOrEmail)

	resp, err := c.core.CallApi(http.MethodGet, "customers", funcName, data)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return ret, err
		}
		return ret, errors.New(strings.ToLower(errResponse.Message))
	}

	if err = json.Unmarshal(bytesResp, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

func (c *customer) ChangePassword(customerId, newPassword string) error {
	if !matchPasswordWithPattern(newPassword, true) {
		return errors.New("invalid password format")
	}

	data := url.Values{}
	data.Add("customer-id", customerId)
	data.Add("new-passwd", newPassword)

	resp, err := c.core.CallApi(http.MethodPost, "customers/v2", "change-password", data)
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

	boolResult, err := strconv.ParseBool(string(bytesResp))
	if err != nil {
		return err
	}
	if !boolResult {
		return core.ErrRcOperationFailed
	}

	return nil
}

func (c *customer) SignUp(regForm *RegistrationForm) error {
	urlValues, err := regForm.UrlValues()
	if err != nil {
		return err
	}
	resp, err := c.core.CallApi(http.MethodPost, "customers/v2", "signup", urlValues)
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

	regForm.CustomerId = string(bytesResp)
	return nil
}

func New(c core.Core) Customer {
	return &customer{
		core: c,
	}
}
