package service

import (
	"context"
	"errors"
	"github.com/quocsi014/common"
	"github.com/quocsi014/modules/user_information/service"

	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/conversation-request/entity"
	conversationEntity "github.com/quocsi014/modules/conversation/entity"
)

type ConversationRequestService struct {
	repo        IConversationRequestRepository
	userService service.IUserService
}

func NewConversationRequestService(repo IConversationRequestRepository, userService service.IUserService) *ConversationRequestService {
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

func (s *ConversationRequestService) AcceptConversationRequest(ctx context.Context, senderId, recipientId string) (*conversationEntity.Conversation, error) {
	conversation, err := s.repo.AcceptConversationRequest(ctx, senderId, recipientId)
	if err != nil {
		if errors.Is(err, app_error.ErrRecordNotFound) {
			return nil, app_error.ErrNotFound(err, "CONV_REQ_NOT_EXIST", "no conversation requests found")
		}
		return nil, app_error.ErrDatabase(err)
	}
	return conversation, nil
}

func (s *ConversationRequestService) DeleteConversationRequest(ctx context.Context, senderId, recipientId string) error {
	err := s.repo.DeleteConversationRequest(ctx, senderId, recipientId)
	if err != nil {
		if errors.Is(err, app_error.ErrRecordNotFound) {
			return app_error.ErrNotFound(err, "CONV_REQ_NOT_EXIST", "no conversation requests found")
		}
		return app_error.ErrDatabase(err)
	}
	return nil
}

func (s *ConversationRequestService) GetConversationRequestSent(ctx context.Context, senderId string, paging *common.Paging) ([]entity.ConversationRequestDetail, error) {
	conversationReqs, err := s.repo.GetConversationRequestSent(ctx, senderId, paging)
	if err != nil {
		return nil, app_error.ErrDatabase(err)
	}
	return conversationReqs, nil
}

func (s *ConversationRequestService) GetConversationRequestReceived(ctx context.Context, recipientId string, paging *common.Paging) ([]entity.ConversationRequestDetail, error) {
	conversationReqs, err := s.repo.GetConversationRequestReceived(ctx, recipientId, paging)
	if err != nil {
		return nil, app_error.ErrDatabase(err)
	}
	return conversationReqs, nil
}
