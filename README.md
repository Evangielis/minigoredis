# PACKAGE DOCUMENTATION

```` go
package minigoredis
    import "minigoredis"
````

# TYPES

```` go
type Minigoredis struct {
    Miniredis *miniredis.Miniredis
}
    Minigoredis -- struct which wraps a miniredis instance

func Run() (*Minigoredis, error)
    Run -- starts a miniredis server and serves it wrapped in Gominiredis

func (g *Minigoredis) Close() error
    Close -- Closes the miniredis server. Never throws an error

func (g *Minigoredis) Get(key string) *redis.StringCmd
    Get -- Retrieves a value from miniredis wrapped in a go-redis StringCmd

func (g *Minigoredis) Set(key string, val string, _ time.Duration) *redis.StatusCmd
    Set -- Sets a key, value pair in miniredis
````
