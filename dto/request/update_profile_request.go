package request

import (
	"encoding/json"
	"go-agreenery/entities"
)

type UpdateProfileRequest struct {
	ID          string
	DisplayName string `json:"display_name,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty" validate:"email"`
	Photo       string `json:"photo,omitempty"`
}

func (updateProfileRequest UpdateProfileRequest) ToEntity() entities.User {
	return entities.User{
		Base: entities.Base{
			ID: updateProfileRequest.ID,
		},
		DisplayName: updateProfileRequest.DisplayName,
		Phone:       updateProfileRequest.Phone,
		Photo:       updateProfileRequest.Photo,
		Credential: entities.Credential{
			Email: updateProfileRequest.Email,
		},
	}
}

func (updateProfileRequest UpdateProfileRequest) ToCleanFields() []string {
	var fields []string
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(&updateProfileRequest)
	json.Unmarshal(inrec, &inInterface)

	for field, val := range inInterface {
		if val == nil {
			continue
		}

		fields = append(fields, field)
	}

	return fields
}
