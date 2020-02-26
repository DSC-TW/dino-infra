package model

/*
User contains user information while user register the game
*/
type User struct {
	Name           string `json:"name" form:"name"`
	Gender         string `json:"gender" form:"gender"`
	Mail           string `json:"email" form:"mail"`
	School         string `json:"school" form:"school"`
	Department     string `json:"department" form:"department"`
	IsAgreePrivacy bool   `json:"is_agree_privacy" form:"is_agree_privacy"`
}
