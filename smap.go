package smap

import "golang.org/x/exp/constraints"

type NType interface {
	constraints.Integer | constraints.Float
}

type BType interface {
	[]byte | string
}

func N[T NType, R NType](from []T) []R {
	to := make([]R, len(from))
	for i, v := range from {
		to[i] = R(v)
	}
	return to
}

func S[T BType, R BType](from []T) []R {
	to := make([]R, len(from))
	for i, v := range from {
		to[i] = R(v)
	}
	return to
}

func F[T any, R any](from []T, f func(T) R) []R {
	to := make([]R, len(from))
	for i, v := range from {
		to[i] = f(v)
	}
	return to
}
