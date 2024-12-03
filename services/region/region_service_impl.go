package region

import (
	"encoding/json"
	"go-agreenery/entities"
	"net/http"
)

type regionService struct{}

type responseBody struct {
	Data []entities.Region
}

func NewRegionService() *regionService {
	return &regionService{}
}

func (service regionService) GetProvinces() ([]entities.Region, error) {
	var err error
	var client = &http.Client{}

	request, err := http.NewRequest("GET", "https://wilayah.id/api/provinces.json", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res responseBody
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (service regionService) GetRegencies(provinceCode string) ([]entities.Region, error) {
	var err error
	var client = &http.Client{}

	request, err := http.NewRequest("GET", "https://wilayah.id/api/regencies/"+provinceCode+".json", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res responseBody
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (service regionService) GetDistricts(regencyCode string) ([]entities.Region, error) {
	var err error
	var client = &http.Client{}

	request, err := http.NewRequest("GET", "https://wilayah.id/api/districts/"+regencyCode+".json", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res responseBody
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (service regionService) GetVillages(districtCode string) ([]entities.Region, error) {
	var err error
	var client = &http.Client{}

	request, err := http.NewRequest("GET", "https://wilayah.id/api/villages/"+districtCode+".json", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res responseBody
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
