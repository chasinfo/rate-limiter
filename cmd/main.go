package main

import (
	"log"
	"net/http"
	"time"

	"github.com/chasinfo/rate-limiter/internal/config"
	"github.com/chasinfo/rate-limiter/internal/limiter"
)

func main() {
	// Carrega as configurações do projeto
	cfg := config.LoadConfig()

	// Inicializa o rate limiter
	rateLimiter := limiter.NewRateLimiter(map[string]int{
		"BlockDuration":  cfg.BlockDuration,
		"RateLimitIP":    cfg.RateLimitIP,
		"RateLimitToken": cfg.RateLimitToken,
	}, time.Hour)

	// Configura o servidor web
	http.Handle("/", rateLimiter.Limit(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Requisição bem-sucedida!"))
	})))

	log.Printf("Servidor escutando na porta %s", cfg.ServerPort)
	if err := http.ListenAndServe(cfg.ServerPort, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
