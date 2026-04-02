package types

import (
	"net/http"
	"time"
)

type User struct {
	Id        int
	FirstName string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Post struct {
	Id        int
	UserId    int
	Content   string
	ImgUrl    string
	CreatedAt time.Time
}
type Follow struct {
	FollowerId  int
	FollowingId int
	CreatedAt   time.Time
}

type PostPayload struct {
	Content string
	ImgUrl  string
}

type RegisterPayload struct {
	FirstName string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type FollowPayload struct {
	FollowerId  int
	FollowingId int `json:"followingid"`
}

type UserStore interface {
	CreateUser(User) error
	GetUserByGmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
}

type PostStore interface {
	CreatePost(http.Request, Post) error
	//get feed from post ids???
	GetPosts(int) ([]Post, error)
}

type FolowStore interface {
	AddFollower(FollowPayload) error
}
