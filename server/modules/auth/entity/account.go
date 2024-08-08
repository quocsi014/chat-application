package entity


type Account struct{
	Id string `json:"id" gorm:"column:id"`
	Email string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func NewAccount(email, password string) *Account{
	return &Account{
		Email: email,
		Password: password,
	}
}

func (a *Account)TableName() string{
	return "accounts"
}

