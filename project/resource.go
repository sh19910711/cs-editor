package project

import (
	"github.com/codestand/editor/db" // TODO: replace with model
	_ "github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resource struct {
}

func (r *Resource) Create(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func (r *Resource) List(c *gin.Context) {
	var projects []Project
	db.ORM.Find(&projects)
	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

func (r *Resource) Show(c *gin.Context) {
	// TODO
}

func (r *Resource) ReadFile(c *gin.Context) {
	// TODO
}

func (r *Resource) CreateFile(c *gin.Context) {
	// TODO
}

func (r *Resource) UpdateFile(c *gin.Context) {
	// TODO
}
