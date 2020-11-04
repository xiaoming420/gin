package models

type Users struct {
	ID int `gorm:"primary_key" json:"id"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}

func CheckAuth(phone, password string) int {
	var auth Users
	db.Select("id").Where(Users{Phone : phone, Password : password}).First(&auth)
	if auth.ID > 0 {
		return auth.ID
	}

	return 0
}