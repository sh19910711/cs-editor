package user

import "testing"
import "github.com/codestand/editor/db"
import "flag"
import "os"

func TestFindWithPasswordCorrect(t *testing.T) {
	_, err := FindWithPassword("foo", "bar")
	if err != nil {
		t.Errorf("can't login")
	}
}

func TestFindWithPasswordWrong(t *testing.T) {
	_, err := FindWithPassword("foo", "baZ")
	if err == nil {
		t.Errorf("can login?")
	}
}

func TestMain(m *testing.M) {
	// init database
	db.Init()
	defer db.Close()
	db.ORM.AutoMigrate(User{})

	// create a fake user
	Save(&User{LoginID: "foo", Password: "bar"})

	flag.Parse()
	os.Exit(m.Run())
}
