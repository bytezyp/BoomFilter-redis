# BoomFilter-redis
The BoomFilter is a filter based on redis bitmap. 

A Bloom filter has two parameters : key (filer key), K (func number), option (redis-client).
It has three functions, which are one of BoomFilter, anthor of DelKey, last of Count.

In this implementation, redis-client used is [go-redis](github.com/go-redis).
# For example
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


