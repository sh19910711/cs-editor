package main

import (
	_ "github.com/Sirupsen/logrus"
	"github.com/codestand/editor/db"
	"github.com/codestand/editor/dev" // TODO: unload unless development
	"github.com/codestand/editor/project"
	"github.com/codestand/editor/user"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"path/filepath"
)

func createRender() multitemplate.Render {
	r := multitemplate.New()
	includes, err := filepath.Glob("templates/includes/*.html.tmpl")
	if err != nil {
		panic(err)
	}
	for _, include := range includes {
		r.AddFromFiles(filepath.Base(include), "templates/layout.html.tmpl", include)
	}
	return r
}

func main() {
	db.Init()
	defer db.Close()
	db.ORM.AutoMigrate(&user.User{})
	db.ORM.AutoMigrate(&project.Project{})

	store := sessions.NewCookieStore(uuid.NewV4().Bytes()) // TODO: use redis

	r := gin.Default()
	r.HTMLRender = createRender()

	r.Use(sessions.Sessions("_", store))
	r.Static("/assets", "./assets")

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

	// TODO:
	r.Static("/swagger", "./swagger")
	dr := dev.Resource{}
	devpages := r.Group("dev")
	devpages.GET("/doc", dr.Doc)

	r.Run()
}
