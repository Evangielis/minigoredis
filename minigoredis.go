package minigoredis

import "github.com/alicebob/miniredis"
import "github.com/go-redis/redis"
import "time"

// Minigoredis -- struct which wraps a miniredis instance
type Minigoredis struct {
	Miniredis *miniredis.Miniredis
}

// Run -- starts a miniredis server and serves it wrapped in Gominiredis
func Run() (*Minigoredis, error) {
	mr, err := miniredis.Run()
	return &Minigoredis{
		Miniredis: mr,
	}, err
}

// Get -- Retrieves a value from miniredis wrapped in a go-redis StringCmd
func (g *Minigoredis) Get(key string) *redis.StringCmd {
	val, err := g.Miniredis.Get(key)
	if err != nil {
		panic(err)
	}
	return redis.NewStringResult(val, nil)
}

// Set -- Sets a key, value pair in miniredis
func (g *Minigoredis) Set(key string, val interface{}, _ time.Duration) *redis.StatusCmd {
	err := g.Miniredis.Set(key, val.(string))
	if err != nil {
		panic(err)
	}
	return redis.NewStatusResult("OK", err)
}

// Close -- Closes the miniredis server.  Never throws an error
func (g *Minigoredis) Close() error {
	g.Miniredis.Close()
	return nil
}
