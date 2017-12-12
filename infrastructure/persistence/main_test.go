package persistence

import (
	"os"
	"testing"

	"github.com/takashabe/go-ddd-sample/config"
	fixture "github.com/takashabe/go-fixture"
	_ "github.com/takashabe/go-fixture/mysql"
)

func MainTest(m *testing.M) {
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
	err = fixture.Load("testdata/schema.sql")
	if err != nil {
		panic(err.Error())
	}
}
