package login

import "net/http"

type LoginUser interface {
	LoginUser(rw http.ResponseWriter, r *http.Request)
}
