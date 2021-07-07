package customer

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/xpartacvs/go-resellerclub/core"
)

type customer struct {
	core core.Core
}

type Customer interface {
	SignUp(regForm *RegistrationForm) error
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
