package repository

import (
	"github.com/quocsi014/modules/conversation/entity"
	"gorm.io/gorm"
)

type ConversationRepository interface {
	GetConversations(userId string) ([]entity.ConversationResponse, error)
}

type conversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) *conversationRepository {
	return &conversationRepository{db: db}
}

func (r *conversationRepository) GetConversations(userId string) ([]entity.ConversationResponse, error) {
	var conversations []entity.ConversationResponse
	
	err := r.db.Raw(`
		select c.id, c.is_group, coalesce(gd.name, concat_ws(' ',u.firstname, u.lastname)) as name, 
		coalesce(gd.avatar_url, u.avatar_url) as avatar_url, m.message, 
		concat_ws(' ', u2.firstname, u2.lastname) as user_name_sender 
		from conversation_memberships cm 
		left join conversations c on cm.conversation_id = c.id 
		left join group_details gd on c.id = gd.id and c.is_group = 1 
		left join conversation_memberships cm2 on cm2.conversation_id = c.id and c.is_group = 0 and cm2.user_id != ?
		left join users u on u.id = cm2.user_id 
		left join messages m on c.last_message_id = m.id 
		left join users u2 on u2.id = m.user_id 
		where cm.user_id = ? 
		order by c.last_message_time
	`, userId, userId).Scan(&conversations).Error

	return conversations, err
}
