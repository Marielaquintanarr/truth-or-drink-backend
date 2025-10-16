package handlers

import (
	"context"
	"encoding/json"
	"go-truth-or-drink-api/models"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetLevels(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := pool.Query(context.Background(), "SELECT id, level FROM level LIMIT 10")
		if err != nil {
			http.Error(w, "Error haciendo query: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var levels []models.Level
		for rows.Next() {
			var l models.Level
			if err := rows.Scan(&l.ID, &l.Level); err != nil {
				http.Error(w, "Error escaneando row: "+err.Error(), http.StatusInternalServerError)
				return
			}
			levels = append(levels, l)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(levels)
	}
}
