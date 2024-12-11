package post_report

import "github.com/labstack/echo/v4"

type PostReportRepository interface {
	GetPostReports(c echo.Context) error
	CreatePostReport(c echo.Context) error
	DeletePostReport(c echo.Context) error
	SendWarning(c echo.Context) error
	DeletePostWithMessage(c echo.Context) error
}
