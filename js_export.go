package rand

import (
	"apigo.cc/go/jsmod"
)

func init() {
	jsmod.Register("rand", map[string]any{
		"int": func(min, max int64) int64 {
			return Int(min, max)
		},
		"fastInt": func(min, max int64) int64 {
			return FastInt(min, max)
		},
		"float": func(min, max float64) float64 {
			return Float(min, max)
		},
		"fastFloat": func(min, max float64) float64 {
			return FastFloat(min, max)
		},
		"byte":  Byte,
		"bytes": Bytes,
		"perm":  Perm,
	})
}
