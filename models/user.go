package models

type User struct {
	ID     	  uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Phone 	  string `json:"phone"`
}

type RegisterUser struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Phone 	  string `json:"phone"    binding:"required"`
}

type LoginUser struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
