package domain

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xpartacvs/go-resellerclub/core"
)

var d = New(core.New(
	os.Getenv("RESELLER_ID"),
	os.Getenv("API_KEY"),
	false),
)

var (
	domainName = os.Getenv("TEST_DOMAIN_NAME")
	orderID    = os.Getenv("TEST_ORDER_ID")
	cns        = os.Getenv("TEST_CNS")
)

func TestSuggestNames(t *testing.T) {
	res, err := d.SuggestNames("domain", "", false, false)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestGetOrderID(t *testing.T) {
	res, err := d.GetOrderID(domainName)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestGetRegistrationOrderDetails(t *testing.T) {
	res, err := d.GetRegistrationOrderDetails(orderID, []string{"All"})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestModifyNameServers(t *testing.T) {
	res, err := d.ModifyNameServers(orderID, []string{"ns1.domain.asia"})
	require.NoError(t, err)
	require.NotNil(t, res)

	res, err = d.ModifyNameServers(orderID, []string{"ns2.domain.asia"})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestAddChildNameServer(t *testing.T) {
	res, err := d.AddChildNameServer(orderID, cns, []string{"0.0.0.0", "1.1.1.1"})
	require.NoError(t, err)
	require.NotNil(t, res)
}
