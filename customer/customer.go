package customer

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
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
}

func (c *customer) Details(customerIdOrEmail string) (CustomerDetail, error) {
	ret := CustomerDetail{}
	data := url.Values{}
	rgxEmail := regexp.MustCompile("^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
	rgxNumber := regexp.MustCompile("^[0-9]+$")

	var funcName, query string
	switch {
	case rgxEmail.MatchString(customerIdOrEmail):
		funcName = "details"
		query = "username"
	case rgxNumber.MatchString(customerIdOrEmail):
		funcName = "details-by-id"
		query = "customer-id"
	default:
		return ret, errors.New("invalid username or email")
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
		return errors.New("password unchanged")
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
