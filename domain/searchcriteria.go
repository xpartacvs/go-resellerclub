package domain

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

type SearchCriteria struct {
	Limit             uint16        `validate:"required,min=10,max=500" query:"no-of-records"`
	Offset            uint8         `validate:"required,min=1" query:"page-no"`
	SortOrderBy       []SortOrder   `validate:"omitempty" query:"order-by,omitempty"`
	OrderIDs          []uint        `validate:"omitempty" query:"order-id,omitempty"`
	ResellerIDs       []uint        `validate:"omitempty" query:"reseller-id,omitempty"`
	CustomerIDs       []uint        `validate:"omitempty" query:"customer-id,omitempty"`
	DomainKeys        []DomainKey   `validate:"omitempty" query:"product-key,omitempty"`
	DomainName        string        `validate:"omitempty" query:"domain-name,omitempty"`
	OrderStatuses     []OrderStatus `validate:"omitempty" query:"status,omitempty"`
	PrivacyStatus     PrivacyState  `validate:"omitempty" query:"privacy-enabled,omitempty"`
	ShowChildOrders   bool          `validate:"omitempty" query:"show-child-orders,omitempty"`
	TimeCreationStart time.Time     `validate:"omitempty" query:"creation-date-start,omitempty"`
	TimeCreationEnd   time.Time     `validate:"omitempty" query:"creation-date-end,omitempty"`
	TimeExpiryStart   time.Time     `validate:"omitempty" query:"expiry-date-start,omitempty"`
	TimeExpiryEnd     time.Time     `validate:"omitempty" query:"expiry-date-start,omitempty"`
}

func (c SearchCriteria) UrlValues() (url.Values, error) {
	if err := validator.New().Struct(c); err != nil {
		return url.Values{}, err
	}

	var wg sync.WaitGroup

	urlValues := url.Values{}
	valueCriteria := reflect.ValueOf(c)
	typeCriteria := reflect.TypeOf(c)

	for i := 0; i < valueCriteria.NumField(); i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			vField := valueCriteria.Field(idx)
			tField := typeCriteria.Field(idx)
			fieldTag := tField.Tag.Get("query")
			if len(fieldTag) > 0 {
				if strings.HasSuffix(fieldTag, "omitempty") && vField.IsZero() {
					return
				}
				queryField := strings.TrimSuffix(fieldTag, ",omitempty")

				switch vField.Kind() {
				case reflect.Uint8, reflect.Uint16, reflect.Uint:
					urlValues.Add(queryField, fmt.Sprintf("%d", vField.Uint()))
				case reflect.String:
					urlValues.Add(queryField, vField.String())
				case reflect.Bool:
					urlValues.Add(queryField, fmt.Sprintf("%t", vField.Bool()))
				case reflect.Struct:
					if vField.Type().ConvertibleTo(reflect.TypeOf(time.Time{})) {
						timeField := vField.Interface().(time.Time)
						urlValues.Add(queryField, fmt.Sprintf("%d", timeField.Unix()))
					}
				case reflect.Slice:
					for j := 0; j < vField.Len(); j++ {
						vSlice := vField.Index(j)
						switch vSlice.Type().Kind() {
						case reflect.Uint:
							urlValues.Add(queryField, fmt.Sprintf("%d", vSlice.Uint()))
						case reflect.String:
							urlValues.Add(queryField, vSlice.String())
						case reflect.Map:
							if vSlice.Type().ConvertibleTo(reflect.TypeOf(SortOrder{})) {
								vSortOrder := vSlice.Interface().(SortOrder)
								var wgSortOrder sync.WaitGroup
								for k, desc := range vSortOrder {
									wgSortOrder.Add(1)
									go func(key SortBy, value bool) {
										defer wgSortOrder.Done()
										vQuery := string(key)
										if value {
											vQuery += " desc"
										}
										urlValues.Add(queryField, vQuery)
									}(k, desc)
								}
								wgSortOrder.Wait()
							}
						}
					}
				}
			}

		}(i)
	}

	wg.Wait()
	return urlValues, nil
}
