package students

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/sajinlama/go-backend/internal/storage"
	"github.com/sajinlama/go-backend/internal/types"
	"github.com/sajinlama/go-backend/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
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
		lastId, err := storage.CreateStudents(
			student.Name,
			student.Email,
			student.Age,
		)
		slog.Info("user created sucessfulyy ", slog.String("userId", fmt.Sprint(lastId)))
		if err != nil {
			response.Writejson(w, http.StatusInternalServerError, err)
		}

		response.Writejson(w, http.StatusCreated, map[string]int64{
			"id": int64(lastId),
		})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		slog.Info("getting student id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.Writejson(
				w,
				http.StatusBadRequest,
				response.GeneralError(err),
			)
			return
		}

		student, err := storage.GetStudentById(intId)

		if err != nil {
			response.Writejson(
				w,
				http.StatusNotFound,
				response.GeneralError(err),
			)
			return
		}

		response.Writejson(w, http.StatusOK, student)
	}
}
