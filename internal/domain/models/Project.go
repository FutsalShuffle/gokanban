package models

type Project struct {
	Name     string
	Id       int32
	Slug     string
	Workflow string
}

func (u *Project) ShouldTreatUserAsAdmin(user *User) bool {
	return true
}
