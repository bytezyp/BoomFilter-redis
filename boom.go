package BoomFilter_redis

import (
"github.com/go-redis/redis"
)

const (
	HASH_LOOP_NUM = 4
)

type hashFunc func(string)int64

/*
*	Hash function list
*/
var hashList = []hashFunc{
	BKDRHash,
	APHash,
	DJBHash64,
	ELFHash64,
}
/*
*	type BoomFilter
*/
type Boom struct {
	Key string
	K int
	Con *redis.Client
}
/*
*	return Boom
*	method: a new and default boomFilter
*/
func NewDefaultBoom(Key string, options *redis.Options) *Boom  {
	return &Boom{
		Key:Key,
		K:HASH_LOOP_NUM,
		Con:redis.NewClient(options),
	}
}
/*
*	return Boom
*	method: a new and customiz boomFilter
*/
func NewBoom(Key string, K int,  options *redis.Options) *Boom {
	if K <= HASH_LOOP_NUM && K > 0 {
		return NewDefaultBoom(Key,options)
	}
	return &Boom{
		Key:Key,
		K:K,
		Con:redis.NewClient(options),
	}
}
/*
*	return int64
*	method: generate random numbers
*/
func randomGenerator(hash_value int64) int64 {
	return hash_value % (1 << 32)
}
/*
*	return bool
*	method: verify whether the value exists
*/
func (Bf *Boom)BoomFilter(filterVal string)  bool {
	flag := false
	pipeliner := Bf.Con.Pipeline()
	for i:= 0; i < Bf.K; i++ {
		hashFunc := hashList[i]
		hash := hashFunc(filterVal)
		offset := randomGenerator(hash)
		if getBitRes, _ := Bf.Con.GetBit(Bf.Key, offset).Result();getBitRes == 0 {
			pipeliner.Process(Bf.Con.SetBit(Bf.Key,offset, 1))
			//if _, err:= Bf.Con.SetBit(Bf.Key,offset, 1).Result(); err == nil  {
			flag = true
			//}
		}
	}
	_,err := pipeliner.Exec()
	if err != nil {
		return false
	}
	return flag
}
/*
*	return int
*	method: sum amount of key bitmap
*/
func (Bf *Boom)Count() int{
	if count, err := Bf.Con.BitCount(Bf.Key, nil).Result(); err == nil {
		return int(count)
	}
	return 0
}
/*
*	return bool
* 	method: del key bitmap
*/

func (Bf *Boom)DelKey(key string) bool {
	if _, err := Bf.Con.Del(key).Result();err != nil {
		return true
	}
	return false
}


