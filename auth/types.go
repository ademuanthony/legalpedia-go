package auth

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	CreatedAt uint64 `json:"created_at"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
