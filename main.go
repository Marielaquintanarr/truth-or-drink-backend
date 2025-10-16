package main

import (
	"context"
	"log"
	"net/http"

	"go-truth-or-drink-api/handlers"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dsn := "postgresql://postgres.omnmlbciwkbykopnidrm:Marikolas2004ssssss@aws-1-us-east-2.pooler.supabase.com:6543/postgres"

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Error parseando config: %v", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Error conectando a la DB: %v", err)
	}

	// Mux (router)
	mux := http.NewServeMux()

	// Endpoints
	mux.HandleFunc("/levels", handlers.GetLevels(pool))
	mux.HandleFunc("/tellEasy", handlers.GetTellByLevelEasy(pool))
	mux.HandleFunc("/tellMedium", handlers.GetTellByLevelMedium(pool))
	mux.HandleFunc("/tellHard", handlers.GetTellByLevelHard(pool))
	mux.HandleFunc("/drinkEasy", handlers.GetDrinkEasy(pool))
	mux.HandleFunc("/drinkMedium", handlers.GetDrinkMedium(pool))
	mux.HandleFunc("/drinkHard", handlers.GetDrinkHard(pool))

	// 🌍 Envolvemos todo con el middleware CORS
	handler := enableCORS(mux)

	log.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}

// Middleware CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permitir peticiones desde cualquier origen (o tu frontend)
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Responder rápido a las peticiones OPTIONS (preflight)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Continuar al siguiente handler
		next.ServeHTTP(w, r)
	})
}
