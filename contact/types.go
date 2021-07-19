package contact

import (
	"errors"
	"net/url"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ContactType string

// type ContactDetail struct {
// 	Id               string      `json:"entityid,omitempty" validate:"-" query:"-"`
// 	CustomerId       string      `json:"customerid,omitempty" validate:"required,number" query:"customer-id"`
// 	ParentKey        string      `json:"parentkey,omitempty" validate:"-" query:"-"`
// 	Name             string      `json:"name,omitempty" validate:"required" query:"name"`
// 	Company          string      `json:"company,omitempty" validate:"omitempty" query:"company"`
// 	Type             ContactType `json:"type,omitempty" validate:"omitempty" query:"type"`
// 	Email            string      `json:"emailaddr,omitempty" validate:"required,email" query:"email"`
// 	PhoneCountryCode string      `json:"telnocc,omitempty" validate:"omitempty,len=2,number" query:"phone-cc"`
// 	Phone            string      `json:"telno,omitempty" validate:"omitempty,number" query:"phone"`
// 	Address          string      `json:"address1,omitempty" validate:"omitempty" query:"address-line-1"`
// 	AddressLine2     string      `json:"address2,omitempty" validate:"omitempty" query:"address-line-2,omitempty"`
// 	AddressLine3     string      `json:"address3,omitempty" validate:"omitempty" query:"address-line-3,omitempty"`
// 	City             string      `json:"city,omitempty" validate:"omitempty" query:"city"`
// 	State            string      `json:"state,omitempty" validate:"omitempty" query:"state"`
// 	CountryCode      string      `json:"country,omitempty" validate:"omitempty,iso3166_1_alpha2" query:"country"`
// 	Zipcode          string      `json:"zip,omitempty" validate:"omitempty" query:"zipcode"`
// 	SystemStatus     string      `json:"currentstatus,omitempty" validate:"-" query:"-"`
// 	RegistryStatus   string      `json:"contactstatus,omitempty" validate:"-" query:"-"`
// }

type ContactDetail struct {
	// `query-add:"name" query-mod:"name" json:"name,omitempty" validate:"required"`
	Id               string      `json:"entityid" query-mod:"contact-id" query-add:"-"`
	Type             ContactType `json:"type" query-add:"type" query-mod:"-" validate:"required"`
	CustomerId       string      `json:"customerid" query-add:"customer-id" query-mod:"-" validate:"required,number"`
	StatusSystem     string      `json:"currentstatus" query-add:"-" query-mod:"-"`
	StatusRegistry   string      `json:"contactstatus" query-add:"-" query-mod:"-"`
	ParentKey        string      `json:"parentkey" query-add:"-" query-mod:"-"`
	Name             string      `json:"name" query-add:"name" query-mod:"name" validate:"required,max=255"`
	Email            string      `json:"emailaddr" query-add:"email" query-mod:"email" validate:"required,email"`
	Company          string      `json:"company" query-add:"company" query-mod:"company" validate:"required,max=255"`
	Address          string      `json:"address1" query-add:"address-line-1" query-mod:"address-line-1" validate:"required,max=64"`
	AddressLine2     string      `json:"address2" query-add:"address-line-2,optional" query-mod:"address-line-2,optional" validate:"-"`
	AddressLine3     string      `json:"address3" query-add:"address-line-3,optional" query-mod:"address-line-3,optional" validate:"-"`
	City             string      `json:"city" query-add:"city" query-mod:"city" validate:"required,max=64"`
	State            string      `json:"state" query-add:"state,optional" query-mod:"state,optional" validate:"omitempty,max=64"`
	CountryCode      string      `json:"country" query-add:"country" query-mod:"country" validate:"required,iso3166_1_alpha2"`
	Zipcode          string      `json:"zip" query-add:"zipcode" query-mod:"zipcode" validate:"required,max=16"`
	PhoneCountryCode string      `json:"telnocc" query-add:"phone-cc" query-mod:"phone-cc" validate:"required,min=1,max=3"`
	Phone            string      `json:"telno" query-add:"phone" query-mod:"phone" validate:"required,min=4,max=12"`
	FaxCountryCode   string      `json:"-" query-add:"fax-cc,optional" query-mod:"fax-cc,optional" validate:"omitempty,min=1,max=3"`
	Fax              string      `json:"-" query-add:"fax,optional" query-mod:"fax,optional" validate:"omitempty,min=4,max=12"`
}

const (
	TypeContact   ContactType = "Contact"
	TypeAt        ContactType = "AtContact"
	TypeBr        ContactType = "BrContact"
	TypeBrOrg     ContactType = "BrOrgContact"
	TypeCa        ContactType = "CaContact"
	TypeCl        ContactType = "ClContact"
	TypeCn        ContactType = "CnContact"
	TypeCo        ContactType = "CoContact"
	TypeCoop      ContactType = "CoopContact"
	TypeDe        ContactType = "DeContact"
	TypeEs        ContactType = "EsContact"
	TypeEu        ContactType = "EuContact"
	TypeFr        ContactType = "FrContact"
	TypeMx        ContactType = "MxContact"
	TypeNl        ContactType = "NlContact"
	TypeNyc       ContactType = "NycContact"
	TypeUk        ContactType = "UkContact"
	TypeUKService ContactType = "UkServiceContact"
)

func (c *ContactDetail) extract(queryTag string) (*url.Values, error) {
	if len(queryTag) <= 0 {
		return nil, errors.New("queryTag must not empty")
	}

	v := validator.New()
	if err := v.Struct(c); err != nil {
		return nil, err
	}

	valueCurrent := reflect.ValueOf(c)
	typeCurrent := reflect.TypeOf(c)

	ret := &url.Values{}
	for i := 0; i < valueCurrent.Elem().NumField(); i++ {
		vFieldCurrent := valueCurrent.Elem().Field(i)
		tFieldCurrent := typeCurrent.Elem().Field(i)
		tagFieldCurrent := tFieldCurrent.Tag.Get(queryTag)
		if len(tagFieldCurrent) <= 0 || tagFieldCurrent == "-" || vFieldCurrent.Kind() != reflect.String {
			continue
		}
		if vFieldCurrent.IsZero() {
			if !strings.HasSuffix(tagFieldCurrent, ",optional") {
				return nil, errors.New(strings.ToLower(tFieldCurrent.Name) + " must not empty")
			}
			continue
		}
		ret.Add(strings.TrimSuffix(tagFieldCurrent, ",optional"), vFieldCurrent.String())
	}

	return ret, nil
}
