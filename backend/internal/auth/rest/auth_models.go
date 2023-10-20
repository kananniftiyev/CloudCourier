package rest

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Error struct {
	ErrorN string `json:"error"`
}
