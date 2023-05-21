package service

import (
	"errors"
	"placio-app/models"

	"gorm.io/gorm"
)

type PostService interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetPost(postID string) (*models.Post, error)
	UpdatePost(post *models.Post) (*models.Post, error)
	DeletePost(postID string) error
	ListPosts() ([]*models.Post, error)
}

type PostServiceImpl struct {
	db           *gorm.DB
	store        *models.Post
	// accountStore *models.Account
}

func NewPostService(db *gorm.DB) PostService {
	return &PostServiceImpl{db: db, store: &models.Post{}}
}

func (ps *PostServiceImpl) CreatePost(post *models.Post) (*models.Post, error) {
	if err := ps.db.Create(&post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (ps *PostServiceImpl) GetPost(postID string) (*models.Post, error) {
	var post models.Post
	if err := ps.db.Preload("Medias").Preload("Comments").Preload("Likes").First(&post, "id = ?", postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (ps *PostServiceImpl) UpdatePost(post *models.Post) (*models.Post, error) {
	if err := ps.db.Model(&models.Post{}).Where("id = ?", post.ID).Updates(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (ps *PostServiceImpl) DeletePost(postID string) error {
	if err := ps.db.Delete(&models.Post{}, "id = ?", postID).Error; err != nil {
		return err
	}
	return nil
}

func (ps *PostServiceImpl) ListPosts() ([]*models.Post, error) {
	var posts []*models.Post
	if err := ps.db.Preload("Medias").Preload("Comments").Preload("Likes").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
