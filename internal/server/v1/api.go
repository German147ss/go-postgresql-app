// internal/server/v1/api.go
package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/orlmonteverde/go-postgres-microblog/internal/data"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{Repository: &data.UserRepository{Data: data.New()}}

	r.Mount("/users", ur.Routes())

	pr := &PostRouter{Repository: &data.PostRepository{Data: data.New()}}

	r.Mount("/posts", pr.Routes())
	return r
}
