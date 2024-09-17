package appTypes

type User struct {
	Id       int    `db:"user_id" json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
