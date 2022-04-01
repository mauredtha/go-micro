package tests

import (
	"database/sql"
	"microservices/libraries/api"
	"microservices/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
)

func TestUser(t *testing.T) {
	db, teardown := NewUnit(t)
	defer teardown()

	u := User{Db: db}
	t.Run("CRUD", u.Crud)
	t.Run("List", u.List)
}

// User struct for test users
type User struct {
	Db *sql.DB
}

//Crud : unit test  for create get and delete user function
func (u *User) Crud(t *testing.T) {
	u0 := models.User{
		Username: "Aladin",
		Email:    "aladin@gmail.com",
		Password: "1234",
		IsActive: false,
	}

	err := u0.Create(u.Db)
	if err != nil {
		t.Fatalf("creating user u0: %s", err)
	}

	u1 := models.User{
		ID: u0.ID,
	}

	err = u1.Get(u.Db)
	if err != nil {
		t.Fatalf("getting user u1: %s", err)
	}

	if diff := cmp.Diff(u1, u0); diff != "" {
		t.Fatalf("fetched != created:\n%s", diff)
	}

	u1.IsActive = false
	err = u1.Update(u.Db)
	if err != nil {
		t.Fatalf("update user u1: %s", err)
	}

	u2 := models.User{
		ID: u1.ID,
	}

	err = u2.Get(u.Db)
	if err != nil {
		t.Fatalf("getting user u2: %s", err)
	}

	if diff := cmp.Diff(u1, u2); diff != "" {
		t.Fatalf("fetched != updated:\n%s", diff)
	}

	err = u2.Delete(u.Db)
	if err != nil {
		t.Fatalf("delete user u2: %s", err)
	}

	u3 := models.User{
		ID: u2.ID,
	}

	err = u3.Get(u.Db)

	apiErr, ok := err.(*api.Error)
	if !ok || apiErr.Err != sql.ErrNoRows {
		t.Fatalf("getting user u3: %s", err)
	}
}

//List : unit test for user list function
func (u *User) List(t *testing.T) {
	u0 := models.User{
		Username: "Aladin",
		Email:    "aladin@gmail.com",
		Password: "1234",
		IsActive: false,
	}

	err := u0.Create(u.Db)
	if err != nil {
		t.Fatalf("creating user u0: %s", err)
	}

	var user models.User
	users, err := user.List(u.Db)
	if err != nil {
		t.Fatalf("listing users: %s", err)
	}
	if exp, got := 1, len(users); exp != got {
		t.Fatalf("expected users list size %v, got %v", exp, got)
	}
}

//go test -v essentials/tests -run TestUser
