package models

import "time"

// struct para fazer request de aplicações que fazem authenticação com o OAuth do google
type UserRequest struct {
    ID  string `json:"id"`
	Fullname string `json:"fullname"`
	ImageUrl string `json:"imageUrl"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// struct core de user
type User struct {
	ID  string `json:"id"`
	Fullname  string `json:"fullname"`
	ImageUrl  string `json:"imageUrl"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}