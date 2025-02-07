package model

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
	Age      int    `json:"age"`
}
