package contact

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/xpartacvs/go-resellerclub/core"
)

type contact struct {
	core core.Core
}

type Contact interface {
	Add(detail *ContactDetail, attributes core.EntityAttributes) error
}

func (c *contact) Add(details *ContactDetail, attributes core.EntityAttributes) error {
	if details == nil {
		return errors.New("detail must not nil")
	}

	data, err := details.extract("query-add")
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
