package repository

import (
	"context"

	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/conversation/entity"
	"gorm.io/gorm"
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

func (r *ConversationRepository) AcceptConversationRequest(ctx context.Context, senderId, recipientId string) error {
	acceptConversationRequest := entity.NewAcceptedConversationRequest()
	result := r.db.Table(acceptConversationRequest.TableName()).Where("sender_id = ? and recipient_id = ?", senderId, recipientId).Updates(acceptConversationRequest)
	if result.Error != nil{
		return result.Error
	}

	if result.RowsAffected == 0{
		return app_error.ErrRecordNotFound
	}

	return nil
	
}
