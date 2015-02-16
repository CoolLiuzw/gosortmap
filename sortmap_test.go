package sortmap_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/tg/gosortmap"
)

func ExampleByKey() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, e := range sortmap.ByKey(m) {
		fmt.Printf("%s: %d\n", e.K, e.V)
	}
	// Output:
	// apple: 4
	// banana: 1
	// cabbage: 3
	// daikon: 2
}

func ExampleByKeyDesc() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, e := range sortmap.ByKeyDesc(m) {
		fmt.Printf("%s: %d\n", e.K, e.V)
	}
	// Output:
	// daikon: 2
	// cabbage: 3
	// banana: 1
	// apple: 4
}

func ExampleByVal() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, e := range sortmap.ByVal(m) {
		fmt.Printf("%s: %d\n", e.K, e.V)
	}
	// Output:
	// banana: 1
	// daikon: 2
	// cabbage: 3
	// apple: 4
}

func ExampleByValDesc() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, e := range sortmap.ByValDesc(m) {
		fmt.Printf("%s: %d\n", e.K, e.V)
	}
	// Output:
	// apple: 4
	// cabbage: 3
	// daikon: 2
	// banana: 1
}

var benchMap = func() map[int]int {
	m := make(map[int]int)
	for n := 0; n < 10000; n++ {
		m[rand.Int()] = rand.Int()
	}
	return m
}()

type kv struct{ k, v int }
type kvs []kv

func (m kvs) Len() int           { return len(m) }
func (m kvs) Less(i, j int) bool { return m[i].k < m[j].k }
func (m kvs) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

type kvs_nosort []kv

func (m kvs_nosort) Len() int           { return len(m) }
func (m kvs_nosort) Less(i, j int) bool { return false }
func (m kvs_nosort) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func BenchmarkManualSorted(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(kvs_nosort, 0, len(benchMap))
		for k, v := range benchMap {
			m = append(m, kv{k, v})
		}
		sort.Sort(m)
	}
}

func BenchmarkManualFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(kvs, 0, len(benchMap))
		for k, v := range benchMap {
			m = append(m, kv{k, v})
		}
		sort.Sort(m)
	}
}

func BenchmarkManualKey(b *testing.B) {
	for n := 0; n < b.N; n++ {
		keys := make([]int, 0, len(benchMap))
		for k, _ := range benchMap {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		values := make([]int, len(keys))
		for n := range keys {
			values[n] = benchMap[keys[n]]
		}
	}
}

func BenchmarkSortSorted(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortmap.ByFunc(benchMap, func(x, y sortmap.KV) bool { return false })
	}
}

func BenchmarkSortFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortmap.ByFunc(benchMap, func(x, y sortmap.KV) bool { return x.K.(int) < y.K.(int) })
	}
}

func BenchmarkSortKey(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortmap.ByKey(benchMap)
	}
}
