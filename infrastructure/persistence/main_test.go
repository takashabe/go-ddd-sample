package persistence

import (
	"log"
	"testing"

	"github.com/takashabe/go-ddd-sample/config"
	fixture "github.com/takashabe/go-fixture"
	_ "github.com/takashabe/go-fixture/mysql"
)

func MainTest(m *testing.M) {
	setup()
	log.Fatal(m.Run())
}

func setup() {
	db, err := config.NewDBConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fixture := fixture.NewFixture(db, "mysql")
	err = fixture.Load("testdata/schema.sql")
	if err != nil {
		panic(err.Error())
	}
}
