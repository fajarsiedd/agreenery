package comment

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/services/comment"

	"github.com/labstack/echo/v4"
)

type commentHandler struct {
	service comment.CommentService
}

func NewCommentHandler(s comment.CommentService) *commentHandler {
	return &commentHandler{
		service: s,
	}
}

func (h commentHandler) GetComments(c echo.Context) error {
	postID := c.Param("postID")

	filter, err := helpers.GetFilter(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	filter.PostID = postID

	result, pagination, err := h.service.GetComments(filter)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponsePagination(c, constants.GetCommentsSuccess, pagination, response.ListCommentResponse{}.FromListEntity(result))
}

func (h commentHandler) CreateComment(c echo.Context) error {
	req := request.CommentRequest{}
	postID := c.Param("postID")

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	comment := req.ToEntity()
	comment.PostID = postID
	comment.UserID = claims.UserID

	result, err := h.service.CreateComment(comment)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.CreateCommentSuccess, response.CommentResponse{}.FromEntity(result))
}

func (h commentHandler) UpdateComment(c echo.Context) error {
	commentID := c.Param("id")
	postID := c.Param("postID")

	req := request.CommentRequest{ID: commentID}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	comment := req.ToEntity()
	comment.PostID = postID

	result, err := h.service.UpdateComment(comment, claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UpdateCommentSuccess, response.CommentResponse{}.FromEntity(result))
}

func (h commentHandler) DeleteComment(c echo.Context) error {
	commentID := c.Param("id")
	postID := c.Param("postID")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := h.service.DeleteComment(commentID, postID, claims.UserID); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeleteCommentSuccess, nil)
}
