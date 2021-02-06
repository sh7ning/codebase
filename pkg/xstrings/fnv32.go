package xstring

import "hash/fnv"

// Fnv32 将 s hash 为一个 uint32 数字
// 分区 i := fnv32(symbol) % ShardCount
func Fnv32(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
