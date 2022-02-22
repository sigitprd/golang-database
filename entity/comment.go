package entity

// struct comment
type Comment struct {
	Id      int32  `json:"id"`
	Email   string `json:"email"`
	Comment string `json:"comment"`
}
