package main 

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"` 
	JWT string `json:"jwt"`
	Otp string `json:"otp"`
}

type Post struct {
	author string     // username of the author
	CreatedAt string //TODO: make it a date type compatable with the one in the database
	Content string
	Likes int
	commentsNumber int

}

type Comment struct {
	PostID string
	CommentID string
	content string
	author string
	CreatedAt string
	Likes int
}

// `json:"state"`

type Set map[interface{}]struct{}

func (s Set) Add(value interface{}) {
	s[value] = struct{}{}
}

func (s Set) Remove(value interface{}) {
	delete(s, value)
}

func (s Set) Contains(value interface{}) bool {
	_, ok := s[value]
	return ok
}

type tmpUser struct {
	Email string
	Otp string
	Token string
}