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
	return QueryAPI(5, api.GetFlowers)
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
	return Filter(GetAllFlowersR(), func(f models.Flower) bool {
		return f.Price < 10
	})
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
	return Map(GetTheCheapFlowersR(), func(f models.Flower) models.FlowerResponse {
		return models.FlowerResponse{
			Name:  f.Name,
			Price: f.Price,
			Photo: GetOrElse(&f.Photo, nil, func(p *string) bool {
				return *p != ""
			}),
		}
	})
}

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result = make([]T, 0)
	for _, i := range items {
		if predicate(i) {
			result = append(result, i)
		}
	}
	return result
}

func Map[A, B any](items []A, mapper func(A) B) []B {
	var result = make([]B, len(items), len(items))
	for i, item := range items {
		result[i] = mapper(item)
	}
	return result
}

func GetOrElse[T any](item T, orElse T, fn func(T) bool) T {
	if fn(item) {
		return item
	}
	return orElse
}

func QueryAPI[T any](batchSize int, procFn func(from, to int) []T) []T {
	var result = make([]T, 0)

	for {
		f := procFn(len(result), len(result)+batchSize)
		if len(f) == 0 {
			break
		}
		result = append(result, f...)
	}

	return result
}
