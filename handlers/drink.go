package handlers

import (
	"context"
	"encoding/json"
	"go-truth-or-drink-api/models"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDrinkEasy(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := pool.Query(context.Background(), "SELECT id, drink, level_id FROM drink WHERE level_id = 1")
		if err != nil {
			http.Error(w, "Error haciendo query: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var drinks []models.Drink
		for rows.Next() {
			var d models.Drink
			if err := rows.Scan(&d.ID, &d.Drink, &d.LevelId); err != nil {
				http.Error(w, "Error escaneando row: "+err.Error(), http.StatusInternalServerError)
				return
			}
			drinks = append(drinks, d)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(drinks)
	}
}

func GetDrinkMedium(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := pool.Query(context.Background(), "SELECT id, drink, level_id FROM drink WHERE level_id = 2")
		if err != nil {
			http.Error(w, "Error haciendo query: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var drinks []models.Drink
		for rows.Next() {
			var d models.Drink
			if err := rows.Scan(&d.ID, &d.Drink, &d.LevelId); err != nil {
				http.Error(w, "Error escaneando row: "+err.Error(), http.StatusInternalServerError)
				return
			}
			drinks = append(drinks, d)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(drinks)
	}
}

func GetDrinkHard(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := pool.Query(context.Background(), "SELECT id, drink, level_id FROM drink WHERE level_id = 3")
		if err != nil {
			http.Error(w, "Error haciendo query: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var drinks []models.Drink
		for rows.Next() {
			var d models.Drink
			if err := rows.Scan(&d.ID, &d.Drink, &d.LevelId); err != nil {
				http.Error(w, "Error escaneando row: "+err.Error(), http.StatusInternalServerError)
				return
			}
			drinks = append(drinks, d)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(drinks)
	}
}
