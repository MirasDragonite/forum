package structs

type User struct {
	id             int
	username       string
	email          string
	hashedPassword string
}

func (u *User) GetUserID() int {
	return u.id
}

func (u *User) GetUserName() string {
	return u.username
}

func (u *User) GetUserEmail() string {
	return u.email
}

func (u *User) GetUserHashPassword() string {
	return u.hashedPassword
}

func (u *User) ChangeUserName(s string) {
	u.username = s
}

func (u *User) ChangeUserEmail(s string) {
	u.email = s
}

func (u *User) ChangeUserHashPassword(s string) {
	u.hashedPassword = s
}
