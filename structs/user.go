package structs

type User struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
}
