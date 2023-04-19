package utils_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/MickStanciu/go-fn/pkg/v1/fn/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRetry_WhenAlwaysFails(t *testing.T) {
	_, err := utils.Retry[string](2, time.Second, func() (string, error) {
		return thisWillAlwaysFail("George")
	})

	require.NotNil(t, err)
	assert.Equal(t, "failing internally", err.Error())
}

func TestRetry_WhenAlwaysFails_Pointer(t *testing.T) {
	_, err := utils.Retry[*string](2, time.Second, func() (*string, error) {
		res, err := thisWillAlwaysFail("George")
		return &res, err
	})

	require.NotNil(t, err)
	assert.Equal(t, "failing internally", err.Error())
}

func TestRetry_WhenAlwaysSucceeds(t *testing.T) {
	res, err := utils.Retry[string](2, time.Second, func() (string, error) {
		return thisWillAlwaysSucceed("George")
	})

	require.NoError(t, err)
	assert.Equal(t, "Hello George", res)
}

func thisWillAlwaysFail(arg string) (string, error) {
	return "", fmt.Errorf("failing internally")
}

func thisWillAlwaysSucceed(arg string) (string, error) {
	return "Hello " + arg, nil
}
