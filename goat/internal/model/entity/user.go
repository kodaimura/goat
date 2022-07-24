package entity


type User struct {
	UserId int `db:"USER_ID" json:"userid"`
	Username string `db:"USERNAME" json:"username"`
	Password string `db:"PASSWORD" json:"password"`
	CreateAt string `db:"CREATE_AT" json:"createat"`
	UpdateAt string `db:"UPDATE_AT" json:"updateat"`
}