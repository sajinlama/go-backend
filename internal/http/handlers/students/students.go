package students

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/sajinlama/go-backend/internal/types"
	"github.com/sajinlama/go-backend/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Students

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.Writejson(w, http.StatusBadRequest, "request body is empty")
			return
		}

		if err != nil {
			response.Writejson(w, http.StatusBadRequest, err.Error())
			return
		}

		slog.Info("creating a student")

		response.Writejson(w, http.StatusCreated, map[string]string{
			"success": "ok",
		})
	}
}
