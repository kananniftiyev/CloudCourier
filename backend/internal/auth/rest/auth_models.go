package rest

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Error struct {
	ErrorN string `json:"error"`
}

type AfterLoginData struct {
}

type AuthorizedUserMessage struct {
	Message        string `json:"message"`
	Token          string `json:"access_token"`
	AfterLoginData `json:"user"`
}
