package batch_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/MickStanciu/go-fn/batch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParallelExecByKey_WhenAllArePassing(t *testing.T) {
	ctx := context.TODO()
	res, err := batch.ParallelExecByKey(
		ctx, 2,
		[]string{
			"item_1", "item_2", "item_3", "item_4", "item_5", "item_6",
			"item_7", "item_8", "item_9", "item_10", "item_11", "item_12",
		},
		func(ctx context.Context, itemID string) (string, error) {
			return fmt.Sprintf("_%s_", itemID), nil
		},
	)
	require.NoError(t, err)
	expected := map[string]string{
		"item_1": "_item_1_", "item_2": "_item_2_", "item_3": "_item_3_", "item_4": "_item_4_",
		"item_5": "_item_5_", "item_6": "_item_6_", "item_7": "_item_7_", "item_8": "_item_8_",
		"item_9": "_item_9_", "item_10": "_item_10_", "item_11": "_item_11_", "item_12": "_item_12_",
	}
	assert.EqualValues(t, expected, res)
}

func TestParallelExecByKey_WhenSomeAreFailing(t *testing.T) {
	ctx := context.TODO()
	res, err := batch.ParallelExecByKey(
		ctx, 2,
		[]string{
			"item_1", "item_2", "item_3", "item_4", "item_5", "item_6",
			"item_7", "item_8", "item_9", "item_10", "item_11", "item_12",
		},
		func(ctx context.Context, itemID string) (string, error) {
			if itemID == "item_5" {
				return "", fmt.Errorf("some error with element %s", itemID)
			}
			return fmt.Sprintf("%s_", itemID), nil
		},
	)
	require.EqualError(t, err, "some error with element item_5")
	assert.Nil(t, res)
}
