package service

import (
	"context"
	"github.com/quocsi014/common"
	"github.com/quocsi014/modules/message/entity"
)

type IMessageRepo interface {
	InsertMessage(ctx context.Context, message *entity.Message) error
	GetMessages(ctx context.Context, paging *common.Paging, conversationId string) ([]entity.Message, error)
}

type IMessageService interface {
	SendMessage(ctx context.Context, message *entity.Message) error
	GetMessages(ctx context.Context, paging *common.Paging, conversationId string) ([]entity.Message, error)
}
