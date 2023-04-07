package cookies

import (
	"net/http"
	"time"
	"web/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configure encrypt the cookie data
func Configure() {
	s = securecookie.New(config.HASH_KEY, config.BLOCK_KEY)
}

// Save is to the auth data
func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encryptedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encryptedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Read return the values on cookie
func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	if err = s.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}

// Remove delete the value on cookie
func Remove(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
