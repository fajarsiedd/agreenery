package response

import "go-agreenery/entities"

type RegionResponse struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	PostalCode string `json:"postal_code,omitempty"`
}

type ListRegionResponse []RegionResponse

func (regionResponse RegionResponse) FromEntity(regionEntity entities.Region) RegionResponse {
	return RegionResponse{
		Code:       regionEntity.Code,
		Name:       regionEntity.Name,
		PostalCode: regionEntity.PostalCode,
	}
}

func (listRegionResponse ListRegionResponse) FromListEntity(regionEntities []entities.Region) ListRegionResponse {
	regions := []RegionResponse{}

	for _, v := range regionEntities {
		regions = append(regions, RegionResponse{}.FromEntity(v))
	}

	return regions
}
