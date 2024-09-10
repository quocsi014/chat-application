package service

import (
	"context"
	"errors"

	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/conversation-request/entity"
	userEntity "github.com/quocsi014/modules/user_information/entity"
	conversationEntity "github.com/quocsi014/modules/conversation/entity"
)

type IConversationRequestRepository interface {
	CreateConversationRequest(ctx context.Context, req *entity.ConversationRequest) error
	AcceptConversationRequest(ctx context.Context, senderId, recipientId string) (*conversationEntity.Conversation, error)
	DeleteConversationRequest(ctx context.Context, senderId, recipientId string) error
	GetConversationRequestSent(ctx context.Context, senderId string) ([]entity.ConversationRequestDetail, error)
	GetConversationRequestReceived(ctx context.Context, recipientId string) ([]entity.ConversationRequestDetail, error)
}

type IUserService interface {
	GetUserById(ctx context.Context, userId string) (*userEntity.User, error)
}

type ConversationRequestService struct {
	repo        IConversationRequestRepository
	userService IUserService
}

func NewConversationRequestService(repo IConversationRequestRepository, userService IUserService) *ConversationRequestService {
	return &ConversationRequestService{
		repo:        repo,
		userService: userService,
	}
}

func (s *ConversationRequestService) CreateConversationRequest(ctx context.Context, senderId, recipientId string) error {
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

	req := entity.NewConversationRequest(senderId, recipientId)

	if err := s.repo.CreateConversationRequest(ctx, req); err != nil {
		return app_error.ErrInternal(err)
	}

	return nil
}

func (s *ConversationRequestService) AcceptConversationRequest(ctx context.Context, senderId, recipientId string) (*conversationEntity.Conversation, error){
	conversation, err := s.repo.AcceptConversationRequest(ctx, senderId, recipientId)
	if err != nil{
		if errors.Is(err, app_error.ErrRecordNotFound){
			return nil, app_error.ErrNotFound(err, "CONV_REQ_NOT_EXIST", "no conversation requests found")
		}
		return nil, app_error.ErrDatabase(err)
	}	
	return conversation, nil
}

func (s *ConversationRequestService) DeleteConversationRequest(ctx context.Context, senderId, recipientId string) error{
	err := s.repo.DeleteConversationRequest(ctx, senderId, recipientId)
	if err != nil{
		if errors.Is(err, app_error.ErrRecordNotFound){
			return app_error.ErrNotFound(err, "CONV_REQ_NOT_EXIST", "no conversation requests found")
		}
		return app_error.ErrDatabase(err)
	}
	return nil
}

func (s *ConversationRequestService) GetConversationRequestSent(ctx context.Context, senderId string) ([]entity.ConversationRequestDetail, error){
	conversationReqs, err := s.repo.GetConversationRequestSent(ctx, senderId)
	if err != nil {
		return nil, app_error.ErrDatabase(err)
	}
	return conversationReqs, nil
}

func (s *ConversationRequestService) GetConversationRequestReceived(ctx context.Context, recipientId string) ([]entity.ConversationRequestDetail, error){
	conversationReqs, err := s.repo.GetConversationRequestReceived(ctx, recipientId)
	if err != nil {
		return nil, app_error.ErrDatabase(err)
	}
	return conversationReqs, nil
}
