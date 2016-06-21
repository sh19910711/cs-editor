package main

import(
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/sessions"
  "github.com/jinzhu/gorm"
  log "github.com/Sirupsen/logrus"
  _ "github.com/mattn/go-sqlite3"
  "encoding/hex"
  "golang.org/x/crypto/scrypt"
)

type UserAccount struct {
  ID int32 `json:"id"`
  LoginID string `json:"login_id" form:"user.login_id" db:"login_id" binding:"required,max=32"`
  Password string `json:"-" form:"user.password" db:"password" binding:"required,max=128"`
}

func (ua *UserAccount) encryptedPassword() string {
  salt := []byte("TODO: change me")
  converted, _ := scrypt.Key([]byte(ua.Password), salt, 16384, 8, 1, 32)
  return hex.EncodeToString(converted[:])
}

type UserResource struct {
  db *gorm.DB
}

func NewUserResource() *UserResource {
  return &UserResource {}
}

func (ur *UserResource) ShowLoginPage(c *gin.Context) {
  c.HTML(http.StatusOK, "login.html.tmpl", gin.H {})
}

func (ur *UserResource) CreateUser(c *gin.Context) {
  var user UserAccount
  c.Bind(&user)

  var existUser UserAccount
  ur.db.Where("login_id = ?", user.LoginID).First(&existUser)
  if user.LoginID == existUser.LoginID {
    c.JSON(http.StatusConflict, gin.H {"msg": "ERROR"})
    return
  }

  user.Password = user.encryptedPassword()
  ur.db.Save(&user)
  c.JSON(http.StatusOK, gin.H {"msg": "OK"})
}

func (ur *UserResource) Auth(c *gin.Context) {
  // query
  var loginUser UserAccount
  c.Bind(&loginUser)
  loginUser.Password = loginUser.encryptedPassword()

  // target user
  var targetUser UserAccount
  ur.db.Where("login_id = ?", loginUser.LoginID).First(&targetUser)

  if targetUser.Password == loginUser.Password {
    // generate session
    s := sessions.Default(c)
    s.Set("user.login_id", targetUser.LoginID)
    s.Save()

    c.Redirect(http.StatusFound, "/dashboard")
  } else {
    c.JSON(http.StatusForbidden, gin.H {"msg": "login failed"})
  }
}

func (ur *UserResource) ShowRegisterPage(c *gin.Context) {
  c.HTML(http.StatusOK, "register.html.tmpl", gin.H {})
}

func (ur *UserResource) ListUsers(c *gin.Context) {
  var users []UserAccount
  ur.db.Find(&users)
  c.JSON(http.StatusOK, gin.H {"users": users})
}

type Service struct {
  engine *gin.Engine
}

func NewService() *Service {
  return &Service {engine: gin.Default()}
}

func showHomePage(c *gin.Context) {
  s := sessions.Default(c)
  var loginUser UserAccount
  loginID := s.Get("user.login_id")
  database.Where("login_id = ?", loginID).First(&loginUser)
  c.HTML(http.StatusOK, "index.html.tmpl", gin.H {"loginUser": loginUser})
}

func showDashboardPage(c *gin.Context) {
  c.HTML(http.StatusOK, "dashboard.html.tmpl", gin.H {})
}

var database *gorm.DB

func (rt *Service) Run() error {
  // setup database
  db, err := gorm.Open("sqlite3", ":memory:") // TODO: change driver
  if err != nil {
    return err
  }
  db.DB()
  db.AutoMigrate(&UserAccount{})
  defer db.Close()
  database = db

  // setup sessions (cookie-based)
  store := sessions.NewCookieStore([]byte("TODO: secret?"))
  rt.engine.Use(sessions.Sessions("user", store))

  // setup routes
  rt.engine.LoadHTMLGlob("templates/*.tmpl")

  rt.engine.GET("/", showHomePage)
  rt.engine.GET("/dashboard", showDashboardPage)

  ur := UserResource { db: db }
  rt.engine.GET("/login", ur.ShowLoginPage)
  rt.engine.POST("/login", ur.Auth)
  rt.engine.GET("/register", ur.ShowRegisterPage)
  rt.engine.POST("/register", ur.CreateUser)
  rt.engine.GET("/users", ur.ListUsers)

  rt.engine.Run()

  return nil
}

func main() {
  es := NewService()
  err := es.Run()
  if err != nil {
    log.Warn(err)
  }
}
