package workshop

import (
	"github.com/MickStanciu/go-fn/workshop/api"
	"github.com/MickStanciu/go-fn/workshop/models"
)

func GetAllFlowers() []models.Flower {
	var flowers = make([]models.Flower, 0)
	batchSize := 5

	for {
		f := api.GetFlowers(len(flowers), len(flowers)+batchSize)
		if len(f) == 0 {
			break
		}
		flowers = append(flowers, f...)
	}

	return flowers
}

func GetAllFlowersR() []models.Flower {
	panic("not implemented")
}

func GetTheCheapFlowers() []models.Flower {
	var flowers = GetAllFlowers()

	var cheapFlowers = make([]models.Flower, 0)
	for _, f := range flowers {
		if f.Price < 10 {
			cheapFlowers = append(cheapFlowers, f)
		}
	}

	return cheapFlowers
}

func GetTheCheapFlowersR() []models.Flower {
	panic("not implemented")
}

func ConvertCheapFlowers() []models.FlowerResponse {
	var flowers = GetTheCheapFlowers()
	r := make([]models.FlowerResponse, len(flowers), len(flowers))

	for i, f := range flowers {
		copyF := f
		var photo *string = nil
		if copyF.Photo != "" {
			photo = &copyF.Photo
		}
		r[i] = models.FlowerResponse{
			Name:  copyF.Name,
			Price: copyF.Price,
			Photo: photo,
		}
	}

	return r
}

func ConvertCheapFlowersR() []models.FlowerResponse {
	panic("not implemented")
}
