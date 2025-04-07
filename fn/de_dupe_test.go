package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/fn"
	"github.com/stretchr/testify/assert"
)

func TestDeduplicateList(t *testing.T) {
	type human struct {
		ID   string
		Name string
		Age  int
	}

	tests := []struct {
		name         string
		input        []*human
		expect       []*human
		customAssert func(t *testing.T, output []*human)
	}{
		{
			name:   "when empty",
			input:  []*human{},
			expect: []*human{},
			customAssert: func(t *testing.T, output []*human) {
				assert.Empty(t, output)
			},
		},
		{
			name: "when there no duplicates",
			input: []*human{
				{
					ID:   "1",
					Name: "George",
					Age:  10,
				},
				{
					ID:   "2",
					Name: "Michael",
					Age:  20,
				},
			},
			expect: []*human{
				{
					ID:   "1",
					Name: "George",
					Age:  10,
				},
				{
					ID:   "2",
					Name: "Michael",
					Age:  20,
				},
			},
			customAssert: func(t *testing.T, output []*human) {
				assert.Len(t, output, 2)
			},
		},
		{
			name: "when there are duplicates",
			input: []*human{
				{
					ID:   "1",
					Name: "George",
					Age:  10,
				},
				{
					ID:   "2",
					Name: "Michael",
					Age:  20,
				},
				{
					ID:   "1",
					Name: "Thomas",
					Age:  30,
				},
				{
					ID:   "1",
					Name: "John",
					Age:  40,
				},
			},
			expect: []*human{
				{
					ID:   "1",
					Name: "John",
					Age:  40,
				},
				{
					ID:   "2",
					Name: "Michael",
					Age:  20,
				},
			},
			customAssert: func(t *testing.T, output []*human) {
				assert.Len(t, output, 2)
				assert.Equal(t, "1", output[0].ID)
				assert.Equal(t, "John", output[0].Name)
				assert.Equal(t, 40, output[0].Age)
				assert.Equal(t, "2", output[1].ID)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fn.DeduplicateList(tt.input, func(element *human) string {
				return element.ID
			})
			tt.customAssert(t, result)
		})
	}
}
