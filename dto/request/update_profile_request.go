package request

import (
	"encoding/json"
	"go-agreenery/entities"
)

type UpdateProfileRequest struct {
	ID          string `json:"-"`
	DisplayName string `json:"display_name,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty" validate:"email"`
	Photo       string `json:"photo,omitempty"`
}

func (r UpdateProfileRequest) ToEntity() entities.User {
	return entities.User{
		Base: entities.Base{
			ID: r.ID,
		},
		DisplayName: r.DisplayName,
		Phone:       r.Phone,
		Photo:       r.Photo,
		Credential: entities.Credential{
			Email: r.Email,
		},
	}
}

func (r UpdateProfileRequest) ToCleanFields() []string {
	var fields []string
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(&r)
	json.Unmarshal(inrec, &inInterface)

	for field, val := range inInterface {
		if val == nil {
			continue
		}

		fields = append(fields, field)
	}

	return fields
}
