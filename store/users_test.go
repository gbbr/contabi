package store

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var testUser1 User = User{
	Email:      "a@b.com",
	Password:   "1234",
	FirstName:  "Abraham",
	LastName:   "Incoherent",
	SessionKey: "abcd",
}

func setupUsers(t *testing.T) (UserStore, string) {
	tmpDir, err := ioutil.TempDir("", "contabi-tests-")
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(testUser1)
	if err != nil {
		t.Fatal(err)
	}
	fname := filepath.Join(tmpDir, "a@b.com")
	err = ioutil.WriteFile(fname, b, 0666)
	if err != nil {
		t.Fatal(err)
	}
	return &userStore{location: tmpDir}, tmpDir
}

func TestUsersGet(t *testing.T) {
	tst, tstDir := setupUsers(t)
	u, err := tst.Get("a@b.com")
	if err != nil {
		t.Fatal(err)
	}
	if testUser1.Email != u.Email || testUser1.Password != u.Password ||
		testUser1.FirstName != u.FirstName || testUser1.LastName != u.LastName ||
		testUser1.SessionKey != u.SessionKey {
		log.Fatalf("did not get correct user, got:\n%v vs.\n%v\n", *u, testUser1)
	}
	if err := os.RemoveAll(tstDir); err != nil {
		log.Printf("cannot clean up: %v\n", err)
	}
}
