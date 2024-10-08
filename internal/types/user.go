package types

type User struct {
	Id string `json:"id"`

	Name  string `json:"name"`
	Email string `json:"email"`

	RefreshToken []byte `json:"-"`
}
