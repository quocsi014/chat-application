package repository

import (
	"context"

	"gorm.io/gorm"
	"github.com/quocsi014/modules/conversation/entity"
	"github.com/quocsi014/common/app_error"
)

type ConversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) *ConversationRepository {
	return &ConversationRepository{
		db: db,
	}
}

func (r *ConversationRepository) CreateConversationRequest(ctx context.Context, req *entity.ConversationRequest) error {
	if err := r.db.Table((&entity.ConversationRequest{}).TableName()).WithContext(ctx).Create(req).Error; err != nil {
		return app_error.ErrDatabase(err)
	}
	return nil
}