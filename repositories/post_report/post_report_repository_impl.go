package post_report

import (
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type postReportRepository struct {
	db *gorm.DB
}

func NewPostReportRepository(db *gorm.DB) *postReportRepository {
	return &postReportRepository{
		db: db,
	}
}

func (r postReportRepository) GetPostReports(filter entities.Filter) ([]entities.PostReport, entities.Pagination, error) {
	postReportModel := models.ListPostReport{}

	query := r.db.Model(&postReportModel)

	if filter.Search != "" {
		query = query.Where("report_type LIKE ?", "%"+filter.Search+"%")
	}

	if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
		query = query.Where("post_reports.created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}

	query = query.Order("post_reports." + filter.SortBy + " " + filter.Sort)

	var totalItems int64

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	offset := (filter.Page - 1) * filter.Limit

	if err := query.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Limit(filter.Limit).Offset(offset).Find(&postReportModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return postReportModel.ToListEntity(), pagination, nil
}

func (r postReportRepository) CreatePostReport(postReport entities.PostReport) (entities.PostReport, error) {
	postReportModel := models.PostReport{}.FromEntity(postReport)

	if err := r.db.Omit("User").Create(&postReportModel).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Find(&postReportModel).Error; err != nil {
		return entities.PostReport{}, err
	}

	return postReportModel.ToEntity(), nil
}

func (r postReportRepository) DeletePostReport(id string) error {
	postReportModel := models.PostReport{}

	if err := r.db.Unscoped().Delete(&postReportModel, &id).Error; err != nil {
		return err
	}

	return nil
}

func (r postReportRepository) DeletePostWithMessage(postReport entities.PostReport) (string, error) {
	var media string
	err := r.db.Transaction(func(tx *gorm.DB) error {
		postReportModel := models.PostReport{}
		if err := r.db.First(&postReportModel, &postReport.ID).Error; err != nil {
			return err
		}

		postDb := models.Post{}
		if err := tx.First(&postDb, &postReportModel.PostID).Error; err != nil {
			return err
		}

		media = postDb.Media

		if err := tx.Unscoped().Where("post_id = ?", postReportModel.PostID).Delete(&models.Comment{}).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Where("post_id = ?", postReportModel.PostID).Delete(&models.Like{}).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Delete(&postDb).Error; err != nil {
			return err
		}

		// CREATE NOTIFICATION
		userNotifModel := models.UserNotification{
			UserID:   postDb.UserID,
			Title:    "Pelanggaran Aturan: Postingan Anda telah dihapus",
			Subtitle: postReport.Message,
			Icon:     "https://storage.googleapis.com/agreenery/uploads/agreenery-logo.png",
		}
		if err := tx.Create(&userNotifModel).Error; err != nil {
			return err
		}

		if err := tx.Model(&postReportModel).Where("id = ?", postReport.ID).Update("status_done", true).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return media, nil
}

func (r postReportRepository) SendWarning(postReport entities.PostReport) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		postReportModel := models.PostReport{}
		if err := r.db.First(&postReportModel, &postReport.ID).Error; err != nil {
			return err
		}

		postDb := models.Post{}
		if err := tx.First(&postDb, &postReportModel.PostID).Error; err != nil {
			return err
		}

		// CREATE NOTIFICATION
		userNotifModel := models.UserNotification{
			UserID:   postDb.UserID,
			Title:    "Peringatan: " + postReportModel.ReportType,
			Subtitle: postReport.Message,
			Icon:     "https://storage.googleapis.com/agreenery/uploads/agreenery-logo.png",
		}

		if err := tx.Create(&userNotifModel).Error; err != nil {
			return err
		}

		if err := tx.Model(&postReportModel).Where("id = ?", postReport.ID).Update("status_done", true).Error; err != nil {
			return err
		}

		return nil
	})
}
