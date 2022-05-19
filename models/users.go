package models

type User struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Address   string `json:"address"`
}
