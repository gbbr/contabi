package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"sync"
	"time"
)

// UserStore allows managing user storage and authentication.
type UserStore interface {
	// Get returns the user based on his ID (email).
	Get(id string) (*User, error)

	// Update updates a given user with new data. The user to be updated is
	// set by setting the appropriate Email field.
	Update(*User) (*User, error)
}

// User contains data about a user.
type User struct {
	// Email holds the email of the user.
	Email string

	// Password holds the password for this user.
	Password string

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

func DefaultUserStore() UserStore {
	// TODO(gbbr): Handle error
	usr, _ := user.Current()
	return &userStore{
		location: filepath.Join(usr.HomeDir, ".contabi/users"),
	}
}

var ErrUserNotFound = errors.New("user not found")

type userStore struct {
	location string
	mu       sync.RWMutex
}

var _ UserStore = (*userStore)(nil)

func (u userStore) Get(id string) (*User, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.get(id)
}

func (u userStore) get(id string) (*User, error) {
	b, err := ioutil.ReadFile(filepath.Join(u.location, id))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	var usr User
	err = json.Unmarshal(b, &usr)
	return &usr, err
}

func (u userStore) Update(up *User) (*User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()
	usr, err := u.get(up.Email)
	if err != nil {
		return nil, err
	}
	if up.Password != "" {
		usr.Password = up.Password
	}
	if up.FirstName != "" {
		usr.FirstName = up.FirstName
	}
	if up.LastName != "" {
		usr.LastName = up.LastName
	}
	if up.SessionKey != "" {
		usr.SessionKey = up.SessionKey
		usr.SessionKeyTime = time.Now()
	}
	if !up.SessionKeyTime.IsZero() {
		usr.SessionKeyTime = up.SessionKeyTime
	}
	b, err := json.Marshal(usr)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(filepath.Join(u.location, up.Email), b, 0666)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
