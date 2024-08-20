package req

import (
	"encoding/json"
	"net/http"
)

func Decode[T any](r *http.Request, dst *T) error {
	d := json.NewDecoder(r.Body)
	return d.Decode(dst)
}
