package router

import (
	"io"
	"net/http"

	"github.com/aaruel/kite-testenv/gateway/repo"
	"github.com/aaruel/kite-testenv/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Ideally, generating the router structure from the repo structure
func NewRouter(repo *repo.Repo) {
	r := chi.NewRouter()

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	r.Use(corsOpts.Handler)

	r.Route("/math", func(r chi.Router) {
		r.Route("/square/{number}", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				number := chi.URLParam(r, "number")
				response := repo.Query("math", "square", utils.StringToFloat(number))
				n := response.MustFloat64()
				io.WriteString(w, utils.FloatToString(n))
			})
		})
	})

	http.ListenAndServe(":3579", r)
}
