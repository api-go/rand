package rand

import (
	"apigo.cc/go/jsmod"
)

func init() {
	jsmod.Register("rand", map[string]any{
		"Int":       jsInt,
		"FastInt":   jsFastInt,
		"Float":     jsFloat,
		"FastFloat": jsFastFloat,
		"Byte":      Byte,
		"Bytes":     Bytes,
		"Perm":      Perm,
	})
}

func jsInt(min, max int64) int64 {
	return Int(min, max)
}

func jsFastInt(min, max int64) int64 {
	return FastInt(min, max)
}

func jsFloat(min, max float64) float64 {
	return Float(min, max)
}

func jsFastFloat(min, max float64) float64 {
	return FastFloat(min, max)
}
