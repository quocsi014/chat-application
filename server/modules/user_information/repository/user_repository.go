package repository

import (
	"context"
	"errors"

	"github.com/quocsi014/common"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/user_information/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) InsertUser(ctx context.Context, user *entity.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepository) FindUserById(ctx context.Context, id string) (*entity.User, error) {
	user := entity.User{}
	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app_error.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return repo.db.Model(&entity.User{}).Where("id = ?", user.Id).Updates(user).Error
}

func (repo *UserRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := entity.User{}
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app_error.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetUsersByUsername(ctx context.Context, username string, paging *common.Paging) ([]entity.User, error) {
	users := make([]entity.User, 0)
	var totalRows int64

	db := repo.db.Table((&entity.User{}).TableName()).Where("username LIKE ?", "%"+username+"%")
	if err := db.Count(&totalRows).Error; err != nil {
		return nil, err
	}

	offset := (paging.Page - 1) * paging.Limit
	db = repo.db.Raw("select * from users u left join user_relationships ur on u.id = ur.user_id or u.id = ur.friend_id where u.username like ? limit ? offset ?", "%"+username+"%", paging.Limit, offset)
	if err := db.Scan(&users).Error; err != nil {
		return nil, err
	}
	paging.TotalPage = int64(totalRows) / int64(paging.Limit)
	if int64(totalRows)%int64(paging.Limit) != 0 {
		paging.TotalPage++
	}

	return users, nil
}
