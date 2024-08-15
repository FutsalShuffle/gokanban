package models

type User struct {
	Name    string
	Id      int32
	IsAdmin bool
}

func (u *User) HasRole(role int32, project *Project) bool {
	return true
}
