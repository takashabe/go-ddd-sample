package interfaces

import (
	"os"
	"testing"

	"github.com/takashabe/go-ddd-sample/config"
	fixture "github.com/takashabe/go-fixture"
	_ "github.com/takashabe/go-fixture/mysql"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	db, err := config.NewDBConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fixture, err := fixture.NewFixture(db, "mysql")
	if err != nil {
		panic(err.Error())
	}
	err = fixture.LoadSQL("testdata/schema.sql")
	err = fixture.Load("testdata/users.yml")
	if err != nil {
		panic(err.Error())
	}
}
