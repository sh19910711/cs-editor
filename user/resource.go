package user

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resource struct {
}

type UserForm struct {
	LoginId  string `form:"user.login_id" binding:"required,max=32"`
	Password string `form:"user.password" binding:"required,max=128"`
}

func (r *Resource) HomePage(c *gin.Context) {
	s := sessions.Default(c)
	h := gin.H{}
	u, err := Find(s.Get("current_user"))
	if err == nil {
		h["loginUser"] = u
		h["logined"] = true
	}
	c.HTML(http.StatusOK, "index.html.tmpl", h)
}

func (r *Resource) DashboardPage(c *gin.Context) {
	s := sessions.Default(c)
	u, err := Find(s.Get("current_user"))
	log.Infof("user: %s", s.Get("current_user"))
	if err != nil {
		s.Clear()
		c.Redirect(http.StatusFound, "/")
	} else {
		c.HTML(http.StatusOK, "dashboard.html.tmpl", gin.H{"loginUser": u})
	}
}

func (r *Resource) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html.tmpl", gin.H{})
}

func (r *Resource) Create(c *gin.Context) {
	var uf UserForm
	c.Bind(&uf)

	if Exist(uf.LoginId) {
		c.JSON(http.StatusConflict, gin.H{"msg": "ERROR: already exists"})
		return
	}

	u := User{LoginId: uf.LoginId, Password: uf.Password}
	Save(&u)

	c.HTML(http.StatusOK, "ok.html.tmpl", gin.H{"msg": "created"})
}

func (r *Resource) Auth(c *gin.Context) {
	var uf UserForm
	c.Bind(&uf)

	u, err := FindWithPassword(uf.LoginId, uf.Password)
	if err != nil {
		log.Warn(err)
		c.JSON(http.StatusForbidden, gin.H{"msg": "login failed"})
	} else {
		s := sessions.Default(c)
		s.Set("current_user", u.LoginId)
		s.Save()

		c.Redirect(http.StatusFound, "/dashboard")
	}
}

func (r *Resource) RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html.tmpl", gin.H{})
}

func (r *Resource) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": AllUsers()})
}
