package watering_schedule

import (
	"go-agreenery/entities"
	"go-agreenery/helpers"
	"go-agreenery/repositories/watering_schedule"
	"strings"
)

type wateringScheduleService struct {
	repository watering_schedule.WateringScheduleRepository
}

func NewWateringScheduleService(r watering_schedule.WateringScheduleRepository) *wateringScheduleService {
	return &wateringScheduleService{
		repository: r,
	}
}

func (s wateringScheduleService) GetWateringSchedules(filter entities.Filter) ([]entities.WateringSchedule, entities.Pagination, error) {
	return s.repository.GetWateringSchedules(filter)
}

func (s wateringScheduleService) GetWateringSchedule(id string) (entities.WateringSchedule, error) {
	return s.repository.GetWateringSchedule(id)
}

func (s wateringScheduleService) CreateWateringSchedule(schedule entities.WateringSchedule) (entities.WateringSchedule, error) {
	var url string
	if schedule.ImageFile != nil {
		params := helpers.UploaderParams{
			File: schedule.ImageFile,
		}

		result, err := helpers.UploadFile(params)
		if err != nil {
			return entities.WateringSchedule{}, err
		}

		url = result
	}

	schedule.Image = url

	result, err := s.repository.CreateWateringSchedule(schedule)
	if err != nil {
		if schedule.ImageFile != nil {
			splittedStr := strings.Split(url, "/")
			object := splittedStr[len(splittedStr)-1]

			if err := helpers.DeleteFile(object); err != nil {
				return entities.WateringSchedule{}, err
			}
		}

		return entities.WateringSchedule{}, err
	}

	return result, nil
}

func (s wateringScheduleService) UpdateWateringSchedule(schedule entities.WateringSchedule, currUserID string) (entities.WateringSchedule, error) {
	scheduleDb, err := s.repository.GetWateringSchedule(schedule.ID)
	if err != nil {
		return entities.WateringSchedule{}, err
	}

	var url string
	if schedule.ImageFile != nil {
		params := helpers.UploaderParams{
			File: schedule.ImageFile,
		}

		result, err := helpers.UploadFile(params)
		if err != nil {
			return entities.WateringSchedule{}, err
		}

		url = result
	}

	schedule.Image = url

	result, err := s.repository.UpdateWateringSchedule(schedule, currUserID)
	if err != nil {
		if schedule.ImageFile != nil {
			splittedStr := strings.Split(url, "/")
			object := splittedStr[len(splittedStr)-1]

			if err := helpers.DeleteFile(object); err != nil {
				return entities.WateringSchedule{}, err
			}
		}

		return entities.WateringSchedule{}, err
	} else {
		if schedule.ImageFile != nil && scheduleDb.Image != "" {
			splittedStr := strings.Split(scheduleDb.Image, "/")
			oldObj := splittedStr[len(splittedStr)-1]

			if err := helpers.DeleteFile(oldObj); err != nil {
				return entities.WateringSchedule{}, err
			}
		}
	}

	return result, nil
}

func (s wateringScheduleService) DeleteWateringSchedule(id, currUserID string) error {
	media, err := s.repository.DeleteWateringSchedule(id, currUserID)
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
