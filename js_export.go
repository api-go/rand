package rand

import (
	"apigo.cc/go/jsmod"
)

func init() {
	jsmod.Register("rand", map[string]any{
		"Int": func(min, max int64) int64 {
			return Int(min, max)
		},
		"FastInt": func(min, max int64) int64 {
			return FastInt(min, max)
		},
		"Float": func(min, max float64) float64 {
			return Float(min, max)
		},
		"FastFloat": func(min, max float64) float64 {
			return FastFloat(min, max)
		},
		"Byte":  Byte,
		"Bytes": Bytes,
		"Perm":  Perm,
	})
}
