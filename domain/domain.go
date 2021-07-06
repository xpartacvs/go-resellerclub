package domain

import (
	"net/http"

	"github.com/xpartacvs/go-resellerclub/core"
)

type domain struct {
	core core.Core
}

type Domain interface {
	GetRegistrationOrders(criteria SearchCriteria) error
}

func New(c core.Core) Domain {
	return &domain{
		core: c,
	}
}

func (d *domain) GetRegistrationOrders(criteria SearchCriteria) error {
	urlValues, err := criteria.UrlValues()
	if err != nil {
		return err
	}
	resp, err := d.core.CallApi(http.MethodGet, "domains", "search", urlValues)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
