package dto


type User struct {
	Id int `json:"user_id"`
	Name string `json:"user_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}