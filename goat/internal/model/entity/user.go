package entity


type User struct {
	UserId int `db:"USER_ID" json:"user_id"`
	Username string `db:"USERNAME" json:"username"`
	Password string `db:"PASSWORD" json:"password"`
	CreateAt string `db:"CREATE_AT" json:"create_at"`
	UpdateAt string `db:"UPDATE_AT" json:"update_at"`
}