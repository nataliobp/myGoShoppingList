package controller

type UserVM struct {
	Id    int64  `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}
