package storage

type StorageStrategy interface {
    Increment(key string) (int, error)
    Get(key string) (int, error)
    Reset(key string) error
}

type RedisStrategy struct {
    // Adicione aqui os campos necessários para a conexão com o Redis
}

func NewRedisStrategy() *RedisStrategy {
    return &RedisStrategy{
        // Inicialize a conexão com o Redis aqui
    }
}

func (r *RedisStrategy) Increment(key string) (int, error) {
    // Implemente a lógica para incrementar o contador no Redis
    return 0, nil
}

func (r *RedisStrategy) Get(key string) (int, error) {
    // Implemente a lógica para obter o contador do Redis
    return 0, nil
}

func (r *RedisStrategy) Reset(key string) error {
    // Implemente a lógica para resetar o contador no Redis
    return nil
}