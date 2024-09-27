package repository

import (
	"context"
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
