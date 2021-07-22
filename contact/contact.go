package contact

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/xpartacvs/go-resellerclub/core"
)

type contact struct {
	core core.Core
}

type Contact interface {
	Add(detail *ContactDetail, attributes core.EntityAttributes) error
	Details(contactId string) (*ContactDetail, error)
	Delete(contactId string) (*Action, error)
	Search(criteria ContactCriteria, offset, limit uint16) (*ContactSearchResult, error)
}

func (c *contact) Search(criteria ContactCriteria, offset, limit uint16) (*ContactSearchResult, error) {
	if offset <= 0 || limit <= 0 {
		return nil, errors.New("offset or limit must greater than zero")
	}

	if err := validator.New().Struct(criteria); err != nil {
		return nil, err
	}

	data, err := criteria.UrlValues()
	if err != nil {
		return nil, err
	}
	data.Add("no-of-records", strconv.FormatUint(uint64(limit), 10))
	data.Add("page-no", strconv.FormatUint(uint64(offset), 10))

	resp, err := c.core.CallApi(http.MethodGet, "contacts", "search", data)
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

	replacer := strings.NewReplacer("entity.", "", "contact.", "")
	strResp := replacer.Replace(string(bytesResp))

	var buffer map[string]core.JSONBytes
	if err := json.Unmarshal([]byte(strResp), &buffer); err != nil {
		return nil, err
	}

	var dataBuffers []ContactDetail
	var numMatched int

	for key, dataBytes := range buffer {
		switch {
		case key == "recsindb":
			numMatched, err = strconv.Atoi(string(dataBytes))
			if err != nil {
				numMatched = 0
			}
		case key == "result":
			if err := json.Unmarshal(dataBytes, &dataBuffers); err != nil {
				return nil, err
			}
		}
	}

	return &ContactSearchResult{
		RequestedLimit:  limit,
		RequestedOffset: offset,
		Contacts:        dataBuffers,
		TotalMatched:    numMatched,
	}, nil
}

func (c *contact) Delete(contactId string) (*Action, error) {
	if !core.RgxNumber.MatchString(contactId) {
		return nil, core.ErrRcInvalidCredential
	}

	resp, err := c.core.CallApi(http.MethodPost, "contacts", "delete", url.Values{"contact-id": {contactId}})
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

	ret := new(Action)
	if err = json.Unmarshal(bytesResp, ret); err != nil {
		return nil, err
	}

	return ret, nil
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
