package main

import (
	"log"
	"net/http"

	"github.com/go-delve/delve/pkg/config"
)

func main() {
	// Carrega as configurações do projeto
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Inicializa o rate limiter
	rateLimiter := limiter.NewRateLimiter(cfg)

	// Configura o servidor web
	http.Handle("/", rateLimiter.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Requisição bem-sucedida!"))
	})))

	log.Printf("Servidor escutando na porta %s", cfg.ServerPort)
	if err := http.ListenAndServe(cfg.ServerPort, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
