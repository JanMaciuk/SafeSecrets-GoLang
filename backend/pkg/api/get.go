package api

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
	"sb14project_backend/pkg/database"
	"sb14project_backend/pkg/encryption"
	"strconv"
	"strings"
)

func (s *server) parseKeyAndID(r *http.Request) (int, string) {
	arr := strings.Split(r.FormValue("key"), ":")

	if len(arr) != 2 {
		return 0, ""
	}

	id, _ := strconv.Atoi(arr[0])
	key := arr[1]

	return id, key
}

func (s *server) getSecret(r *http.Request) (*database.Secret, string, error) {
	id, key := s.parseKeyAndID(r)
	secret, err := s.db.GetSecretByID(id)

	if err != nil {
		return nil, "", err
	}

	output, err := encryption.Decrypt(key, secret.Content)
	if err != nil {
		return nil, "", err
	}

	return secret, output, nil
}

func (s *server) handleGetSecret(w http.ResponseWriter, r *http.Request) {
	secret, content, err := s.getSecret(r)

	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}

	// -1 = infinite usages
	if secret.UsagesLeft != nil {
		if *secret.UsagesLeft == 0 {
			render.Status(r, http.StatusGone)
			render.JSON(w, r, "usage limit exceeded")
			return
		}

		*secret.UsagesLeft--

		if err := s.db.UpdateSecret(secret); err != nil {
			fmt.Println(err)
			return
		}
	}

	secret.Content = content

	render.JSON(w, r, render.M{"secret": secret})
}
