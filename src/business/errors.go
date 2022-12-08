package business

import (
	"fmt"
	"net/http"
)

func NewError(w http.ResponseWriter, errorCode int, errorMessage string) {
	w.WriteHeader(errorCode)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"code":"%d","message":"%s"}`, errorCode, errorMessage)))
}
