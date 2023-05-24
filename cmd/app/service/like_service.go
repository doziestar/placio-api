package service

import (
	"errors"
	"placio-app/models"

	"gorm.io/gorm"
)

type LikeService interface {
	CreateLike(like *models.Like) (*models.Like, error)
	GetLike(likeID string) (*models.Like, error)
	DeleteLike(likeID string) error
	ListLikes(postID string) ([]*models.Like, error)
	LikePost(postID string, accountID string) error
	DeleteLikeByPostID(postID string) error
	UnlikePost(postID string, accountID string) error
	GetLikeCount(postID string) (int64, error)
}

type likeServiceImpl struct {
	db *gorm.DB
}

func NewLikeService(db *gorm.DB) LikeService {
	return &likeServiceImpl{db: db}
}

func (ls *likeServiceImpl) CreateLike(like *models.Like) (*models.Like, error) {
	if err := ls.db.Create(&like).Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (ls *likeServiceImpl) GetLike(likeID string) (*models.Like, error) {
	var like models.Like
	if err := ls.db.First(&like, "id = ?", likeID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &like, nil
}

func (ls *likeServiceImpl) DeleteLike(likeID string) error {
	if err := ls.db.Delete(&models.Like{}, "id = ?", likeID).Error; err != nil {
		return err
	}
	return nil
}

func (ls *likeServiceImpl) ListLikes(postID string) ([]*models.Like, error) {
	var likes []*models.Like
	if err := ls.db.Where("post_id = ?", postID).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}

func (ls *likeServiceImpl) DeleteLikeByPostID(postID string) error {
	if err := ls.db.Delete(&models.Like{}, "post_id = ?", postID).Error; err != nil {
		return err
	}
	return nil
}

func (ls *likeServiceImpl) LikePost(postID string, accountID string) error {
	like := &models.Like{
		PostID: postID,
		UserId: accountID,
	}
	if err := ls.db.Create(&like).Error; err != nil {
		return err
	}
	return nil
}

func (ls *likeServiceImpl) UnlikePost(postID string, accountID string) error {
	if err := ls.db.Delete(&models.Like{}, "post_id = ? AND account_id = ?", postID, accountID).Error; err != nil {
		return err
	}
	return nil
}

func (ls *likeServiceImpl) GetLikeCount(postID string) (int64, error) {
	var count int64
	if err := ls.db.Model(&models.Like{}).Where("post_id = ?", postID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
