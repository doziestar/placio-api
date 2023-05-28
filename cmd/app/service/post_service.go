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
	ListPosts(page, pageSize int, sortBy string, filters map[string]interface{}) ([]*models.Post, error)

	CreateComment(postID string, comment *models.Comment) (*models.Comment, error)
	UpdateComment(comment *models.Comment) (*models.Comment, error)
	DeleteComment(commentID string) error
	GetComment(commentID string) (*models.Comment, error)
	GetComments(postID string, page, pageSize int, sortBy string, filters map[string]interface{}) ([]*models.Comment, error)

	LikePost(postID string, userID string) error
	UnlikePost(postID string, userID string) error
	GetLikes(postID string, page, pageSize int) ([]*models.Like, error)
}

type PostServiceImpl struct {
	db    *gorm.DB
	store *models.Post
	// accountStore *models.Account
}

func NewPostService(db *gorm.DB) PostService {
	return &PostServiceImpl{db: db, store: &models.Post{}}
}

func (ps *PostServiceImpl) CreatePost(post *models.Post) (*models.Post, error) {
	if post == nil {
		return nil, errors.New("post cannot be nil")
	}
	if err := ps.db.Create(&post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (ps *PostServiceImpl) GetPost(postID string) (*models.Post, error) {
	if postID == "" {
		return nil, errors.New("postID cannot be empty")
	}

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
	if post == nil {
		return nil, errors.New("post cannot be nil")
	}
	if err := ps.db.Model(&models.Post{}).Where("id = ?", post.ID).Updates(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (ps *PostServiceImpl) DeletePost(postID string) error {
	if postID == "" {
		return errors.New("postID cannot be empty")
	}
	if err := ps.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Post{}, "id = ?", postID).Error; err != nil {
			return err
		}
		// Delete associated comments, likes, or perform any other necessary cleanup
		// ...
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (ps *PostServiceImpl) CreateComment(postID string, comment *models.Comment) (*models.Comment, error) {
	if postID == "" {
		return nil, errors.New("postID cannot be empty")
	}
	if comment == nil {
		return nil, errors.New("comment cannot be nil")
	}

	comment.PostID = postID
	if err := ps.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (ps *PostServiceImpl) UpdateComment(comment *models.Comment) (*models.Comment, error) {
	if comment == nil {
		return nil, errors.New("comment cannot be nil")
	}
	if err := ps.db.Model(&models.Comment{}).Where("id = ?", comment.ID).Updates(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (ps *PostServiceImpl) DeleteComment(commentID string) error {
	if commentID == "" {
		return errors.New("commentID cannot be empty")
	}
	if err := ps.db.Delete(&models.Comment{}, "id = ?", commentID).Error; err != nil {
		return err
	}
	return nil
}

func (ps *PostServiceImpl) LikePost(postID string, userID string) error {
	if postID == "" {
		return errors.New("postID cannot be empty")
	}
	if userID == "" {
		return errors.New("userID cannot be empty")
	}

	like := &models.Like{
		PostID: postID,
		UserId: userID,
	}
	if err := ps.db.Create(&like).Error; err != nil {
		return err
	}
	return nil
}

func (ps *PostServiceImpl) ListPosts(page, pageSize int, sortBy string, filters map[string]interface{}) ([]*models.Post, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, errors.New("page and pageSize must be positive integers")
	}

	offset := (page - 1) * pageSize

	query := ps.db.Model(&models.Post{})

	// Apply filters
	for key, value := range filters {
		query = query.Where(key, value)
	}

	// Apply sorting
	if sortBy != "" {
		query = query.Order(sortBy)
	}

	var posts []*models.Post
	if err := query.Offset(offset).Limit(pageSize).Preload("Medias").Preload("Comments").Preload("Likes").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (ps *PostServiceImpl) GetComments(postID string, page, pageSize int, sortBy string, filters map[string]interface{}) ([]*models.Comment, error) {
	if postID == "" {
		return nil, errors.New("postID cannot be empty")
	}
	if page <= 0 || pageSize <= 0 {
		return nil, errors.New("page and pageSize must be positive integers")
	}

	offset := (page - 1) * pageSize

	query := ps.db.Model(&models.Comment{}).Where("post_id = ?", postID)

	// Apply filters
	for key, value := range filters {
		query = query.Where(key, value)
	}

	// Apply sorting
	if sortBy != "" {
		query = query.Order(sortBy)
	}

	var comments []*models.Comment
	if err := query.Offset(offset).Limit(pageSize).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (ps *PostServiceImpl) GetComment(commentID string) (*models.Comment, error) {
	if commentID == "" {
		return nil, errors.New("commentID cannot be empty")
	}

	var comment models.Comment
	if err := ps.db.First(&comment, "id = ?", commentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

func (ps *PostServiceImpl) UnlikePost(postID string, userID string) error {
	if postID == "" {
		return errors.New("postID cannot be empty")
	}
	if userID == "" {
		return errors.New("userID cannot be empty")
	}

	if err := ps.db.Delete(&models.Like{}, "post_id = ? AND user_id = ?", postID, userID).Error; err != nil {
		return err
	}
	return nil
}

func (ps *PostServiceImpl) GetLikes(postID string, page, pageSize int) ([]*models.Like, error) {
	if postID == "" {
		return nil, errors.New("postID cannot be empty")
	}
	if page <= 0 || pageSize <= 0 {
		return nil, errors.New("page and pageSize must be positive integers")
	}

	offset := (page - 1) * pageSize
	var likes []*models.Like
	if err := ps.db.Where("post_id = ?", postID).Offset(offset).Limit(pageSize).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}
