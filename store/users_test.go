package store

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var testUsers []User = []User{
	{
		Email:      "a@b.com",
		Password:   "1234",
		FirstName:  "Abraham",
		LastName:   "Incoherent",
		SessionKey: "abcd",
	},
	{
		Email:      "d@x.com",
		Password:   "5678",
		FirstName:  "Jonathan",
		LastName:   "Davis",
		SessionKey: "efgh",
	},
}

func setupUsers(t *testing.T) (UserStore, string) {
	tmpDir, err := ioutil.TempDir("", "contabi-tests-")
	if err != nil {
		t.Fatal(err)
	}
	for _, u := range testUsers {
		b, err := json.Marshal(u)
		if err != nil {
			t.Fatal(err)
		}
		fname := filepath.Join(tmpDir, u.Email)
		err = ioutil.WriteFile(fname, b, 0666)
		if err != nil {
			t.Fatal(err)
		}
	}
	return &userStore{location: tmpDir}, tmpDir
}

func TestUsersGet(t *testing.T) {
	tst, tstDir := setupUsers(t)
	for _, tu := range testUsers {
		u, err := tst.Get(tu.Email)
		if err != nil {
			t.Fatal(err)
		}
		if tu.Email != u.Email || tu.Password != u.Password ||
			tu.FirstName != u.FirstName || tu.LastName != u.LastName ||
			tu.SessionKey != u.SessionKey {
			log.Fatalf("did not get correct user, got:\n%v vs.\n%v\n", *u, tu)
		}
	}
	if err := os.RemoveAll(tstDir); err != nil {
		log.Printf("cannot clean up: %v\n", err)
	}
}

func TestUsersUpdate(t *testing.T) {
	u1 := testUsers[0]
	start := time.Now()
	tst, tstDir := setupUsers(t)
	if _, err := tst.Update(&User{
		Email:      u1.Email,
		Password:   "password",
		FirstName:  "firstname",
		LastName:   "lastname",
		SessionKey: "sessionkey",
	}); err != nil {
		t.Fatal(err)
	}
	u, err := tst.Get(u1.Email)
	if err != nil {
		t.Fatal(err)
	}
	if u.Password != "password" || u.FirstName != "firstname" ||
		u.LastName != "lastname" || u.SessionKey != "sessionkey" ||
		!u.SessionKeyTime.After(start) || !u.SessionKeyTime.Before(time.Now()) {
		t.Fatal("did not update correctly")
	}
	if err := os.RemoveAll(tstDir); err != nil {
		log.Printf("cannot clean up: %v\n", err)
	}
}
