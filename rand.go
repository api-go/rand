package rand

import (
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

var seedCounter atomic.Uint64
var pcgPool = sync.Pool{
	New: func() any {
		return rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), seedCounter.Add(1)))
	},
}

// IntegerType 约束支持所有整数类型
type IntegerType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// FloatType 约束支持所有浮点类型
type FloatType interface {
	~float32 | ~float64
}

// --- 基础整数 ---

// Int 获取 [min, max] 闭区间随机整数
func Int[T IntegerType](min, max T) T {
	if max < min { return min }
	diff := uint64(max - min)
	return min + T(rand.Uint64N(diff+1))
}

// FastInt 获取 [min, max] 闭区间随机整数 (高性能并发模式)
func FastInt[T IntegerType](min, max T) T {
	if max < min { return min }
	inst := pcgPool.Get().(*rand.Rand)
	diff := uint64(max - min)
	res := inst.Uint64N(diff + 1)
	pcgPool.Put(inst)
	return min + T(res)
}

// --- 浮点数 ---

// Float 获取 [min, max) 区间随机浮点数
func Float[T FloatType](min, max T) T {
	if max <= min { return min }
	return min + T(rand.Float64())*T(max-min)
}

// FastFloat 获取 [min, max) 区间随机浮点数 (高性能并发模式)
func FastFloat[T FloatType](min, max T) T {
	if max <= min { return min }
	inst := pcgPool.Get().(*rand.Rand)
	res := inst.Float64()
	pcgPool.Put(inst)
	return min + T(res)*T(max-min)
}

// --- 字节与切片 ---

// Byte 获取一个随机字节
func Byte() byte {
	return byte(rand.Uint32N(256))
}

// Bytes 获取长度为 n 的随机字节切片
func Bytes(n int) []byte {
	if n <= 0 { return []byte{} }
	b := make([]byte, n)
	for i := range b {
		b[i] = Byte()
	}
	return b
}

// --- 集合辅助 ---

// Perm 获取 [0, n) 的随机排列
func Perm(n int) []int {
	if n <= 0 { return []int{} }
	return rand.Perm(n)
}

// Shuffle 随机洗牌
func Shuffle(n int, swap func(i, j int)) {
	if n <= 0 { return }
	rand.Shuffle(n, swap)
}
