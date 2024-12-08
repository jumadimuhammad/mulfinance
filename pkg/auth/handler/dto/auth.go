package dto

type Login struct {
	Username string `param:"id" json:"username"`
	Password string `param:"id" json:"password"`
}
