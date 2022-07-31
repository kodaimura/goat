package entity


type User struct {
	UserId int `db:"USER_ID" json:"user_id"`
	UserName string `db:"USER_NAME" json:"user_name"`
	Password string `db:"PASSWORD" json:"password"`
	CreateAt string `db:"CREATE_AT" json:"create_at"`
	UpdateAt string `db:"UPDATE_AT" json:"update_at"`
}