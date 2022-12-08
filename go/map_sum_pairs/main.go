package main

import (
	"sort"
	"strings"
)

type MapSum struct {
	Map  map[string]int
	keys []string
}

func Constructor() MapSum {
	return MapSum{
		Map:  make(map[string]int, 128),
		keys: make([]string, 0, 128),
	}
}

func (mapSum *MapSum) Insert(key string, val int) {
	mapSum.Map[key] = val

	i := sort.SearchStrings(mapSum.keys, key)
	if i < len(mapSum.keys) && mapSum.keys[i] == key {
		return
	}

	mapSum.keys = append(mapSum.keys, "")
	copy(mapSum.keys[i+1:], mapSum.keys[i:])
	mapSum.keys[i] = key
}

func (mapSum *MapSum) Sum(prefix string) int {
	sum := 0
	i := sort.SearchStrings(mapSum.keys, prefix)
	for i < len(mapSum.keys) && strings.HasPrefix(mapSum.keys[i], prefix) {
		sum += mapSum.Map[mapSum.keys[i]]
		i++
	}
	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
