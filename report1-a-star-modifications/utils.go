package AStar

import (
	"math"
	"reflect"
)

type Number interface {
	int64 | float64
}

func MaxValue[T Number](n T) T {
	switch reflect.TypeOf(n).Kind() {
	case reflect.Int64:
		return T(math.MaxInt64)
	case reflect.Float64:
		return T(math.Inf(1))
	default:
		panic("unsupported type")
	}
}

type Edge[T Number] struct {
	to int
	c  T
}
