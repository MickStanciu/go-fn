package workshop_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/MickStanciu/go-fn/workshop"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllFlowers(t *testing.T) {
	r := workshop.GetAllFlowers()
	assert.Len(t, r, 34)
}

func TestGetAllFlowersR(t *testing.T) {
	r := workshop.GetAllFlowersR()
	assert.Len(t, r, 34)
}

//func TestQueryAPI(t *testing.T) {
//	tfn := func(from, to int) []string {
//		fmt.Println(from, to)
//		if to == 5 {
//			return []string{"a", "b", "c", "d", "e"}
//		}
//		return nil
//	}
//	r := workshop.QueryAPI(5, tfn)
//	require.Len(t, r, 5)
//}

func TestGetTheCheapFlowers(t *testing.T) {
	r := workshop.GetTheCheapFlowers()
	require.Len(t, r, 14)
	for _, flower := range r {
		assert.True(t, flower.Price < 10)
	}
}
func TestGetTheCheapFlowersR(t *testing.T) {
	r := workshop.GetTheCheapFlowersR()
	require.Len(t, r, 14)
	for _, flower := range r {
		assert.True(t, flower.Price < 10)
	}
}

func TestMapCheapFlowers(t *testing.T) {
	r := workshop.ConvertCheapFlowers()
	require.Len(t, r, 14)
	for _, flower := range r {
		assert.True(t, flower.Price < 10)
	}

	jsonResponse, err := json.Marshal(r)
	require.NoError(t, err)

	b, err := os.ReadFile("testdata/expected_response.json")
	if err != nil {
		panic(err)
	}

	require.JSONEq(t, string(b), string(jsonResponse))
}

func TestMapCheapFlowersR(t *testing.T) {
	r := workshop.ConvertCheapFlowersR()
	require.Len(t, r, 14)
	for _, flower := range r {
		assert.True(t, flower.Price < 10)
	}

	jsonResponse, err := json.Marshal(r)
	require.NoError(t, err)

	b, err := os.ReadFile("testdata/expected_response.json")
	if err != nil {
		panic(err)
	}

	require.JSONEq(t, string(b), string(jsonResponse))
}
