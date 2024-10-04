package service

import (
	"context"
	"github.com/quocsi014/common"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/message/entity"
)

type MessageService struct {
	repo IMessageRepo
}

func NewMessageService(repo IMessageRepo) *MessageService {
	return &MessageService{
		repo: repo,
	}
}

func (ms *MessageService) SendMessage(ctx context.Context, message *entity.Message) error {
	if err := ms.repo.InsertMessage(ctx, message); err != nil {
		return app_error.ErrDatabase(err)
	}
	return nil
}

func (ms *MessageService) GetMessages(ctx context.Context, paging *common.Paging, conversationId string) ([]entity.Message, error) {
	return ms.repo.GetMessages(ctx, paging, conversationId)
}
