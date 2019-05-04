package controller

type SignUpVM struct {
	Name    string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
