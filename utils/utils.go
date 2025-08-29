package utils

import (
	"github.com/jackc/pgx/v5/pgtype"

	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func StringToPgText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{
			Valid: false,
		}
	}
	return pgtype.Text{
		String: s,
		Valid:  true,
	}
}

func Resp(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ReadBody[T any](r *http.Request) (*T, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("Error reading request body")
		return nil, err
	}

	var data T
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling request body")
		return nil, err
	}
	return &data, nil
}
