package strings_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn/utils/strings"
	"github.com/stretchr/testify/assert"
)

func TestGetOrElse(t *testing.T) {
	assert.EqualValues(t, "Z", strings.GetOrElse("Z", "should not happen"))
	assert.EqualValues(t, "Z", strings.GetOrElse("", "Z"))
}
