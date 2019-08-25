package persistence

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/takashabe/go-ddd-sample/config"
	"github.com/takashabe/go-ddd-sample/domain"
	fixture "github.com/takashabe/go-fixture"
	_ "github.com/takashabe/go-fixture/mysql"
)

func loadFixture(t *testing.T, conn *sql.DB, file string) {
	fixture, err := fixture.NewFixture(conn, "mysql")
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	err = fixture.Load(file)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
}

func TestGetUser(t *testing.T) {
	conn, err := config.NewDBConnection()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	loadFixture(t, conn, "testdata/users.yml")

	cases := []struct {
		input      int
		expectUser *domain.User
		expectErr  error
	}{
		{
			1,
			&domain.User{
				ID:   1,
				Name: "satoshi",
			},
			nil,
		},
		{
			0,
			nil,
			sql.ErrNoRows,
		},
	}
	for i, c := range cases {
		repo := NewUserRepository(conn)
		user, err := repo.Get(context.Background(), c.input)
		if err != c.expectErr {
			t.Fatalf("#%d: want error %#v, got %#v", i, c.expectErr, err)
		}
		if err != nil {
			continue
		}
		if !reflect.DeepEqual(user, c.expectUser) {
			t.Errorf("#%d: want %#v, got %#v", i, c.expectUser, user)
		}
	}
}

func TestGetUsers(t *testing.T) {
	conn, err := config.NewDBConnection()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	loadFixture(t, conn, "testdata/users.yml")

	cases := []struct {
		expectIDs []int
	}{
		{[]int{1, 2, 3, 4, 5}},
	}
	for i, c := range cases {
		repo := NewUserRepository(conn)
		users, err := repo.GetAll(context.Background())
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}
		ids := make([]int, 0)
		for _, u := range users {
			ids = append(ids, u.ID)
		}
		if !reflect.DeepEqual(ids, c.expectIDs) {
			t.Errorf("#%d: want %#v, got %#v", i, c.expectIDs, ids)
		}
	}
}

func TestSaveUser(t *testing.T) {
	conn, err := config.NewDBConnection()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	loadFixture(t, conn, "testdata/users.yml")

	cases := []struct {
		input string
	}{
		{"foo"},
		{"foo"}, // duplicate
	}
	for i, c := range cases {
		ctx := context.Background()
		repo := NewUserRepository(conn)
		err := repo.Save(ctx, &domain.User{
			Name: c.input,
		})
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}

		users, err := repo.GetAll(ctx)
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}
		find := false
		for _, u := range users {
			if u.Name == c.input {
				find = true
				break
			}
		}
		if !find {
			t.Errorf("#%d: want contain name %s, but not found it", i, c.input)
		}
	}
}
