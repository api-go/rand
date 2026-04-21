package rand_test

import (
	"testing"

	"apigo.cc/go/rand"
)

func TestGenericInt(t *testing.T) {
	// 1. 普通正数区间
	for range 100 {
		v := rand.Int(10, 20)
		if v < 10 || v > 20 { t.Errorf("Int failed: %d", v) }
	}

	// 2. 负数区间
	for range 100 {
		v := rand.Int(-50, -10)
		if v < -50 || v > -10 { t.Errorf("Negative range failed: %d", v) }
	}

	// 3. 防御性测试
	if rand.Int(100, 10) != 100 { t.Error("Reverse range should return min") }
}

func TestFastInt(t *testing.T) {
	for range 100 {
		v := rand.FastInt(1, 100)
		if v < 1 || v > 100 { t.Errorf("FastInt failed: %d", v) }
	}
}

func TestBytes(t *testing.T) {
	// 已纠正：不再检查 byte < 0
	b := rand.Byte()
	_ = b 

	bs := rand.Bytes(10)
	if len(bs) != 10 { t.Error("Bytes(10) length failed") }
}
