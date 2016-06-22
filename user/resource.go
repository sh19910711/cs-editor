package user

import (
	"github.com/codestand/editor/db" // TODO: replace with model
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resource struct {
}

func (r *Resource) HomePage(c *gin.Context) {
	s := sessions.Default(c)

	var loginUser User
	loginID := s.Get("user.login_id")
	db.ORM.Where("login_id = ?", loginID).First(&loginUser)

	c.HTML(http.StatusOK, "index.html.tmpl", gin.H{"loginUser": loginUser})
}

func (r *Resource) DashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html.tmpl", gin.H{})
}

func (r *Resource) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html.tmpl", gin.H{})
}

func (r *Resource) Create(c *gin.Context) {
	var newUser User
	c.Bind(&newUser)

	var existUser User
	db.ORM.Where("login_id = ?", newUser.LoginID).First(&existUser)
	if newUser.LoginID == existUser.LoginID {
		c.JSON(http.StatusConflict, gin.H{"msg": "ERROR"})
		return
	}

	newUser.Password = newUser.EncryptedPassword()
	db.ORM.Save(&newUser)
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}

func (r *Resource) Auth(c *gin.Context) {
	// query
	var loginUser User
	c.Bind(&loginUser)
	loginUser.Password = loginUser.EncryptedPassword()

	// target user
	var targetUser User
	db.ORM.Where("login_id = ?", loginUser.LoginID).First(&targetUser)

	if targetUser.Password == loginUser.Password {
		// generate session
		s := sessions.Default(c)
		s.Set("login_id", targetUser.LoginID)
		s.Save()

		c.Redirect(http.StatusFound, "/dashboard")
	} else {
		c.JSON(http.StatusForbidden, gin.H{"msg": "login failed"})
	}
}

func (r *Resource) RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html.tmpl", gin.H{})
}

func (r *Resource) List(c *gin.Context) {
	var users []User
	db.ORM.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}
