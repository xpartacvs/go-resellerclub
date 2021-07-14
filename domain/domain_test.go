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

func TestSuggestNames(t *testing.T) {
	res, err := d.SuggestNames("domain", "", false, false)
	require.NoError(t, err)
	require.NotNil(t, res)
}
