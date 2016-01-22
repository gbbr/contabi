package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

const sessionExpires = 3 * time.Hour

// UserStore allows managing user storage and authentication.
type UserStore interface {
	Get(id string) (*User, error)
	Update(*User) (*User, error)
}

// User contains data about a user.
type User struct {
	// Email holds the email of the user.
	Email string

	// FirstName holds the first name of this user.
	FirstName string

	// LastName holds the first name of this user.
	LastName string

	// SessionKey is the last session key generated for this user. If none was
	// generated or the user is logged out, it will be empty.
	SessionKey string

	// SessionKeyTime is the time when the session key was generated.
	SessionKeyTime time.Time
}

type userStore struct{ location string }

func DefaultUserStore() UserStore {
	// TODO(gbbr): Handle error
	usr, _ := user.Current()
	return &userStore{filepath.Join(usr.HomeDir, ".contabi/users")}
}

var ErrorUserNotFound = errors.New("user not found")

func (u userStore) Get(id string) (*User, error) {
	b, err := ioutil.ReadFile(filepath.Join(u.location, id))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrorUserNotFound
		}
		return nil, err
	}
	var usr User
	err = json.Unmarshal(b, &usr)
	return &usr, err
}

func (u userStore) Update(*User) (*User, error) {
	return nil, nil
}
