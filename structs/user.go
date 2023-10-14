package structs

type User struct {
	Id             int64  `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
	CreatedDate    string `json:"createdDate"`
}

type Session struct {
	Id           int64
	UserId       int64
	Token        string
	ExpairedData string
}

type Data struct {
	Status int `json:"status"`
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
