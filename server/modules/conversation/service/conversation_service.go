package service

import (
	"context"
	"errors"

	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/conversation/entity"
	conversationEntity "github.com/quocsi014/modules/conversation/entity"
	userEntity "github.com/quocsi014/modules/user_information/entity"
)

type IConversationRepository interface {
	CreateConversationRequest(ctx context.Context, req *conversationEntity.ConversationRequest) error
	AcceptConversationRequest(ctx context.Context, senderId, recipientId string) (*entity.Conversation, error)
	DeleteConversationRequest(ctx context.Context, senderId, recipientId string) error
	GetConversationRequestSent(ctx context.Context, senderId string) ([]entity.ConversationRequestDetail, error)
}

type IUserService interface {
	GetUserById(ctx context.Context, userId string) (*userEntity.User, error)
}

type ConversationService struct {
	repo        IConversationRepository
	userService IUserService
}

func NewConversationService(repo IConversationRepository, userService IUserService) *ConversationService {
	return &ConversationService{
		repo:        repo,
		userService: userService,
	}
}

func (s *ConversationService) CreateConversationRequest(ctx context.Context, senderId, recipientId string) error {
	if senderId == recipientId {
		return app_error.ErrInvalidRequest(errors.New("sender and recipient cannot be the same"))
	}

	// Kiểm tra sự tồn tại của người nhận
	_, err := s.userService.GetUserById(ctx, recipientId)
	if err != nil {
		if errors.Is(err, app_error.ErrRecordNotFound) {
			return app_error.ErrInvalidRequest(errors.New("recipient does not exist"))
		}
		return app_error.ErrInternal(err)
	}

	req := conversationEntity.NewConversationRequest(senderId, recipientId)

	if err := s.repo.CreateConversationRequest(ctx, req); err != nil {
		return app_error.ErrInternal(err)
	}

	return nil
}

func (s *ConversationService) AcceptConversationRequest(ctx context.Context, senderId, recipientId string) (*entity.Conversation, error){
	conversation, err := s.repo.AcceptConversationRequest(ctx, senderId, recipientId)
	if err != nil{
		if errors.Is(err, app_error.ErrRecordNotFound){
			return nil, app_error.ErrNotFound(err, "CONV_REQ_NOT_EXIST", "no conversation requests found")
		}
		return nil, app_error.ErrDatabase(err)
	}	
	return conversation, nil
}

func (s *ConversationService) DeleteConversationRequest(ctx context.Context, senderId, recipientId string) error{
	err := s.repo.DeleteConversationRequest(ctx, senderId, recipientId)
	if err != nil{
		if errors.Is(err, app_error.ErrRecordNotFound){
			return app_error.ErrNotFound(err, "CONV_REQ_NOT_EXIST", "no conversation requests found")
		}
		return app_error.ErrDatabase(err)
	}
	return nil
}

func (s *ConversationService) GetConversationRequestSent(ctx context.Context, senderId string) ([]entity.ConversationRequestDetail, error){
	conversationReqs, err := s.repo.GetConversationRequestSent(ctx, senderId)
	if err != nil {
		return nil, app_error.ErrDatabase(err)
	}
	return conversationReqs, nil
}
