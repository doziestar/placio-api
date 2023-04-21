package service

import (
	"errors"
	"gorm.io/gorm"
	"placio-app/models"
)

type CommentService interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	GetComment(commentID string) (*models.Comment, error)
	UpdateComment(comment *models.Comment) (*models.Comment, error)
	DeleteComment(commentID string) error
	ListComments(postID string) ([]*models.Comment, error)
}

type CommentServiceImpl struct {
	db           *gorm.DB
	store        *models.Comment
	accountStore *models.Account
}

func NewCommentService(db *gorm.DB) CommentService {
	return &CommentServiceImpl{db: db, store: &models.Comment{}, accountStore: &models.Account{}}
}

func (cs *CommentServiceImpl) CreateComment(comment *models.Comment) (*models.Comment, error) {
	if err := cs.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (cs *CommentServiceImpl) GetComment(commentID string) (*models.Comment, error) {
	var comment models.Comment
	if err := cs.db.First(&comment, "id = ?", commentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

func (cs *CommentServiceImpl) UpdateComment(comment *models.Comment) (*models.Comment, error) {
	if err := cs.db.Model(&models.Comment{}).Where("id = ?", comment.ID).Updates(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (cs *CommentServiceImpl) DeleteComment(commentID string) error {
	if err := cs.db.Delete(&models.Comment{}, "id = ?", commentID).Error; err != nil {
		return err
	}
	return nil
}

func (cs *CommentServiceImpl) ListComments(postID string) ([]*models.Comment, error) {
	var comments []*models.Comment
	if err := cs.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
