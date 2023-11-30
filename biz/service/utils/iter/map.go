package iter

import (
	"freezonex/openiiot/biz/service/utils/monad/option"
	"freezonex/openiiot/biz/service/utils/pair"
)

// FromMap creates an iterator from a map.
func FromMap[K comparable, V any](data map[K]V) func() option.Option[pair.Pair[K, V]] {
	pairs := make([]pair.Pair[K, V], len(data))
	index := 0
	for k, v := range data {
		pairs[index] = pair.New(k, v)
		index++
	}
	return FromSlice(pairs)
}

// CollectToMap collects an iterator into a map.
func CollectToMap[K comparable, V any](iter func() option.Option[pair.Pair[K, V]]) map[K]V {
	out := make(map[K]V, 0)
	elem := iter()
	for elem.IsSome() {
		p := elem.Unwrap()
		out[p.First] = p.Second
	}
	return out
}

// Keys returns an iterator of keys from a map.
func Keys[K comparable, V any](data map[K]V) func() option.Option[K] {
	return Map(FromMap(data), pair.GetFirst[K, V])
}

// Values returns an iterator of values from a map.
func Values[K comparable, V any](data map[K]V) func() option.Option[V] {
	return Map(FromMap(data), pair.GetSecond[K, V])
}
