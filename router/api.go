package router

import (
	"be-golang-chapter-36-implem/infra"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupReouter(ctx infra.Context) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", ctx.Handler.CustomerHandler.Login)
	})

	r.Route("/customer", func(r chi.Router) {
		r.Post("/", ctx.Handler.CustomerHandler.Create)
		r.Get("/", ctx.Handler.CustomerHandler.GetAll)
		r.Get("/{id}", ctx.Handler.CustomerHandler.Create)
	})

	fmt.Println("server start on port ", ctx.Config.Port)
	http.ListenAndServe(":"+ctx.Config.Port, r)
}
