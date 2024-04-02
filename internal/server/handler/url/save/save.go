package save

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/VadimRight/Go_WebApp/internal/lib/api/response"
	"github.com/VadimRight/Go_WebApp/internal/lib/logger/sl"
	"github.com/VadimRight/Go_WebApp/internal/lib/random"
	"github.com/VadimRight/Go_WebApp/internal/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

//	"github.com/VadimRight/Go_WebApp/internal/storage/postgres"
)

type Request struct {
	Id uuid.UUID
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

type URLHadler interface {
	 SaveURL(urltosave string, alias_name string) (string, error)
}

const aliasLength = 6

func New(log *slog.Logger, urlSave URLHadler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "hadlers.url.save.New"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("failed to decode request body")
			render.JSON(w, r, response.Error("empty request"))
			return
		}

		if err != nil {
			log.Error("failed to decode request body", sl.Error(err))
			render.JSON(w, r, response.Error("failed to decode request"))
			return

		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("invalid request", sl.Error(err))
			render.JSON(w, r, response.ValidationError(validateErr))
			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRundomString(aliasLength)
		}

		url, err := urlSave.SaveURL(req.URL, alias)
		if errors.Is(err, storage.ErrUrlExists) {
			log.Info("url already exists", slog.String("url", req.URL))
			render.JSON(w, r, response.Error("url already exists"))
			return
		}
		if err != nil {
			log.Error("failed to add url", sl.Error(err))
			render.JSON(w, r, response.Error("failed to add url"))
			return
		}
		log.Info("url added",  slog.String("url", url))
		responseOK(w, r, alias)
	}

}

func responseOK(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Response: response.OK(),
		Alias:    alias,
	})
}
