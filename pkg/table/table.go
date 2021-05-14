package table

import (
	"DataStructures/HashTable/pkg/hash"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

type KeyValue struct {
	key   string
	value interface{}
}

type IObject interface {
	Copy()
	ToString()
}

func NewKeyValue(key string, value interface{},
) KeyValue {
	return KeyValue{
		key:   key,
		value: value,
	}
}

type Table struct {
	content [][]KeyValue
	length  uint32
	seek    uint32
}

func keyToString(key interface{}) string {
	return fmt.Sprintf("%+v-%T", key, key)
}

func (table *Table) Get(key interface{}) (interface{}, error) {
	keyString := keyToString(key)
	index := hash.Murmur3_32Bits(keyString, table.seek) % table.length
	for _, keyValue := range table.content[index] {
		if keyValue.key == keyString {
			return keyValue.value, nil
		}
	}
	return nil, errors.New("element not found in hash table with the provided key")
}

func (table *Table) Set(key interface{}, value interface{}) {
	keyString := keyToString(key)
	index := hash.Murmur3_32Bits(keyString, table.seek) % table.length
	for keyValueIndex, keyValue := range table.content[index] {
		if keyValue.key == keyString {
			table.content[index][keyValueIndex].value = value
			return
		}
	}
	table.content[index] = append(table.content[index], NewKeyValue(keyString, value))
}

func NewTable(length uint32) *Table {
	seek, generationError := rand.Int(rand.Reader, big.NewInt(2147483647))
	if generationError != nil {
		panic(generationError)
	}
	return &Table{
		content: make([][]KeyValue, length),
		length:  length,
		seek:    uint32(seek.Uint64()),
	}
}
