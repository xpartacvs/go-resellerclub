package general

import (
	"github.com/xpartacvs/go-resellerclub/core"
)

type general struct {
	core       core.Core
	currencies currencyDB
}

type General interface {
	CurrencyOf(iso CurrencyISO) Currency
}

func (g *general) CurrencyOf(iso CurrencyISO) Currency {
	return g.currencies[iso]
}

func New(c core.Core) (General, error) {
	curr, err := fetchCurrencyDB(c)
	if err != nil {
		return nil, err
	}
	return &general{
		core:       c,
		currencies: curr,
	}, nil
}
