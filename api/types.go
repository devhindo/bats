package api

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"` 
}

type Post struct {
	Author string   // username of the author
	Content string
	CreatedAt string //TODO: make it a date type compatable with the one in the database
}



// `json:"state"`