package main

import (
	"context"
	"log"
	"net/http"
	"os"

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

	mux := http.NewServeMux()

	mux.HandleFunc("/levels", handlers.GetLevels(pool))
	mux.HandleFunc("/tellEasy", handlers.GetTellByLevelEasy(pool))
	mux.HandleFunc("/tellMedium", handlers.GetTellByLevelMedium(pool))
	mux.HandleFunc("/tellHard", handlers.GetTellByLevelHard(pool))
	mux.HandleFunc("/drinkEasy", handlers.GetDrinkEasy(pool))
	mux.HandleFunc("/drinkMedium", handlers.GetDrinkMedium(pool))
	mux.HandleFunc("/drinkHard", handlers.GetDrinkHard(pool))

	handler := enableCORS(mux)

	// 📦 Render usa PORT dinámico, usa 8080 por defecto si no está definida
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor corriendo en puerto %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		allowedOrigins := []string{
			"http://localhost:5173",
			"https://truthordrinkmq.netlify.app",
		}

		for _, o := range allowedOrigins {
			if origin == o {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
