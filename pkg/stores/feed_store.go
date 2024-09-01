package stores

import (
	"github.com/ochiengotieno304/feedpulse-go/internal/utils"
	"github.com/ochiengotieno304/feedpulse-go/pkg/db"
	"github.com/ochiengotieno304/feedpulse-go/pkg/models"
	"gorm.io/gorm"
)

type FeedStore interface {
	ReadAll(filters map[string]string, page, pageSize int) ([]*models.News, error)
	Read(feedID int) (*models.News, error)
}

type feedStore struct {
	db *gorm.DB
}

func NewFeedStore() FeedStore {
	return &feedStore{
		db: db.DB(),
	}
}

func (s *feedStore) ReadAll(filters map[string]string, page, pageSize int) ([]*models.News, error) {
	var feeds []*models.News
	if err := utils.QueryBuilder(filters, s.db).Scopes(utils.Paginate(page, pageSize)).Find(&feeds).Order("published_date DESC").Error; err != nil {
		return nil, err
	}

	return feeds, nil
}

func (s *feedStore) Read(feedID int) (*models.News, error) {
	var feed models.News
	if err := s.db.First(&feed, feedID).Error; err != nil {
		return nil, err
	}
	return &feed, nil
}
