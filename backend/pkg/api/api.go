package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"net/http"
	"sb14project_backend/pkg/database"
)

type server struct {
	db *database.DB
}

func NewServer(addr string, db *database.DB) error {
	s := &server{db}

	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
		MaxAge:         300,
	}))

	r.Get("/", s.handleGetSecret)
	r.Post("/", s.handleAddSecret)
	r.Delete("/", s.handleDeleteSecret)

	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "route not found")
	})

	fmt.Printf("listening on %v\n", addr)

	return http.ListenAndServe(addr, r)
}
