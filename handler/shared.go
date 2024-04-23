package handler

import (
	"log/slog"
	"net/http"
)

func logError(r *http.Request, err error) {
	slog.Error("internal server error", "error", err, "url", r.URL.Path)
}
