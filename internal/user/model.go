package user

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
	Phone        string `json:"phone" bson:"phone"`
}

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
