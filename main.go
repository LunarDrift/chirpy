package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/LunarDrift/chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32 // allows us to safely increment and read int value across multiple goroutines (HTTP requests)
	db             *database.Queries
	platform       string
	secret         string
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("could not find environment file: ", err)
	}

	// get db info from environment
	dbURL := os.Getenv("DB_URL")
	dbPlatform := os.Getenv("PLATFORM")
	dbSecret := os.Getenv("SECRET")

	// open connection to database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println("could not connect to database: ", err)
	}

	dbQueries := database.New(db)

	// create instance of struct
	apiCfg := &apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
		platform:       dbPlatform,
		secret:         dbSecret,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))

	// register new Handlers
	mux.HandleFunc("GET /admin/metrics", apiCfg.handleShowMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handleResetMetrics)

	mux.HandleFunc("GET /api/healthz", handleHealthZ)
	mux.HandleFunc("GET /api/chirps", apiCfg.handleGetAllChirps)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", apiCfg.handleDeleteChirpByID)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handleGetChirpByID)
	mux.HandleFunc("POST /api/chirps", apiCfg.handleCreateChirp)
	mux.HandleFunc("POST /api/users", apiCfg.handleCreateUser)
	mux.HandleFunc("PUT /api/users", apiCfg.handleUpdateUser)
	mux.HandleFunc("POST /api/login", apiCfg.handleUserLogin)
	mux.HandleFunc("POST /api/refresh", apiCfg.handleRefreshAccessToken)
	mux.HandleFunc("POST /api/revoke", apiCfg.handleRevokeRefreshToken)
	mux.HandleFunc("POST /api/polka/webhooks", apiCfg.handleUpgradeUserStatus)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Listening on port: ", httpServer.Addr)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
