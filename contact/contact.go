package contact

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/xpartacvs/go-resellerclub/core"
)

type contact struct {
	core core.Core
}

type Contact interface {
	Add(detail *ContactDetail, attributes core.EntityAttributes) error
	Details(contactId string) (*ContactDetail, error)
}

func (c *contact) Details(contactId string) (*ContactDetail, error) {
	if !core.RgxNumber.MatchString(contactId) {
		return nil, core.ErrRcInvalidCredential
	}

	data := url.Values{}
	data.Add("contact-id", contactId)

	resp, err := c.core.CallApi(http.MethodGet, "contacts", "details", data)
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

	ret := new(ContactDetail)
	if err = json.Unmarshal(bytesResp, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *contact) Add(details *ContactDetail, attributes core.EntityAttributes) error {
	if details == nil {
		return errors.New("detail must not nil")
	}

	data, err := details.UrlValues()
	if err != nil {
		return err
	}

	if attributes != nil {
		attributes.CopyTo(data)
	}

	resp, err := c.core.CallApi(http.MethodPost, "contacts", "add", *data)
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

	details.Id = string(bytesResp)
	return nil
}

func New(c core.Core) Contact {
	return &contact{
		core: c,
	}
}
