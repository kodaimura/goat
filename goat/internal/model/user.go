package model


type User struct {
	Id int `db:"user_id" json:"user_id"`
	Name string `db:"user_name" json:"user_name"`
	Password string `db:"user_password" json:"user_password"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}