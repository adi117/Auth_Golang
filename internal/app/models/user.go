package models

import "time"

type User struct {
	Id        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLoginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserReponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
