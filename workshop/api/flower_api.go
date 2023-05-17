package api

import (
	"encoding/json"
	"os"
	"time"

	"github.com/MickStanciu/go-fn/workshop/models"
)

var flowersMap = make(map[int]models.Flower)
var flowers = make([]models.Flower, 0)

func init() {
	b, err := os.ReadFile("testdata/flowers.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(b, &flowers); err != nil {
		panic(err)
	}

	for _, f := range flowers {
		flowersMap[f.ProductId] = f
	}
}

func GetFlower(id int) *models.Flower {
	v, ok := flowersMap[id]
	if !ok {
		return nil
	}
	return &v
}

func GetFlowers(from, to int) []models.Flower {
	if from < 0 || from > len(flowers) || to < 0 || from > to {
		return nil
	}

	if to > len(flowers) {
		to = len(flowers)
	}

	time.Sleep(200 * time.Millisecond)
	return flowers[from:to]
}
