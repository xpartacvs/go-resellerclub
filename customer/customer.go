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
	SignUp(regForm *SignUpForm) error
	ChangePassword(customerId, newPassword string) error
	Details(customerIdOrEmail string) (CustomerDetail, error)
	Delete(customerId string) error
	ForgotPassword(username string) error
	Suspension(toggle bool, customerId, reason string) error
	Search(criteria CustomerCriteria, offset, limit uint16) (CustomerSearchResult, error)
	Modify(customerIdOrEmail string, changes CustomerDetail) error
	GenerateOTP(customerId string) error
	VerifyOTP(customerId, otp string, authType core.AuthType) (bool, error)
	Authenticate(username, password string) (CustomerDetail, *ErrorAuthentication)
	GenerateToken(username, password, ip string) (string, error)
}

func (c *customer) GenerateToken(username, password, ip string) (string, error) {
	if !matchPasswordWithPattern(password, true) {
		return "", errors.New("invalid format on password")
	}

	if !core.RgxEmail.MatchString(username) {
		return "", errors.New("invalid format on email")
	}

	data := url.Values{}
	data.Add("username", username)
	data.Add("passwd", password)
	data.Add("ip", ip)

	resp, err := c.core.CallApi(http.MethodGet, "customers", "generate-token", data)
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
		if err = json.Unmarshal(bytesResp, &errResponse); err != nil {
			return "", err
		}
		return "", errors.New(strings.ToLower(errResponse.Message))
	}

	return string(bytesResp), nil
}

func (c *customer) Authenticate(username, password string) (CustomerDetail, *ErrorAuthentication) {
	ret := CustomerDetail{}
	errAuth := &ErrorAuthentication{
		JSONStatusResponse: core.JSONStatusResponse{
			Status:  "ERROR",
			Message: "",
		},
	}

	if !core.RgxEmail.MatchString(username) || !matchPasswordWithPattern(password, true) {
		errAuth.Message = "Invalid format of username or password"
		return ret, errAuth
	}

	data := url.Values{}
	data.Add("username", username)
	data.Add("passwd", password)

	resp, err := c.core.CallApi(http.MethodPost, "customers/v2", "authenticate", data)
	if err != nil {
		errAuth.Message = err.Error()
		return ret, errAuth
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		errAuth.Message = err.Error()
		return ret, errAuth
	}

	if resp.StatusCode != http.StatusOK {
		err = json.Unmarshal(bytesResp, errAuth)
		if err != nil {
			errAuth.Message = err.Error()
			return ret, errAuth
		}
		return ret, errAuth
	}

	if err = json.Unmarshal(bytesResp, &ret); err != nil {
		errAuth.Message = err.Error()
		return ret, errAuth
	}

	return ret, nil
}

func (c *customer) VerifyOTP(customerId, otp string, authType core.AuthType) (bool, error) {
	if !core.RgxNumber.MatchString(customerId) {
		return false, core.ErrRcInvalidCredential
	}

	data := url.Values{}
	data.Add("customerid", customerId)
	data.Add("otp", otp)
	data.Add("type", string(authType))

	resp, err := c.core.CallApi(http.MethodPost, "customers/authenticate", "verify-otp", data)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return false, err
		}
		return false, errors.New(strings.ToLower(errResponse.Message))
	}

	return strconv.ParseBool(string(bytesResp))
}

func (c *customer) GenerateOTP(customerId string) error {
	if !core.RgxNumber.MatchString(customerId) {
		return core.ErrRcInvalidCredential
	}

	resp, err := c.core.CallApi(http.MethodGet, "customers/authenticate", "generate-otp", url.Values{"customerid": {customerId}})
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

func (c *customer) Modify(customerIdOrEmail string, changes CustomerDetail) error {
	customerBefore, err := c.Details(customerIdOrEmail)
	if err != nil {
		return nil
	}

	if err := changes.mergePrevious(customerBefore); err != nil {
		return nil
	}

	data, err := changes.UrlValues()
	if err != nil {
		return nil
	}
	data.Add("customer-id", customerBefore.Id)

	resp, err := c.core.CallApi(http.MethodPost, "customers", "modify", data)
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

func (c *customer) Search(criteria CustomerCriteria, offset, limit uint16) (CustomerSearchResult, error) {
	if limit < 10 || limit > 500 {
		return CustomerSearchResult{}, errors.New("limit must be in range of 10 to 500")
	}
	if offset <= 0 {
		return CustomerSearchResult{}, errors.New("offset must greater than 0")
	}

	data, err := criteria.UrlValues()
	if err != nil {
		return CustomerSearchResult{}, err
	}
	data.Add("no-of-records", strconv.FormatUint(uint64(limit), 10))
	data.Add("page-no", strconv.FormatUint(uint64(offset), 10))

	resp, err := c.core.CallApi(http.MethodGet, "customers", "search", data)
	if err != nil {
		return CustomerSearchResult{}, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return CustomerSearchResult{}, err
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := core.JSONStatusResponse{}
		err = json.Unmarshal(bytesResp, &errResponse)
		if err != nil {
			return CustomerSearchResult{}, err
		}
		return CustomerSearchResult{}, errors.New(strings.ToLower(errResponse.Message))
	}

	replacer := strings.NewReplacer("customer.", "")
	strResp := replacer.Replace(string(bytesResp))

	var buffer map[string]core.JSONBytes
	if err := json.Unmarshal([]byte(strResp), &buffer); err != nil {
		return CustomerSearchResult{}, err
	}

	var dataBuffer CustomerDetail
	var dataBuffers []CustomerDetail
	var numMatched int
	for key, dataBytes := range buffer {
		switch {
		case core.RgxNumber.MatchString(key):
			if err := json.Unmarshal(dataBytes, &dataBuffer); err != nil {
				return CustomerSearchResult{}, err
			}
			dataBuffers = append(dataBuffers, dataBuffer)
		case key == "recsindb":
			numMatched, err = strconv.Atoi(string(dataBytes))
			if err != nil {
				numMatched = 0
			}
		}
	}

	return CustomerSearchResult{
		RequestedLimit:  limit,
		RequestedOffset: offset,
		Customers:       dataBuffers,
		TotalMatched:    numMatched,
	}, nil
}

func (c *customer) Suspension(toggle bool, customerId, reason string) error {
	if !core.RgxNumber.MatchString(customerId) {
		return core.ErrRcInvalidCredential
	}

	funcName := "unsuspend"
	if toggle {
		funcName = "suspend"
	}

	data := url.Values{}
	data.Add("customer-id", customerId)
	data.Add("reason", reason)

	resp, err := c.core.CallApi(http.MethodPost, "customers", funcName, data)
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

func (c *customer) ForgotPassword(username string) error {
	if !core.RgxEmail.MatchString(username) {
		return core.ErrRcInvalidCredential
	}

	resp, err := c.core.CallApi(http.MethodGet, "customers", "forgot-password", url.Values{"username": {username}})
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

func (c *customer) SignUp(regForm *SignUpForm) error {
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
