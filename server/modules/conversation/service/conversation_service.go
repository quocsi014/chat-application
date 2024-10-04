package service

import (
	"github.com/quocsi014/common"
	"github.com/quocsi014/modules/conversation/entity"
)

type conversationService struct {
	repo IConversationRepository
}

func NewConversationService(repo IConversationRepository) IConversationService {
	return &conversationService{repo: repo}
}

func (s *conversationService) GetConversations(userId string, paging *common.Paging) ([]entity.ConversationResponse, error) {
	return s.repo.GetConversations(userId, paging)
}
