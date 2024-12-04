package helpers

import (
	"context"
	"go-agreenery/constants"
	"os"

	"googlemaps.github.io/maps"
)

func GetCoordinates(address string) (float64, float64, error) {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	if err != nil {
		return 0, 0, err
	}

	req := &maps.GeocodingRequest{
		Address: address,
	}

	result, err := c.Geocode(context.Background(), req)
	if err != nil {
		return 0, 0, err
	}

	if len(result) > 0 {
		lat := result[0].Geometry.Location.Lat
		lng := result[0].Geometry.Location.Lng
		return lat, lng, nil
	}

	return 0, 0, constants.ErrGetCoordinatesFailed
}
