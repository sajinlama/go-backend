package students

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sajinlama/go-backend/internal/types"
	"github.com/sajinlama/go-backend/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Students

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.Writejson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.Writejson(
				w,
				http.StatusBadRequest,
				response.GeneralError(err),
			)
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)

			response.Writejson(
				w,
				http.StatusBadRequest,
				response.ValidationError(validateErrs),
			)
			return
		}

		slog.Info("creating a student")

		response.Writejson(w, http.StatusCreated, map[string]string{
			"success": "ok",
		})
	}
}
