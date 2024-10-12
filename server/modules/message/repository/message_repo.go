package repository

import (
	"context"
	"fmt"
	"github.com/quocsi014/common"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/message/entity"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (mr *MessageRepository) InsertMessage(ctx context.Context, message *entity.Message) error {
	return mr.db.Create(message).Error
}

func (mr *MessageRepository) GetMessages(ctx context.Context, paging *common.Paging, conversationId string) ([]entity.Message, error) {
	var totalRows int64
	fmt.Println(conversationId)
	db := mr.db.Table(entity.NewMessage().TableName()).Where("conversation_id = ?", conversationId)
	if err := db.Count(&totalRows).Error; err != nil {
		return nil, app_error.ErrDatabase(err)
	}

	paging.TotalPage = int64(totalRows)/int64(paging.Limit) + 1
	messages := make([]entity.Message, 0)
	if err := db.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Preload("Sender").Find(&messages).Error; err != nil {
		return nil, app_error.ErrDatabase(err)
	}

	return messages, nil
}
