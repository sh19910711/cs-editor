package main

import (
	"github.com/codestand/editor/db"
	"github.com/codestand/editor/project"
	"github.com/codestand/editor/user"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	defer db.Close()
	db.ORM.AutoMigrate(&user.User{})

	store := sessions.NewCookieStore([]byte("TODO: secret?"))

	r := gin.Default()
	r.Use(sessions.Sessions("user", store))
	r.Static("/assets", "./assets")

	r.LoadHTMLGlob("templates/*.tmpl")

	ur := user.Resource{}
	pr := project.Resource{}

	pages := r.Group("/")
	pages.GET("/", ur.HomePage)
	pages.GET("/dashboard", ur.DashboardPage)
	pages.GET("/login", ur.LoginPage)
	pages.POST("/login", ur.Auth)
	pages.GET("/register", ur.RegisterPage)
	pages.POST("/register", ur.Create)

	api := r.Group("/api")
	api.GET("/users", ur.List)
	api.GET("/projects", pr.List)
	api.POST("/projects", pr.Create)
	api.GET("/projects/:id", pr.Show)
	api.GET("/projects/:id/*path", pr.ReadFile)
	api.POST("/projects/:id/*path", pr.CreateFile)
	api.PUT("/projects/:id/*path", pr.UpdateFile)

	r.Run()
}
