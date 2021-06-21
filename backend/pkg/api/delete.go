package api

import (
	"github.com/go-chi/render"
	"net/http"
)

func (s *server) handleDeleteSecret(w http.ResponseWriter, r *http.Request) {
	id, removalKey := s.parseKeyAndID(r)

	secret, err := s.db.GetSecretByID(id)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		return
	}

	if secret.RemovalKey != removalKey {
		render.Status(r, http.StatusUnauthorized)
		return
	}

	if err := s.db.RemoveSecret(secret); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
}
