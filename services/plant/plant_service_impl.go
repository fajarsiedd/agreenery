package plant

import (
	"go-agreenery/entities"
	"go-agreenery/helpers"
	"go-agreenery/repositories/plant"
	"strings"
)

type plantService struct {
	repository plant.PlantRepository
}

func NewPlantService(r plant.PlantRepository) *plantService {
	return &plantService{
		repository: r,
	}
}

func (s plantService) GetPlants(filter entities.Filter) ([]entities.Plant, entities.Pagination, error) {
	return s.repository.GetPlants(filter)
}

func (s plantService) GetPlant(id string) (entities.Plant, error) {
	return s.repository.GetPlant(id)
}

func (s plantService) CreatePlant(plant entities.Plant) (entities.Plant, error) {
	var url string
	if plant.ImageFile != nil {
		params := helpers.UploaderParams{
			File: plant.ImageFile,
		}

		result, err := helpers.UploadFile(params)
		if err != nil {
			return entities.Plant{}, err
		}

		url = result
	}

	plant.Image = url

	// GET RECOMMENDATION FROM AI
	fertilizer, tips := helpers.GetFertilzerAndPlantingRecommendation(plant.Name)

	plant.Fertilizer = fertilizer
	plant.PlantingTips = tips

	result, err := s.repository.CreatePlant(plant)
	if err != nil {
		if plant.ImageFile != nil {
			splittedStr := strings.Split(url, "/")
			object := splittedStr[len(splittedStr)-1]

			if err := helpers.DeleteFile(object); err != nil {
				return entities.Plant{}, err
			}
		}

		return entities.Plant{}, err
	}

	return result, nil
}

func (s plantService) UpdatePlant(plant entities.Plant) (entities.Plant, error) {
	plantDb, err := s.repository.GetPlant(plant.ID)
	if err != nil {
		return entities.Plant{}, err
	}

	var url string
	if plant.ImageFile != nil {
		params := helpers.UploaderParams{
			File: plant.ImageFile,
		}

		result, err := helpers.UploadFile(params)
		if err != nil {
			return entities.Plant{}, err
		}

		url = result
	}

	plant.Image = url

	// GET RECOMMENDATION FROM AI
	fertilizer, tips := helpers.GetFertilzerAndPlantingRecommendation(plant.Name)

	plant.Fertilizer = fertilizer
	plant.PlantingTips = tips

	result, err := s.repository.UpdatePlant(plant)
	if err != nil {
		if plant.ImageFile != nil {
			splittedStr := strings.Split(url, "/")
			object := splittedStr[len(splittedStr)-1]

			if err := helpers.DeleteFile(object); err != nil {
				return entities.Plant{}, err
			}
		}

		return entities.Plant{}, err
	} else {
		if plant.ImageFile != nil && plantDb.Image != "" {
			splittedStr := strings.Split(plantDb.Image, "/")
			oldObj := splittedStr[len(splittedStr)-1]

			if err := helpers.DeleteFile(oldObj); err != nil {
				return entities.Plant{}, err
			}
		}
	}

	return result, nil
}

func (s plantService) DeletePlant(id string) error {
	media, err := s.repository.DeletePlant(id)
	if err != nil {
		return err
	}

	if media != "" {
		splittedStr := strings.Split(media, "/")
		oldObj := splittedStr[len(splittedStr)-1]

		if err := helpers.DeleteFile(oldObj); err != nil {
			return err
		}
	}

	return nil
}
