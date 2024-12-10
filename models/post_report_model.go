package models

import "go-agreenery/entities"

type PostReport struct {
	Base
	UserID     string `gorm:"size:191"`
	User       User   `gorm:"foreignKey:UserID;references:ID"`
	PostID     string `gorm:"size:191"`
	ReportType string `gorm:"size:255"`
	StatusDone bool   `gorm:"default:false"`
}

type ListPostReport []PostReport

func (p PostReport) FromEntity(report entities.PostReport) PostReport {
	return PostReport{
		Base:       p.Base.FromEntity(report.Base),
		UserID:     report.UserID,
		User:       p.User.FromEntity(report.User),
		PostID:     report.PostID,
		ReportType: report.ReportType,
		StatusDone: report.StatusDone,
	}
}

func (p PostReport) ToEntity() entities.PostReport {
	return entities.PostReport{
		Base:       p.Base.ToEntity(),
		UserID:     p.UserID,
		User:       p.User.ToEntity(),
		PostID:     p.PostID,
		ReportType: p.ReportType,
		StatusDone: p.StatusDone,
	}
}

func (lp ListPostReport) FromListEntity(categories []entities.PostReport) ListPostReport {
	data := ListPostReport{}

	for _, v := range categories {
		data = append(data, PostReport{}.FromEntity(v))
	}

	return data
}

func (lp ListPostReport) ToListEntity() []entities.PostReport {
	data := []entities.PostReport{}

	for _, v := range lp {
		data = append(data, v.ToEntity())
	}

	return data
}
