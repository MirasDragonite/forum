package structs

type User struct {
	Id             int64
	Username       string
	Email          string
	HashedPassword string
}

type Session struct {
	Id           int64
	UserId       int64
	Token        string
	ExpairedData string
}

func CreateUser(username, email, password string) *User {
	return &User{
		Username:       username,
		Email:          email,
		HashedPassword: password,
	}
}

// func (u *User) GetUserID() int64 {
// 	return u.id
// }

// func (u *User) ChangeUserId(s int64) {
// 	u.id = s
// }

// func (u *User) GetUserName() string {
// 	return u.username
// }

// func (u *User) GetUserEmail() string {
// 	return u.email
// }

// func (u *User) GetUserHashPassword() string {
// 	return u.hashedPassword
// }

// func (u *User) ChangeUserName(s string) {
// 	u.username = s
// }

// func (u *User) ChangeUserEmail(s string) {
// 	u.email = s
// }

// func (u *User) ChangeUserHashPassword(s string) {
// 	u.hashedPassword = s
// }
