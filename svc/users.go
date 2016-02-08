package svc

import (
	"net/http"
	"time"

	"github.com/gbbr/contabi/store"
)

type UserService interface {
	IsValidRequest(r *http.Request) (bool, error)
}

var Users UserService = userService{}

type userService struct{}

var _ UserService = (*userService)(nil)

const sessionExpires = 3 * time.Hour

// IsValidRequest returns true if the request contains valid authentication
// information.
func (_ userService) IsValidRequest(r *http.Request) (bool, error) {
	id := r.Header.Get("X-Auth-User")
	t := r.Header.Get("X-Auth-Token")
	if id == "" || id == "" {
		return false, nil
	}
	st := store.DefaultUserStore()
	usr, err := st.Get(id)
	if err != nil {
		return false, err
	}
	if usr.SessionKey != t {
		return false, nil
	}
	if usr.SessionKeyTime.IsZero() || usr.SessionKeyTime.Sub(time.Now()) > sessionExpires {
		_, err := st.Update(&store.User{SessionKey: ""})
		return false, err
	}
	return true, nil
}
