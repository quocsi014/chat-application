package service

import (
	"github.com/quocsi014/modules/conversation/entity"
)

type IConversationService interface {
	GetConversations(userId string) ([]entity.ConversationResponse, error)
}

type conversationService struct {
	repo IConversationRepository
}

func NewConversationService(repo IConversationRepository) IConversationService {
	return &conversationService{repo: repo}
}

func (s *conversationService) GetConversations(userId string) ([]entity.ConversationResponse, error) {
	return s.repo.GetConversations(userId)
}

type IConversationRepository interface {
	GetConversations(userId string) ([]entity.ConversationResponse, error)
}
