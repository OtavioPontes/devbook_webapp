package cookies

import (
	"devbook_webapp/src/config"
	"net/http"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, Id, token string) error {
	data := map[string]string{
		"id":    Id,
		"token": token,
	}

	dataEncripted, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    dataEncripted,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}
