package repository

import (
	"github.com/quocsi014/common"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/conversation/entity"
	"gorm.io/gorm"
)

type ConversationRepository interface {
	GetConversations(userId string, paging *common.Paging) ([]entity.ConversationResponse, error)
}

type conversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) *conversationRepository {
	return &conversationRepository{db: db}
}

func (r *conversationRepository) GetConversations(userId string, paging *common.Paging) ([]entity.ConversationResponse, error) {
	conversations := make([]entity.ConversationResponse, 0)
	var totalRows int64

	if err := r.db.Table((&entity.ConversationMembership{}).TableName()).Where("user_id = ?", userId).Count(&totalRows).Error; err != nil {
		return conversations, app_error.ErrDatabase(err)
	}

	paging.TotalPage = int64(totalRows) / int64(paging.Limit)
	if int64(totalRows)%int64(paging.Limit) != 0 {
		paging.TotalPage++
	}

	db := r.db.Raw(`
	select c.id, c.is_group, coalesce(gd.name, concat_ws(' ',u.firstname, u.lastname)) as name,
	c.last_message_id, c.last_message_time,
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
	limit ? offset ?
	`, userId, userId, paging.Limit, (paging.Page-1)*paging.Limit)

	if err := db.Scan(&conversations).Error; err != nil {
		return conversations, app_error.ErrDatabase(err)
	}

	return conversations, nil
}
