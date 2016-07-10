package project

import (
	"github.com/codestand/editor/db"
	"github.com/codestand/editor/user"
)

type Project struct {
	Id      int32  `json:"-" db:"id"`
	Name    string `json:"name" db:"name" sql:"unique"`
	OwnerId int32  `json:"owner_id" db:"owner_id" sql:"unique"`
	Files   []File `json:"files" db:"files"`
}

type File struct {
	Id   int32  `json:"-"`
	Path string `json:path`
}

func AutoMigrate() {
	db.ORM.AutoMigrate(&Project{})
	db.ORM.AutoMigrate(&File{})
}

func AllProjects(u user.User) (projects []Project) {
	db.ORM.Where("owner_id = ?", u.Id).Find(&projects)
	return projects
}

func Save(p *Project) {
	db.ORM.Save(p) // TODO: error handling
}

func Find(u user.User, name string) (p Project, err error) {
	if db.ORM.Where("owner_id = ? and name = ?", u.Id, name).First(&p).RecordNotFound() {
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
