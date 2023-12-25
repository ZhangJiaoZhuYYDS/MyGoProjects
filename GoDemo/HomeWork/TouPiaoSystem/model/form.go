package model

type User struct {
	Id uint64
	UserName string   `form:"username"`
	Sex      string   `form:"sex"`
	Hobbies  []string `form:"hobbies"`
	Status int
	VoteFormId uint64
	VoteForm VoteForm



}
type VoteForm struct {
	Title string   `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
	Type string`form:"type" json:"type"`
	Sex string `form:"sex" json:"sex"`
	Hobbies []string `form:"hobbies" json:"hobbies"`
}

type Hobbies struct {
	BaseModel
	UserId uint64
	User User
}