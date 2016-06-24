package project

import (
	"github.com/codestand/editor/db"
	"github.com/codestand/editor/user"
)

type Project struct {
	ID    int32     `json:"-" db:"id"`
	Owner user.User `json:"owner"`
}

func AllProjects() (projects []Project) {
	db.ORM.Find(&projects)
	return projects
}

func Save(p Project) {
	db.ORM.Save(&p) // TODO: error handling
}
