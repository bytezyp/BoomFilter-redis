# BoomFilter-redis
The BoomFilter is a filter based on redis bitmap. 

A Bloom filter has two parameters : key (filer key), K (func number), option (redis-client).
It has three functions, which are one of BoomFilter, anthor of DelKey, last of Count.

In this implementation, redis-client used is [go-redis](github.com/go-redis/redis).

Support data persistence
# For example
### Initialization NewBoomFilter
````go
func main(){
	redisOption := &redis.Options{
    		Addr:               "192.168.99.64:6379",
    		DB:                 0,
    		DialTimeout:        10 * time.Second,
    		ReadTimeout:        30 * time.Second,
    		WriteTimeout:       30 * time.Second,
    		PoolSize:           10,
    		PoolTimeout:        30 * time.Second,
    		IdleTimeout:        500 * time.Millisecond,
    		IdleCheckFrequency: 500 * time.Millisecond,
    	}
    boom := BoomFilters.NewBoom("boomkey", 2 ,redisOption)
    inputString := "testString"
   
}
````
#### filter function
```go
    var flag bool
    flag = boom.BoomFilter(inputString)
```

#### Count function
```go
    var num int
    num = boom.Count()
```
#### DelKey function
```go
    var flag bool
    flag = boom.DelKey(inputString)
```



