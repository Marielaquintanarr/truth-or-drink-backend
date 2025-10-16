package handlers

import (
	"context"
	"encoding/json"
	"go-truth-or-drink-api/models"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetTellByLevelEasy(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := pool.Query(context.Background(), "SELECT id, tell, level_id FROM tell WHERE level_id = 1")
		if err != nil {
			http.Error(w, "Error haciendo query: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tells []models.Tell
		for rows.Next() {
			var t models.Tell
			if err := rows.Scan(&t.ID, &t.Tell, &t.LevelId); err != nil {
				http.Error(w, "Error escaneando row: "+err.Error(), http.StatusInternalServerError)
				return
			}
			tells = append(tells, t)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tells)
	}
}

func GetTellByLevelMedium(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := pool.Query(context.Background(), "SELECT id, tell, level_id FROM tell WHERE level_id = 2")
		if err != nil {
			http.Error(w, "Error haciendo query: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tells []models.Tell
		for rows.Next() {
			var t models.Tell
			if err := rows.Scan(&t.ID, &t.Tell, &t.LevelId); err != nil {
				http.Error(w, "Error escaneando row: "+err.Error(), http.StatusInternalServerError)
				return
			}
			tells = append(tells, t)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tells)
	}
}

func GetTellByLevelHard(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := pool.Query(context.Background(), "SELECT id, tell, level_id FROM tell WHERE level_id = 3")
		if err != nil {
			http.Error(w, "Error haciendo query: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tells []models.Tell
		for rows.Next() {
			var t models.Tell
			if err := rows.Scan(&t.ID, &t.Tell, &t.LevelId); err != nil {
				http.Error(w, "Error escaneando row: "+err.Error(), http.StatusInternalServerError)
				return
			}
			tells = append(tells, t)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tells)
	}
}
