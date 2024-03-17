package api

func newUser(username string, email string) *User {
	return &User{
		Username: username,
		Email: email,
	}
}