package BoomFilter_redis

// go-BKDRHash
func BKDRHash(str string) int64 {
	hash := int64(0)
	len := len(str)
	seed := getBKDRHashSeed(4)
	for i := 0; i < len; i++ {
		hash = (hash * seed) + int64(str[i])
	}
	return hash & 0x7FFFFFFF
}

func getBKDRHashSeed(n int) int64 {
	if n <= 0 {
		return 31
	}
	j := n + 2
	r := 0
	for i:= 0; i <= j; i++ {
		if i % 2 > 0 {
			r = r * 10 + 3
		} else {
			r = r * 10 +1
		}
	}
	return int64(r)
}
// AP Hash Function
func APHash(pstr string) int64 {
	var hash int64 = 0;
	str := []byte(pstr)
	for i := 0; i < len(str); i++ {
		if ((i & 1) == 0) {
			hash ^= ((hash << 7) ^ int64(str[i]) ^ (hash >> 3));
		} else {
			hash ^= (^((hash << 11) ^ int64(str[i]) ^ (hash >> 5)) + 1);
		}
	}
	return (hash & 0x7FFFFFFF);
}
// DJB Hash Function 64
func DJBHash64(pstr string) int64 {
	var hash int64 = 5381;
	str := []byte(pstr)
	for i := 0; i < len(str); i++ {
		hash += (hash << 5) + int64(str[i]);
	}
	return (hash & 0x7FFFFFFFFFFFFFFF);
}

func ELFHash64(pstr string) int64 {
	var hash int64 = 0;
	var x    int64 = 0;
	str := []byte(pstr)
	for i := 0; i < len(str); i++ {
		hash = (hash << 4) + int64(str[i]);
		if x = hash & 0xF0000000; x != 0 {
			hash ^= (x >> 24);
			hash &= ^x + 1;
		}
	}
	return (hash & 0x7FFFFFFFFFFFFFFF);
}