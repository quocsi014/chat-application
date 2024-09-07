package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/conversation/entity"
	"gorm.io/gorm"
)

type ConversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) *ConversationRepository {
	return &ConversationRepository{
		db: db,
	}
}

func (r *ConversationRepository) CreateConversationRequest(ctx context.Context, req *entity.ConversationRequest) error {
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
func (r *ConversationRepository)DeleteConversationRequest(ctx context.Context, senderId, recipientId string) error{
	return DeleteConversationRequest(ctx, r.db, senderId, recipientId)
}

func (r *ConversationRepository) CreateConversation(ctx context.Context, conversation *entity.Conversation) error{
	return r.db.Table(conversation.TableName()).Create(conversation).Error
}

func CreateConversationMembership(ctx context.Context, db *gorm.DB, conversationMembership *entity.ConversationMembership) error{
	return db.Table(conversationMembership.TableName()).Create(conversationMembership).Error
}

func (r *ConversationRepository) CreateConversationMembership(ctx context.Context, conversationMembership *entity.ConversationMembership) error{
	return r.db.Table(conversationMembership.TableName()).Create(conversationMembership).Error
}

func (r *ConversationRepository)AcceptConversationRequest(ctx context.Context, senderId, recipientId string) (*entity.Conversation, error){
	tx := r.db.Begin()
	if err := DeleteConversationRequest(ctx, tx, senderId, recipientId); err != nil{
		tx.Rollback()
		return nil,err
	}
	conversationId := uuid.New()
	conversation := entity.NewConversation(conversationId.String(), false)
	if err := tx.Table(conversation.TableName()).Create(conversation).Error; err != nil{
		tx.Rollback()
		return nil, err
	}

	senderMembership := entity.NewConversationMembershipMemberRole(conversationId.String(), senderId)
	recipientMembership := entity.NewConversationMembershipMemberRole(conversationId.String(), recipientId)

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

func (r *ConversationRepository)GetConversationRequestSent(ctx context.Context, senderId string) ([]entity.ConversationRequestDetail, error){
	var conversationReqs []entity.ConversationRequestDetail

	if err := r.db.Table((&entity.ConversationRequestDetail{}).TableName()).Where("sender_id = ?", senderId).Preload("Sender").Preload("Recipient").Find(&conversationReqs).Error; err != nil{
		return nil, err
	}
	return conversationReqs, nil
}

func (r *ConversationRepository)GetConversationRequestReceived(ctx context.Context, recipientId string) ([]entity.ConversationRequestDetail, error){
	var conversationReqs []entity.ConversationRequestDetail

	if err := r.db.Table((&entity.ConversationRequestDetail{}).TableName()).Where("recipient_id = ?", recipientId).Preload("Sender").Preload("Recipient").Find(&conversationReqs).Error; err != nil{
		return nil, err
	}
	return conversationReqs, nil
}

