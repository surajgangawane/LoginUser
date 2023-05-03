package register

import "net/http"

type RegisterInt interface {
	Register(rw http.ResponseWriter, r *http.Request)
}
