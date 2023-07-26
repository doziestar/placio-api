package Dto

import "mime/multipart"

type CreateCategoryRequest struct {
	Icon  string                  `json:"icon" form:"icon" binding:"required"`
	Name  string                  `json:"name" form:"name" binding:"required"`
	Image []*multipart.FileHeader `json:"image" form:"image" swaggertype:"file"`
}
