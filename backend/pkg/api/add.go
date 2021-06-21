package api

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/thanhpk/randstr"
	"net/http"
	"sb14project_backend/pkg/database"
	"sb14project_backend/pkg/encryption"
)

type addSecretRequest struct {
	Content string `json:"content"`
	Uses    *int   `json:"uses"`
}

func (a addSecretRequest) Bind(*http.Request) error {
	if a.Uses != nil && *a.Uses <= 0 {
		return errors.New("uses should be greater than 0")
	}

	if a.Content == "" {
		return errors.New("content should not be empty")
	}

	return nil
}

func (s *server) handleAddSecret(w http.ResponseWriter, r *http.Request) {
	var data addSecretRequest
	if err := render.Bind(r, &data); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	encryptedContent, key, err := encryption.Encrypt(data.Content)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	removalKey := randstr.String(32)

	secret := database.Secret{
		Content:    encryptedContent,
		RemovalKey: removalKey,
		UsagesLeft: data.Uses,
	}

	if err := s.db.InsertSecret(&secret); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, render.M{
		"key":        key,
		"removalKey": removalKey,
		"secret":     secret,
	})
}
