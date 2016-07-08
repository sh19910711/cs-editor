package project

import (
	"github.com/codestand/editor/db"
	"github.com/codestand/editor/user"
)

type Project struct {
	ID    int32     `json:"-" db:"id"`
	Name  string    `json:"name"`
	Owner user.User `json:"owner"`
	Files []File    `json:"files"`
}

type File struct {
	ID   int32  `json:"-"`
	Path string `json:path`
}

func AutoMigrate() {
	db.ORM.AutoMigrate(&Project{})
	db.ORM.AutoMigrate(&File{})
}

func AllProjects() (projects []Project) {
	db.ORM.Find(&projects)
	return projects
}

func Save(p *Project) {
	db.ORM.Save(p) // TODO: error handling
}

func Find(u user.User, name string) (p Project, err error) {
	if db.ORM.Where("owner.id = ? and name = ?", u.ID, name).First(&p).RecordNotFound() {
		return p, err
	}
	return p, nil
}

func (p *Project) CreateFile(path string) {
	p.Files = append(p.Files, File{Path: path})
	Save(p)
}

func (p *Project) UpdateFile(path string, content string) {
}
