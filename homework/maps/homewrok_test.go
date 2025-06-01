package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap struct {
	keys   []int
	values map[int]int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		keys:   make([]int, 0),
		values: make(map[int]int),
	}
}

func (m *OrderedMap) Insert(key, value int) {
	if !m.Contains(key) {
		index := 0
		for index < len(m.keys) && m.keys[index] < key {
			index++
		}
		m.keys = append(m.keys[:index], append([]int{key}, m.keys[index:]...)...)
	}
	m.values[key] = value
}

func (m *OrderedMap) Erase(key int) {
	if m.Contains(key) {
		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
		delete(m.values, key)
	}
}

func (m *OrderedMap) Contains(key int) bool {
	_, exists := m.values[key]
	return exists
}

func (m *OrderedMap) Size() int {
	return len(m.keys)
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	for _, key := range m.keys {
		action(key, m.values[key])
	}
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
