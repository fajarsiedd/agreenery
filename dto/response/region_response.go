package response

import "go-agreenery/entities"

type RegionResponse struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	PostalCode string `json:"postal_code,omitempty"`
}

type ListRegionResponse []RegionResponse

func (r RegionResponse) FromEntity(region entities.Region) RegionResponse {
	return RegionResponse{
		Code:       region.Code,
		Name:       region.Name,
		PostalCode: region.PostalCode,
	}
}

func (r ListRegionResponse) FromListEntity(regions []entities.Region) ListRegionResponse {
	data := ListRegionResponse{}

	for _, v := range regions {
		data = append(data, RegionResponse{}.FromEntity(v))
	}

	return data
}
