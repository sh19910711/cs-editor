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

func (r *Resource) Create(c *gin.Context) {
	s := sessions.Default(c)
	log.Infof("user: %s", s.Get("current_user"))
	u, err := user.Find(s.Get("current_user"))
	if err == nil {
		Save(Project{Owner: u})
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Forbidden"})
	}
}

func (r *Resource) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"projects": AllProjects()})
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
