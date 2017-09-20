# PACKAGE DOCUMENTATION

```` go
package minigoredis
    import "minigoredis"
````

# INSTALLATION

```` bash
go get github.com/Evangielis/minigoredis
````

# USAGE
In the example below, the function starts up a miniredis server, stores a key then tests a function which retrieves a key from the cache.

```` go
func TestGetPriorTimestampSuccess(t *testing.T) {
	key, val := "foo", "bar"
	cache, _ := minigoredis.Run()   // Here's where the server is started up
    defer cache.Close()                 // Close it to avoid memory leaks
	cache.Set(key, val, 0)          // Setting a key on the cache
	time := getPriorTimestamp(key, cache)
	if time != val {
		t.Fatal("Timestamp retrieved was not equal to cached value")
		t.Fail()
	}
}
````

# TYPES

```` godoc
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
