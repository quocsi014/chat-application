package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/conversation-request/entity"
	conversationEntity "github.com/quocsi014/modules/conversation/entity"
	"gorm.io/gorm"
)

type ConversationRequestRepository struct {
	db *gorm.DB
}

func NewConversationRequestRepository(db *gorm.DB) *ConversationRequestRepository {
	return &ConversationRequestRepository{
		db: db,
	}
}

func (r *ConversationRequestRepository) CreateConversationRequest(ctx context.Context, req *entity.ConversationRequest) error {
	if err := r.db.Table((&entity.ConversationRequest{}).TableName()).WithContext(ctx).Create(req).Error; err != nil {
		return app_error.ErrDatabase(err)
	}
	return nil
}

func DeleteConversationRequest(ctx context.Context, db *gorm.DB, senderId, recipientId string) error {
	conversationRequest := entity.NewConversationRequest(senderId, recipientId)
	result := db.Table(conversationRequest.TableName()).Where("sender_id = ? and recipient_id = ?", senderId, recipientId).Delete(conversationRequest)
	if result.RowsAffected == 0{
		return app_error.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (r *ConversationRequestRepository)DeleteConversationRequest(ctx context.Context, senderId, recipientId string) error{
	return DeleteConversationRequest(ctx, r.db, senderId, recipientId)
}

func (r *ConversationRequestRepository) CreateConversation(ctx context.Context, conversation *conversationEntity.Conversation) error{
	return r.db.Table(conversation.TableName()).Create(conversation).Error
}

func CreateConversationMembership(ctx context.Context, db *gorm.DB, conversationMembership *conversationEntity.ConversationMembership) error{
	return db.Table(conversationMembership.TableName()).Create(conversationMembership).Error
}

func (r *ConversationRequestRepository) CreateConversationMembership(ctx context.Context, conversationMembership *conversationEntity.ConversationMembership) error{
	return r.db.Table(conversationMembership.TableName()).Create(conversationMembership).Error
}

func (r *ConversationRequestRepository)AcceptConversationRequest(ctx context.Context, senderId, recipientId string) (*conversationEntity.Conversation, error){
	tx := r.db.Begin()
	if err := DeleteConversationRequest(ctx, tx, senderId, recipientId); err != nil{
		tx.Rollback()
		return nil,err
	}
	conversationId := uuid.New()
	conversation := conversationEntity.NewConversation(conversationId.String(), false)
	if err := tx.Table(conversation.TableName()).Create(conversation).Error; err != nil{
		tx.Rollback()
		return nil, err
	}

	senderMembership := conversationEntity.NewConversationMembershipMemberRole(conversationId.String(), senderId)
	recipientMembership := conversationEntity.NewConversationMembershipMemberRole(conversationId.String(), recipientId)

	if err := CreateConversationMembership(ctx, tx, senderMembership); err != nil{
		tx.Rollback()
		return nil, err
	}
	if err := CreateConversationMembership(ctx, tx, recipientMembership); err != nil{
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return conversation, nil
}

func (r *ConversationRequestRepository)GetConversationRequestSent(ctx context.Context, senderId string) ([]entity.ConversationRequestDetail, error){
	var conversationReqs []entity.ConversationRequestDetail

	if err := r.db.Table((&entity.ConversationRequestDetail{}).TableName()).Where("sender_id = ?", senderId).Preload("Sender").Preload("Recipient").Find(&conversationReqs).Error; err != nil{
		return nil, err
	}
	return conversationReqs, nil
}

func (r *ConversationRequestRepository)GetConversationRequestReceived(ctx context.Context, recipientId string) ([]entity.ConversationRequestDetail, error){
	var conversationReqs []entity.ConversationRequestDetail

	if err := r.db.Table((&entity.ConversationRequestDetail{}).TableName()).Where("recipient_id = ?", recipientId).Preload("Sender").Preload("Recipient").Find(&conversationReqs).Error; err != nil{
		return nil, err
	}
	return conversationReqs, nil
}

