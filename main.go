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
	user.AutoMigrate()
	project.AutoMigrate()

	store, err := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("TODO: secret"))
	if err != nil {
		panic("Error: setup session store")
	}

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
	r.Static("/swagger-ui", "./vendor/swagger-ui")
	dr := dev.Resource{}
	devpages := r.Group("dev")
	devpages.GET("/doc", dr.Doc)

	r.Run()
}
