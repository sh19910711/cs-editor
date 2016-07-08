package project

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codestand/editor/user"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resource struct {
}

type Query struct {
	Name string `binding:"required"`
}

func (r *Resource) Create(c *gin.Context) {
	var q Query
	err := c.Bind(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "invalid query"})
		return
	}

	s := sessions.Default(c)
	u, err := user.Find(s.Get("current_user"))
	if err == nil {
		p := Project{Name: q.Name, Owner: u}
		Save(&p)
		c.JSON(http.StatusOK, gin.H{"project": p})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Forbidden"})
	}
}

func (r *Resource) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"projects": AllProjects()})
}

func (r *Resource) Show(c *gin.Context) {
	name := c.Param("name")

	s := sessions.Default(c)
	u, err := user.Find(s.Get("current_user"))
	if err == nil {
		p, err := Find(u, name)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"project": p})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "something went wrong"})
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Forbidden"})
	}
}

func (r *Resource) ReadFile(c *gin.Context) {
	// TODO
}

func (r *Resource) CreateFile(c *gin.Context) {
	name := c.Param("name")
	path := c.Param("path")
	log.Info(name, path)

	s := sessions.Default(c)
	u, err := user.Find(s.Get("current_user"))
	if err == nil {
		p, err := Find(u, name)
		if err == nil {
			p.CreateFile(path)
			c.JSON(http.StatusOK, gin.H{"msg": "OK"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "something went wrong"})
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Forbidden"})
	}
}

type FileQuery struct {
	Content string
}

func (r *Resource) UpdateFile(c *gin.Context) {
	name := c.Param("name")
	path := c.Param("path")
	log.Info("name:", name, "path:", path)

	var fq FileQuery
	c.Bind(&fq)

	s := sessions.Default(c)
	u, err := user.Find(s.Get("current_user"))
	if err == nil {
		p, err := Find(u, name)
		if err == nil {
			p.UpdateFile(path, fq.Content)
			c.JSON(http.StatusOK, gin.H{"msg": "OK"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "something is wrong"})
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Forbidden"})
	}
}
