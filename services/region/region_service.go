package region

import "go-agreenery/entities"

type RegionService interface {
	GetProvinces() ([]entities.Region, error)
	GetRegencies(provinceCode string) ([]entities.Region, error)
	GetDistricts(regencyCode string) ([]entities.Region, error)
	GetVillages(districtCode string) ([]entities.Region, error)
}
