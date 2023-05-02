package go_routines_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn/utils"
	"github.com/MickStanciu/go-fn/pkg/v1/fn/utils/go_routines"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParallelExecByKey_WhenAllArePassing(t *testing.T) {
	ctx := context.TODO()
	res, err := go_routines.ParallelExecByKey(
		ctx, 2,
		[]string{
			"key_1", "key_2", "key_3", "key_4", "key_5", "key_6",
			"key_7", "key_8", "key_9", "key_10", "key_11", "key_12",
		},
		func(id string) (*string, error) {
			result := fmt.Sprintf("%s_", id)
			return &result, nil
		},
	)
	require.NoError(t, err)
	expected := []*string{
		utils.ToPtr("key_1_"), utils.ToPtr("key_2_"), utils.ToPtr("key_3_"),
		utils.ToPtr("key_4_"), utils.ToPtr("key_5_"), utils.ToPtr("key_6_"),
		utils.ToPtr("key_7_"), utils.ToPtr("key_8_"), utils.ToPtr("key_9_"),
		utils.ToPtr("key_10_"), utils.ToPtr("key_11_"), utils.ToPtr("key_12_"),
	}
	assert.EqualValues(t, expected, res)
}

func TestParallelExecByKey_WhenSomeAreFailing(t *testing.T) {
	ctx := context.TODO()
	res, err := go_routines.ParallelExecByKey(
		ctx, 2,
		[]string{
			"key_1", "key_2", "key_3", "key_4", "key_5", "key_6",
			"key_7", "key_8", "key_9", "key_10", "key_11", "key_12",
		},
		func(id string) (*string, error) {
			if id == "key_5" {
				return nil, fmt.Errorf("some error with element %s", id)
			}
			result := fmt.Sprintf("%s_", id)
			return &result, nil
		},
	)
	require.EqualError(t, err, "some error with element key_5")
	assert.Nil(t, res)
}
