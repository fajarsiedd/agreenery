package auth

import (
	"go-agreenery/constants"
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (repository *authRepository) Login(user entities.User) (entities.User, error) {
	userModel := models.User{}.FromEntity(user)

	err := repository.db.First(&userModel.Credential, "email = ?", userModel.Credential.Email).Error
	if err != nil {
		return entities.User{}, constants.ErrIncorrectEmail
	}

	err = repository.db.First(&userModel, "credential_id = ?", userModel.Credential.Base.ID).Error
	if err != nil {
		return entities.User{}, constants.ErrUserNotFound
	}

	return userModel.ToEntity(), nil
}

func (repository authRepository) Register(user entities.User) (entities.User, error) {
	userModel := models.User{}.FromEntity(user)

	err := repository.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userModel.Credential).Error; err != nil {
			if err == gorm.ErrDuplicatedKey {
				return constants.ErrDuplicateEmail
			}

			return err
		}

		userModel.CredentialID = userModel.Credential.ID

		if err := tx.Omit("Credential").Create(&userModel).Error; err != nil {
			if err == gorm.ErrDuplicatedKey {
				return constants.ErrDuplicatePhone
			}

			return err
		}

		return nil
	})
	if err != nil {
		return entities.User{}, err
	}

	return userModel.ToEntity(), nil
}

func (repository authRepository) FindUser(id string) (entities.User, error) {
	userModel := models.User{}

	if err := repository.db.Preload("Credential").First(&userModel, &id).Error; err != nil {
		return entities.User{}, constants.ErrUserNotFound
	}

	return userModel.ToEntity(), nil
}

func (repository authRepository) UpdateUser(user entities.User, selectedFields []string) (entities.User, error) {
	userModel := models.User{}.FromEntity(user)

	err := repository.db.Transaction(func(tx *gorm.DB) error {
		userData := models.User{}
		if err := tx.Preload("Credential").First(&userData, &user.ID).Error; err != nil {
			return constants.ErrUserNotFound
		}

		if err := tx.Select(selectedFields).Where("id = ?", userData.CredentialID).Updates(&userModel.Credential).Error; err != nil {
			if err == gorm.ErrDuplicatedKey {
				return constants.ErrDuplicateEmail
			}

			return err
		}

		if err := tx.Select(selectedFields).Where("id = ?", userData.ID).Omit("Credential").Updates(&userModel).Error; err != nil {
			if err == gorm.ErrDuplicatedKey {
				return constants.ErrDuplicatePhone
			}

			return err
		}

		if err := tx.Preload("Credential").First(&userModel, &user.ID).Error; err != nil {
			return constants.ErrUserNotFound
		}

		return nil
	})
	if err != nil {
		return entities.User{}, err
	}

	return userModel.ToEntity(), nil
}
