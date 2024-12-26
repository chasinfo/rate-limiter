package limiter

import (
	"net/http"
	"time"

	"github.com/gorilla/context"
)

type RateLimiter struct {
	// Defina os campos necessários para o rate limiter
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		// Inicialize os campos necessários
	}
}

func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr // Obtenha o endereço IP do cliente
		token := r.Header.Get("API_KEY") // Obtenha o token do cabeçalho

		// Lógica para verificar limites por IP e token
		if exceeded := rl.checkLimits(ip, token); exceeded {
			http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
			return
		}

		// Chame o próximo handler se não houver limite excedido
		next.ServeHTTP(w, r)
	})
}

func (rl *RateLimiter) checkLimits(ip, token string) bool {
	// Implemente a lógica para verificar se os limites foram excedidos
	return false
}